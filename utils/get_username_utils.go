package utils

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func GetUsername(c *gin.Context) (string, error) {
	val, exists := c.Get("email")
	if !exists {
		return "", errors.New("username not found in context")
	}

	username, ok := val.(string)
	if !ok {
		return "", errors.New("username is not of type string")
	}

	return username, nil
}
