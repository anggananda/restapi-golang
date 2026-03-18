package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vmihailenco/msgpack/v5"
)

const MsgPackContentType = "application/x-msgpack"

func Render(c *gin.Context, code int, data any) {
	b, err := msgpack.Marshal(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to marshal msgpack"})
		return
	}
	c.Data(code, MsgPackContentType, b)
}
