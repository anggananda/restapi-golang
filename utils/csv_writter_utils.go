package utils

import (
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendCSV(c *gin.Context, filename string, headers []string, records [][]string) {
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.csv", filename))
	c.Header("Content-Type", "text/csv")
	c.Header("Content-Transfer-Encoding", "binary")

	writer := csv.NewWriter(c.Writer)

	if err := writer.Write(headers); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menulis header CSV"})
		return
	}

	if err := writer.WriteAll(records); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menulis data CSV"})
		return
	}

	writer.Flush()
}
