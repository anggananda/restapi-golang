// // func (h *CASHandler) LoginHandler(c *gin.Context) {
// // 	// service adalah URL callback aplikasi kamu
// // 	serviceURL := "http://localhost:8080/api/v1/auth/callback"

// // 	// bikin URL CAS login dengan parameter service
// // 	loginURL := fmt.Sprintf("%s/login?service=%s", h.casURL.String(), url.QueryEscape(serviceURL))

// // 	// redirect user ke CAS login
// // 	c.Redirect(http.StatusFound, loginURL)
// // }

// package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"net/url"
// 	"restapi-golang/utils"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"gopkg.in/cas.v2"
// )

// type CASHandler struct {
// 	casClient *cas.Client
// 	casURL    *url.URL
// }

// func NewCASHandler(casURL string) (*CASHandler, error) {
// 	parsedURL, err := url.Parse(casURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing CAS URL: %v", err)
// 	}

// 	client := cas.NewClient(&cas.Options{
// 		URL: parsedURL,
// 	})

// 	return &CASHandler{
// 		casClient: client,
// 		casURL:    parsedURL,
// 	}, nil
// }

// func (h *CASHandler) CASMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Buat custom handler yang menangani CAS authentication
// 		handler := h.casClient.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// Set user info to gin context if authenticated
// 			if username := cas.Username(r); username != "" {
// 				c.Set("cas_username", username)
// 				c.Set("cas_attributes", cas.Attributes(r))
// 			}

// 			// Simpan Writer dan Request untuk digunakan nanti
// 			c.Set("cas_writer", w)
// 			c.Set("cas_request", r)

// 			c.Next()
// 		}))

// 		// Execute CAS handler
// 		handler.ServeHTTP(c.Writer, c.Request)
// 	}
// }

// func (h *CASHandler) LoginHandler(c *gin.Context) {
// 	serviceURL := "http://localhost:8080/api/v1/auth/callback"

// 	loginURL := fmt.Sprintf("%s/login?service=%s", h.casURL.String(), url.QueryEscape(serviceURL))

// 	c.JSON(http.StatusOK, gin.H{
// 		"login_url": loginURL,
// 	})
// }

// func (h *CASHandler) CallbackHandler(c *gin.Context) {
// 	username, exists := c.Get("cas_username")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Authentication required",
// 		})
// 		return
// 	}

// 	// Generate JWT token
// 	usernameStr := username.(string)
// 	tokenString, expirationTime, err := utils.GenerateJWT(usernameStr)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Failed to generate token",
// 		})
// 		return
// 	}

// 	attributes, _ := c.Get("cas_attributes")

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"message": "Login successful",
// 		"data": gin.H{
// 			"user": gin.H{
// 				"username":  usernameStr,
// 				"email":     h.getAttributeValue(attributes.(cas.UserAttributes), []string{"email", "mail", "userEmail"}),
// 				"firstName": h.getAttributeValue(attributes.(cas.UserAttributes), []string{"firstName", "givenName", "first_name"}),
// 				"lastName":  h.getAttributeValue(attributes.(cas.UserAttributes), []string{"lastName", "surname", "last_name"}),
// 			},
// 			"token": gin.H{
// 				"access_token": tokenString,
// 				"expires_at":   expirationTime,
// 				"token_type":   "Bearer",
// 			},
// 			"raw_attributes": attributes,
// 		},
// 	})
// }

// func (h *CASHandler) ProfileHandler(c *gin.Context) {
// 	username, exists := c.Get("cas_username")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"error": "Authentication required",
// 		})
// 		return
// 	}

// 	attributes, _ := c.Get("cas_attributes")

// 	c.JSON(http.StatusOK, gin.H{
// 		"success": true,
// 		"data": gin.H{
// 			"user": gin.H{
// 				"username":  username,
// 				"email":     h.getAttributeValue(attributes.(cas.UserAttributes), []string{"email", "mail", "userEmail"}),
// 				"firstName": h.getAttributeValue(attributes.(cas.UserAttributes), []string{"firstName", "givenName", "first_name"}),
// 				"lastName":  h.getAttributeValue(attributes.(cas.UserAttributes), []string{"lastName", "surname", "last_name"}),
// 			},
// 			"attributes": attributes,
// 		},
// 	})
// }

// func (h *CASHandler) LogoutHandler(c *gin.Context) {
// 	// Implement CAS logout if needed
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Logout successful",
// 	})
// }

// func (h *CASHandler) getAttributeValue(attributes cas.UserAttributes, possibleKeys []string) string {
// 	for _, key := range possibleKeys {
// 		if values, exists := attributes[key]; exists && len(values) > 0 {
// 			return values[0]
// 		}
// 	}
// 	return ""
// }

// // Health check untuk test CAS server
// func (h *CASHandler) HealthCheckHandler(c *gin.Context) {
// 	casURL := "https://sso.undiksha.ac.id/cas"

// 	resp, err := http.Get(casURL)
// 	if err != nil {
// 		c.JSON(http.StatusServiceUnavailable, gin.H{
// 			"status":  "error",
// 			"message": fmt.Sprintf("CAS server unavailable: %v", err),
// 		})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		c.JSON(http.StatusServiceUnavailable, gin.H{
// 			"status":  "error",
// 			"message": fmt.Sprintf("CAS server returned status: %d", resp.StatusCode),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  "success",
// 		"message": "CAS server is reachable",
// 		"data": gin.H{
// 			"url":          casURL,
// 			"status_code":  resp.StatusCode,
// 			"service_url":  "http://localhost:8080",
// 			"current_time": time.Now(),
// 		},
// 	})
// }

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/cas.v2"
)

type CASHandler struct {
	casClient   *cas.Client
	casURL      *url.URL
	frontendURL string
	UserService *services.UserService
}

func NewCASHandler(casURL, frontendURL string, service *services.UserService) (*CASHandler, error) {
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
		UserService: service,
	}, nil
}

func (h *CASHandler) CASMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Set headers untuk prevent caching
		c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Writer.Header().Set("Pragma", "no-cache")
		c.Writer.Header().Set("Expires", "0")
		c.Writer.Header().Set("Vary", "*")

		// JANGAN clear cookies di middleware - ini menyebabkan loop
		// Clear cookies hanya dilakukan di login/logout handlers

		handler := h.casClient.Handle(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set headers untuk prevent caching
			w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			w.Header().Set("Expires", "0")

			if username := cas.Username(r); username != "" {
				c.Set("cas_username", username)
				c.Set("cas_attributes", cas.Attributes(r))
				log.Printf("CAS User authenticated: %s", username)
			}

			c.Set("cas_writer", w)
			c.Set("cas_request", r)
			c.Next()
		}))

		handler.ServeHTTP(c.Writer, c.Request)
	}
}

// Helper function untuk clear CAS-related cookies
func (h *CASHandler) clearCASCookies(c *gin.Context) {
	cookiesToClear := []string{
		"CASTGC",
		"JSESSIONID",
		"SESSION",
		"cas-session",
		"TGC",
		"TICKET_GRANTING_COOKIE",
	}

	// Fix domain issue - don't use localhost:8080 as domain
	host := c.Request.Host
	domains := []string{""}

	// Only add domain if it's not localhost with port
	if !strings.Contains(host, "localhost:") && !strings.Contains(host, "127.0.0.1:") {
		domains = append(domains, host)
		if strings.Contains(host, ".") {
			parts := strings.Split(host, ".")
			if len(parts) > 1 {
				domain := "." + strings.Join(parts[len(parts)-2:], ".")
				domains = append(domains, domain)
			}
		}
	}

	for _, cookieName := range cookiesToClear {
		for _, domain := range domains {
			c.SetCookie(cookieName, "", -1, "/", domain, false, true)
		}
	}
}

func (h *CASHandler) LoginHandler(c *gin.Context) {
	log.Printf("Login handler called")

	// Step 1: Destroy existing sessions brutal
	h.destroyAllSessions(c)

	// Step 2: Hit CAS logout untuk clear server-side session SEBELUM login
	logoutURL := fmt.Sprintf("%s/logout", h.casURL.String())
	client := &http.Client{
		Timeout: 3 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Don't follow redirects
		},
	}

	req, err := http.NewRequest("GET", logoutURL, nil)
	if err == nil {
		// Copy cookies from current request
		for _, cookie := range c.Request.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			log.Printf("Pre-login CAS logout completed with status: %d", resp.StatusCode)
		}
	}

	// Step 3: Wait untuk server processing
	time.Sleep(200 * time.Millisecond)

	// Step 4: Generate unique service URL dengan timestamp
	timestamp := time.Now().UnixNano()
	serviceURL := fmt.Sprintf("http://localhost:8080/api/v1/auth/callback?_=%d", timestamp)

	// Step 5: Build login URL dengan parameter paling agresif
	// gateway=true: mencegah transparent authentication
	// renew=true: memaksa fresh credentials input
	loginURL := fmt.Sprintf("%s/login?service=%s&gateway=true&renew=true&_=%d",
		h.casURL.String(),
		url.QueryEscape(serviceURL),
		timestamp)

	log.Printf("Generated fresh login URL: %s", loginURL)

	c.JSON(http.StatusOK, gin.H{
		"login_url":       loginURL,
		"session_cleared": true,
		"timestamp":       timestamp,
	})
}

// GetCASURL returns the CAS URL for external access
func (h *CASHandler) GetCASURL() *url.URL {
	return h.casURL
}

func (h *CASHandler) CallbackHandler(c *gin.Context) {
	username, exists := c.Get("cas_username")
	if !exists || username == "" {
		log.Println("ERROR: CAS username not found in context")
		errorURL := fmt.Sprintf("%s/login?error=auth_failed&_=%d", h.frontendURL, time.Now().UnixNano())
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	usernameStr := username.(string)
	log.Println("Processing login for: " + usernameStr)

	// Check user in database
	_, err := h.UserService.CheckUserByUsername(ctx, usernameStr)
	if err != nil {
		log.Printf("User not found in database: %s, error: %v", usernameStr, err)
		errorURL := fmt.Sprintf("%s/login?error=user_not_found&_=%d", h.frontendURL, time.Now().UnixNano())
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	// Generate token dengan user info yang lengkap
	tokenString, expirationTime, err := utils.GenerateJWT(usernameStr)
	if err != nil {
		log.Printf("Token generation failed: %v", err)
		errorURL := fmt.Sprintf("%s/login?error=token_failed&_=%d", h.frontendURL, time.Now().UnixNano())
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	// Redirect dengan parameter yang unique
	callbackURL := fmt.Sprintf("%s/sso/callback?token=%s&expires_at=%d&username=%s&_=%d",
		h.frontendURL,
		url.QueryEscape(tokenString),
		expirationTime,
		url.QueryEscape(usernameStr),
		time.Now().UnixNano())

	log.Printf("Redirecting to frontend for user: %s", usernameStr)
	c.Redirect(http.StatusFound, callbackURL)
}

func (h *CASHandler) LogoutHandler(c *gin.Context) {
	service := c.Query("service")
	log.Printf("Logout handler called with service: %s", service)

	// Clear ALL cookies yang mungkin terkait session secara lebih agresif
	h.clearAllCookies(c)

	// Build CAS logout URL dengan parameter yang tepat
	var logoutURL string
	timestamp := time.Now().UnixNano()

	if service != "" {
		// Gunakan parameter yang benar untuk CAS logout
		logoutURL = fmt.Sprintf("%s/logout?service=%s&_=%d",
			h.casURL.String(),
			url.QueryEscape(service),
			timestamp)
	} else {
		// Logout tanpa service URL
		logoutURL = fmt.Sprintf("%s/logout?_=%d", h.casURL.String(), timestamp)
	}

	log.Printf("Performing CAS logout to: %s", logoutURL)

	// Set additional headers untuk prevent caching pada logout response
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\"")

	c.Redirect(http.StatusFound, logoutURL)
}

// HealthCheckHandler untuk check status CAS handler
func (h *CASHandler) HealthCheckHandler(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")

	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"cas_url":      h.casURL.String(),
		"frontend_url": h.frontendURL,
		"timestamp":    time.Now().Unix(),
		"message":      "CAS Handler is healthy",
	})
}

// Helper function untuk clear semua cookies lebih agresif
func (h *CASHandler) clearAllCookies(c *gin.Context) {
	// CAS-related cookies
	casCookies := []string{
		"CASTGC", "TGC", "JSESSIONID", "SESSION",
		"cas-session", "TICKET_GRANTING_COOKIE",
		"CAS_COOKIE", "cas_cookie",
	}

	// Application cookies
	appCookies := []string{
		"token", "session", "user", "auth", "authorization",
		"access_token", "refresh_token",
	}

	allCookies := append(casCookies, appCookies...)

	// Clear dari request cookies yang ada
	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
		// Only set domain if it's not localhost with port
		if cookie.Domain != "" && !strings.Contains(cookie.Domain, "localhost:") {
			c.SetCookie(cookie.Name, "", -1, cookie.Path, cookie.Domain, false, true)
		}
	}

	// Fix domain issue untuk localhost
	host := c.Request.Host
	domains := []string{""}

	// Only add domain if it's not localhost with port
	if !strings.Contains(host, "localhost:") && !strings.Contains(host, "127.0.0.1:") {
		domains = append(domains, host)
		if strings.Contains(host, ".") {
			parts := strings.Split(host, ".")
			if len(parts) > 1 {
				domains = append(domains, "."+strings.Join(parts[len(parts)-2:], "."))
			}
		}
	}

	paths := []string{"/", "/cas", "/api"}

	for _, cookieName := range allCookies {
		for _, domain := range domains {
			for _, path := range paths {
				c.SetCookie(cookieName, "", -1, path, domain, false, true)
			}
		}
	}
}

// Tambahkan handler untuk force logout dari CAS server
func (h *CASHandler) ForceLogoutHandler(c *gin.Context) {
	// Clear semua cookies
	h.clearAllCookies(c)

	// Redirect ke CAS logout dengan parameter khusus
	timestamp := time.Now().UnixNano()
	service := c.Query("service")

	var logoutURL string
	if service != "" {
		logoutURL = fmt.Sprintf("%s/logout?service=%s&renew=true&_=%d",
			h.casURL.String(),
			url.QueryEscape(service),
			timestamp)
	} else {
		logoutURL = fmt.Sprintf("%s/logout?renew=true&_=%d", h.casURL.String(), timestamp)
	}

	log.Printf("Force logout to: %s", logoutURL)

	// Set headers untuk complete cache clearing
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\", \"executionContexts\"")

	c.Redirect(http.StatusFound, logoutURL)
}

// Tambahkan method ini ke CAS Handler untuk destroy CAS session

// DestroySessionHandler - menghancurkan session CAS dengan cara yang brutal
func (h *CASHandler) DestroySessionHandler(c *gin.Context) {
	log.Printf("Destroying CAS session...")

	// Step 1: Clear ALL cookies termasuk JSESSIONID
	h.destroyAllSessions(c)

	// Step 2: Hit CAS logout endpoint secara langsung untuk destroy server-side session
	service := c.Query("service")
	timestamp := time.Now().UnixNano()

	// Step 3: First, hit CAS logout tanpa service untuk destroy session
	firstLogoutURL := fmt.Sprintf("%s/logout", h.casURL.String())

	log.Printf("Step 1: Destroying CAS session at: %s", firstLogoutURL)

	// Make HTTP request to CAS logout untuk destroy session di server
	client := &http.Client{
		Timeout: 5 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Don't follow redirects
		},
	}

	req, err := http.NewRequest("GET", firstLogoutURL, nil)
	if err == nil {
		// Copy cookies from current request to logout request
		for _, cookie := range c.Request.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			log.Printf("CAS logout response status: %d", resp.StatusCode)
		}
	}

	// Step 4: Wait a moment for server processing
	time.Sleep(100 * time.Millisecond)

	// Step 5: Now redirect to final service if provided
	var finalURL string
	if service != "" {
		finalURL = fmt.Sprintf("%s?session_destroyed=true&_=%d", service, timestamp)
	} else {
		finalURL = fmt.Sprintf("%s/logout-complete?session_destroyed=true&_=%d", h.frontendURL, timestamp)
	}

	log.Printf("Final redirect after session destruction: %s", finalURL)

	// Set headers untuk complete cache clearing
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Clear-Site-Data", "\"cache\", \"cookies\", \"storage\", \"executionContexts\"")

	c.Redirect(http.StatusFound, finalURL)
}

// destroyAllSessions - menghancurkan semua session cookies dengan brutal
func (h *CASHandler) destroyAllSessions(c *gin.Context) {
	// Semua cookies yang mungkin menyimpan session
	sessionCookies := []string{
		"JSESSIONID",
		"CASTGC",
		"TGC",
		"SESSION",
		"cas-session",
		"TICKET_GRANTING_COOKIE",
		"CAS_COOKIE",
		"jsessionid", // lowercase version
		"castgc",     // lowercase version
	}

	// Clear dengan berbagai kombinasi
	paths := []string{"/", "/cas", "/cas/", "/api", "/api/"}

	for _, cookieName := range sessionCookies {
		for _, path := range paths {
			// Clear dengan empty domain (untuk localhost)
			c.SetCookie(cookieName, "", -1, path, "", false, true)
			c.SetCookie(cookieName, "", -1, path, "", true, true) // dengan secure flag

			// Clear juga dengan uppercase/lowercase variations
			c.SetCookie(strings.ToUpper(cookieName), "", -1, path, "", false, true)
			c.SetCookie(strings.ToLower(cookieName), "", -1, path, "", false, true)
		}
	}

	// Clear semua existing cookies dari request
	for _, cookie := range c.Request.Cookies() {
		c.SetCookie(cookie.Name, "", -1, "/", "", false, true)
		c.SetCookie(cookie.Name, "", -1, cookie.Path, "", false, true)
	}

	log.Printf("Destroyed all session cookies")
}

// PreLogoutHandler - dipanggil sebelum redirect ke CAS login untuk clear session
func (h *CASHandler) PreLogoutHandler(c *gin.Context) {
	log.Printf("Pre-logout: Clearing all sessions before new login")

	// Destroy semua session
	h.destroyAllSessions(c)

	// Hit CAS logout endpoint untuk clear server-side session
	logoutURL := fmt.Sprintf("%s/logout", h.casURL.String())

	client := &http.Client{
		Timeout: 3 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest("GET", logoutURL, nil)
	if err == nil {
		// Copy existing cookies untuk ensure proper logout
		for _, cookie := range c.Request.Cookies() {
			req.AddCookie(cookie)
		}

		resp, err := client.Do(req)
		if err == nil {
			resp.Body.Close()
			log.Printf("Pre-logout CAS hit completed with status: %d", resp.StatusCode)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Session cleared, ready for fresh login",
		"timestamp": time.Now().Unix(),
	})
}

// func (h *CASHandler) LogoutHandler(c *gin.Context) {
// 	fmt.Println("🔴 LOGOUT: Clearing all auth data")

// 	// 1. HAPUS SEMUA COOKIE di response
// 	c.SetCookie("auth_token", "", -1, "/", "", false, true)
// 	c.SetCookie("session_token", "", -1, "/", "", false, true)
// 	c.SetCookie("cas_session", "", -1, "/", "", false, true)

// 	// 2. HAPUS HEADER AUTHORIZATION
// 	c.Header("Authorization", "")

// 	logoutCallbackURL := h.frontendURL + "/logout-complete"
// 	// 3. KIRIM RESPONSE JSON BUKAN REDIRECT
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Logout successful",
// 		"logout_url": fmt.Sprintf("%s/logout?service=%s",
// 			h.casURL.String(),
// 			url.QueryEscape(logoutCallbackURL)),
// 	})
// }

func (h *CASHandler) ProfileHandler(c *gin.Context) {
	username, exists := c.Get("cas_username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authentication required",
		})
		return
	}

	attributes, _ := c.Get("cas_attributes")

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": gin.H{
			"user": gin.H{
				"username":  username,
				"email":     h.getAttributeValue(attributes.(cas.UserAttributes), []string{"email", "mail", "userEmail"}),
				"firstName": h.getAttributeValue(attributes.(cas.UserAttributes), []string{"firstName", "givenName", "first_name"}),
				"lastName":  h.getAttributeValue(attributes.(cas.UserAttributes), []string{"lastName", "surname", "last_name"}),
			},
			"attributes": attributes,
		},
	})
}

func (h *CASHandler) getAttributeValue(attributes cas.UserAttributes, possibleKeys []string) string {
	for _, key := range possibleKeys {
		if values, exists := attributes[key]; exists && len(values) > 0 {
			return values[0]
		}
	}
	return ""
}

// Health check untuk test CAS server
// func (h *CASHandler) HealthCheckHandler(c *gin.Context) {
// 	casURL := "https://sso.undiksha.ac.id/cas"

// 	resp, err := http.Get(casURL)
// 	if err != nil {
// 		c.JSON(http.StatusServiceUnavailable, gin.H{
// 			"status":  "error",
// 			"message": fmt.Sprintf("CAS server unavailable: %v", err),
// 		})
// 		return
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		c.JSON(http.StatusServiceUnavailable, gin.H{
// 			"status":  "error",
// 			"message": fmt.Sprintf("CAS server returned status: %d", resp.StatusCode),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  "success",
// 		"message": "CAS server is reachable",
// 		"data": gin.H{
// 			"url":          casURL,
// 			"status_code":  resp.StatusCode,
// 			"service_url":  "http://localhost:8080",
// 			"current_time": time.Now(),
// 		},
// 	})
// }
