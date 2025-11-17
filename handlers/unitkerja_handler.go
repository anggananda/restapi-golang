package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"restapi-golang/models"
	"restapi-golang/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type UnitKerjaHandler struct {
	UnitKerjaService *services.UnitKerjaService
	RDB              *redis.Client
}

func NewUnitKerjaHandler(service *services.UnitKerjaService, rdb *redis.Client) *UnitKerjaHandler {
	return &UnitKerjaHandler{
		UnitKerjaService: service,
		RDB:              rdb,
	}
}

func (h *UnitKerjaHandler) GetUnitKerja(c *gin.Context) {
	ctx := context.Background()
	key := "unit_kerja"

	cached, err := h.RDB.Get(ctx, key).Bytes()
	if err == nil {
		var data []models.Data
		if err := json.Unmarshal(cached, &data); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"datas":   data,
				"message": "OK (from Redis)",
			})
			return
		}

		h.RDB.Del(ctx, key)
	}

	if err != nil && err != redis.Nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mongoData, svcErr := h.UnitKerjaService.GetUnitKerja(ctx)
	if svcErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": svcErr.Error()})
		return
	}

	if len(mongoData.Data) == 0 {
		c.JSON(500, gin.H{
			"error": "MongoDB returned empty data",
		})
		return
	}

	jsonBytes, _ := json.Marshal(mongoData.Data)
	h.RDB.Set(ctx, key, jsonBytes, 0)

	c.JSON(http.StatusOK, gin.H{
		"datas":   mongoData.Data,
		"message": "OK (from MongoDB, cached)",
	})
}
