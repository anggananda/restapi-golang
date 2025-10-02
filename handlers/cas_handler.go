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

// CASMiddleware - middleware untuk handle CAS authentication
func (h *CASHandler) CASMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Prevent caching
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

// LoginHandler - handle login request
func (h *CASHandler) LoginHandler(c *gin.Context) {
	log.Printf("Login handler called")

	// Clear cookies dengan lebih agresif
	h.clearCookies(c)

	// Hit CAS logout terlebih dahulu untuk clear server session
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

	// Generate service URL untuk callback
	serviceURL := fmt.Sprintf("%s/api/v1/auth/callback", h.hostUrl)

	// Build CAS login URL dengan renew=true
	loginURL := fmt.Sprintf("%s/login?service=%s&renew=true",
		h.casURL.String(),
		url.QueryEscape(serviceURL))

	log.Printf("Redirecting to CAS login: %s", loginURL)

	c.JSON(http.StatusOK, gin.H{
		"login_url": loginURL,
	})
}

// CallbackHandler - handle callback dari CAS setelah login
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

	// Check user in database
	_, err := h.UserService.CheckUserByUsername(ctx, usernameStr)
	if err != nil {
		log.Printf("User not found: %s, error: %v", usernameStr, err)
		errorURL := fmt.Sprintf("%s/login?error=user_not_found", h.frontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	// Generate JWT token
	tokenString, expirationTime, err := utils.GenerateJWT(usernameStr)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		errorURL := fmt.Sprintf("%s/login?error=token_failed", h.frontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	// Redirect ke frontend dengan token
	callbackURL := fmt.Sprintf("%s/sso/callback?token=%s&expires_at=%d&username=%s",
		h.frontendURL,
		url.QueryEscape(tokenString),
		expirationTime,
		url.QueryEscape(usernameStr))

	log.Printf("Login successful, redirecting to frontend: %s", usernameStr)
	c.Redirect(http.StatusFound, callbackURL)
}

func (h *CASHandler) LogoutHandler(c *gin.Context) {
	log.Printf("Logout handler called")

	// Clear cookies terlebih dahulu
	h.clearCookies(c)

	// Buat HTTP request ke CAS logout untuk memastikan server-side session dihapus
	logoutURL := fmt.Sprintf("%s/logout", h.casURL.String())

	// Hit CAS logout endpoint untuk destroy server session
	client := &http.Client{
		Timeout: 3 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Jangan ikuti redirect
		},
	}

	req, err := http.NewRequest("GET", logoutURL, nil)
	if err == nil {
		// Copy cookies dari request saat ini ke logout request
		for _, cookie := range c.Request.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			log.Printf("CAS server logout completed with status: %d", resp.StatusCode)
		}
	}

	// Redirect ke CAS logout dengan service parameter mengarah ke frontend
	// Ini akan memastikan CAS logout complete, lalu redirect ke frontend
	timestamp := time.Now().UnixNano()
	frontendURL := fmt.Sprintf("%s?logout=true&_=%d", h.frontendURL, timestamp)

	finalLogoutURL := fmt.Sprintf("%s/logout?service=%s",
		h.casURL.String(),
		url.QueryEscape(frontendURL))

	log.Printf("Final redirect to CAS logout with service: %s", finalLogoutURL)

	// Set headers untuk prevent caching
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Clear-Site-Data", "\"cookies\"")

	c.Redirect(http.StatusFound, finalLogoutURL)
}

// clearCookies - helper untuk clear cookies
func (h *CASHandler) clearCookies(c *gin.Context) {
	cookiesToClear := []string{
		"CASTGC",
		"JSESSIONID", // Penting untuk session tracking
		"SESSION",
		"TGC",
		"token",
		"auth",
	}

	// Clear cookies dengan berbagai path dan domain
	paths := []string{"/", "/cas", "/cas/"}
	domains := []string{"", "sso.undiksha.ac.id", ".undiksha.ac.id"}

	for _, cookieName := range cookiesToClear {
		for _, path := range paths {
			for _, domain := range domains {
				c.SetCookie(cookieName, "", -1, path, domain, false, true)
				c.SetCookie(cookieName, "", -1, path, domain, true, true) // secure version
			}
		}
	}

	// Clear semua cookies yang ada di request
	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
		if cookie.Path != "" {
			c.SetCookie(cookie.Name, "", -1, cookie.Path, "", false, true)
		}
	}

	log.Printf("Cleared all cookies including JSESSIONID")
}

// Health check untuk test CAS server
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

// LogoutHandler - handle logout request
// func (h *CASHandler) LogoutHandler(c *gin.Context) {
// 	service := c.Query("service")
// 	log.Printf("Logout handler called")

// 	// Clear cookies terlebih dahulu
// 	h.clearCookies(c)

// 	// Buat HTTP request ke CAS logout untuk memastikan server-side session dihapus
// 	logoutURL := fmt.Sprintf("%s/logout", h.casURL.String())

// 	// Hit CAS logout endpoint untuk destroy server session
// 	client := &http.Client{
// 		Timeout: 3 * time.Second,
// 		CheckRedirect: func(req *http.Request, via []*http.Request) error {
// 			return http.ErrUseLastResponse // Jangan ikuti redirect
// 		},
// 	}

// 	req, err := http.NewRequest("GET", logoutURL, nil)
// 	if err == nil {
// 		// Copy cookies dari request saat ini ke logout request
// 		for _, cookie := range c.Request.Cookies() {
// 			req.AddCookie(cookie)
// 		}

// 		resp, err := client.Do(req)
// 		if err == nil {
// 			resp.Body.Close()
// 			log.Printf("CAS server logout completed with status: %d", resp.StatusCode)
// 		}
// 	}

// 	// Build final logout URL dengan service jika ada
// 	var finalLogoutURL string
// 	if service != "" {
// 		finalLogoutURL = fmt.Sprintf("%s/logout?service=%s",
// 			h.casURL.String(),
// 			url.QueryEscape(service))
// 	} else {
// 		finalLogoutURL = logoutURL
// 	}

// 	log.Printf("Final redirect to CAS logout: %s", finalLogoutURL)

// 	// Set headers untuk prevent caching
// 	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
// 	c.Writer.Header().Set("Pragma", "no-cache")
// 	c.Writer.Header().Set("Expires", "0")
// 	c.Writer.Header().Set("Clear-Site-Data", "\"cookies\"")

// 	c.Redirect(http.StatusFound, finalLogoutURL)
// }
