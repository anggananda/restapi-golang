package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"restapi-golang/services"
	"restapi-golang/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/cas.v2"
)

type CASHandler struct {
	casClient   *cas.Client
	casURL      *url.URL
	frontendURL string
	hostUrl     string
	UserService *services.UserService
}

func NewCASHandler(casURL, frontendURL, hostUrl string, service *services.UserService) (*CASHandler, error) {
	parsedURL, err := url.Parse(casURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing CAS URL: %v", err)
	}

	client := cas.NewClient(&cas.Options{
		URL: parsedURL,
	})

	return &CASHandler{
		casClient:   client,
		casURL:      parsedURL,
		frontendURL: frontendURL,
		hostUrl:     hostUrl,
		UserService: service,
	}, nil
}

func (h *CASHandler) CASMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")

		handler := h.casClient.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if username := cas.Username(r); username != "" {
				c.Set("cas_username", username)
				c.Set("cas_attributes", cas.Attributes(r))
				log.Printf("CAS User authenticated: %s", username)
			}
			c.Next()
		}))

		handler.ServeHTTP(c.Writer, c.Request)
	}
}

func (h *CASHandler) LoginHandler(c *gin.Context) {
	log.Printf("Login handler called")

	h.clearCookies(c)

	logoutURL := fmt.Sprintf("%s/logout", h.casURL.String())
	client := &http.Client{
		Timeout: 2 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", logoutURL, nil)
	if err == nil {
		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			log.Printf("Pre-login logout completed")
		}
	}

	serviceURL := fmt.Sprintf("%s/api/v1/auth/callback", h.hostUrl)

	loginURL := fmt.Sprintf("%s/login?service=%s",
		h.casURL.String(),
		url.QueryEscape(serviceURL))

	log.Printf("Redirecting to CAS login: %s", loginURL)

	c.Redirect(http.StatusFound, loginURL)
}

func (h *CASHandler) CallbackHandler(c *gin.Context) {
	username, exists := c.Get("cas_username")
	if !exists || username == "" {
		log.Println("CAS username not found")
		errorURL := fmt.Sprintf("%s/login?error=auth_failed", h.frontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	usernameStr := username.(string)
	log.Printf("Processing login for: %s", usernameStr)

	_, err := h.UserService.CheckUserByUsername(ctx, usernameStr)
	if err != nil {
		log.Printf("User not found: %s, error: %v", usernameStr, err)
		errorURL := fmt.Sprintf("%s/login?error=user_not_found", h.frontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	tokenString, expirationTime, err := utils.GenerateJWT(usernameStr)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		errorURL := fmt.Sprintf("%s/login?error=token_failed", h.frontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	callbackURL := fmt.Sprintf("%s/sso/callback#token=%s&expires_at=%d&username=%s",
		h.frontendURL,
		url.QueryEscape(tokenString),
		expirationTime,
		url.QueryEscape(usernameStr))

	log.Printf("Login successful, redirecting to frontend: %s", usernameStr)
	c.Redirect(http.StatusFound, callbackURL)
}

func (h *CASHandler) LogoutHandler(c *gin.Context) {
	log.Printf("Logout handler called")

	h.clearCookies(c)

	logoutURL := fmt.Sprintf("%s/logout", h.casURL.String())

	client := &http.Client{
		Timeout: 3 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", logoutURL, nil)
	if err == nil {

		for _, cookie := range c.Request.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			log.Printf("CAS server logout completed with status: %d", resp.StatusCode)
		}
	}

	timestamp := time.Now().UnixNano()
	frontendURL := fmt.Sprintf("%s?logout=true&_=%d", h.frontendURL, timestamp)

	finalLogoutURL := fmt.Sprintf("%s/logout?service=%s",
		h.casURL.String(),
		url.QueryEscape(frontendURL))

	log.Printf("Final redirect to CAS logout with service: %s", finalLogoutURL)

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Clear-Site-Data", "\"cookies\"")

	c.Redirect(http.StatusFound, finalLogoutURL)
}

func (h *CASHandler) clearCookies(c *gin.Context) {
	cookiesToClear := []string{
		"CASTGC",
		"JSESSIONID",
		"SESSION",
		"TGC",
		"token",
		"auth",
	}

	paths := []string{"/", "/cas", "/cas/"}
	domains := []string{"", "sso.undiksha.ac.id", ".undiksha.ac.id"}

	for _, cookieName := range cookiesToClear {
		for _, path := range paths {
			for _, domain := range domains {
				c.SetCookie(cookieName, "", -1, path, domain, false, true)
				c.SetCookie(cookieName, "", -1, path, domain, true, true)
			}
		}
	}

	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
		if cookie.Path != "" {
			c.SetCookie(cookie.Name, "", -1, cookie.Path, "", false, true)
		}
	}

	log.Printf("Cleared all cookies including JSESSIONID")
}

func (h *CASHandler) HealthCheckHandler(c *gin.Context) {

	resp, err := http.Get(h.casURL.String())
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("CAS server unavailable: %v", err),
		})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":  "error",
			"message": fmt.Sprintf("CAS server returned status: %d", resp.StatusCode),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "CAS server is reachable",
		"data": gin.H{
			"url":          h.casURL,
			"status_code":  resp.StatusCode,
			"service_url":  h.hostUrl,
			"current_time": time.Now(),
		},
	})
}
