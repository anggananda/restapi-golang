package middlewares

import (
	"net/http"
	"restapi-golang/mocks"
	"strings"

	"github.com/gin-gonic/gin"
)

func SwaggerMockMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get headers
		userAgent := c.GetHeader("User-Agent")
		referer := c.GetHeader("Referer")
		origin := c.GetHeader("Origin")

		// Deteksi Swagger dari berbagai sumber
		isSwaggerUA := strings.Contains(strings.ToLower(userAgent), "swagger")
		isSwaggerReferer := strings.Contains(strings.ToLower(referer), "swagger") ||
			strings.Contains(strings.ToLower(referer), "/swagger")
		isSwaggerOrigin := strings.Contains(strings.ToLower(origin), "swagger")

		isSwagger := isSwaggerUA || isSwaggerReferer || isSwaggerOrigin

		if isSwagger {
			c.Request.Header.Set("X-Mock-Mode", "true")

		}

		mockMode := c.GetHeader("X-Mock-Mode")

		if mockMode == "true" {
			fullPath := c.FullPath()
			method := c.Request.Method

			if mockResp, ok := mocks.GetMockResponse(fullPath, method); ok {

				c.JSON(http.StatusOK, mockResp)
				c.Abort()
				return
			}
		}

		c.Next()
	}
}
