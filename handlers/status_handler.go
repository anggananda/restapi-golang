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

type StatusHandler struct {
	StatusService *services.StatusService
	RDB           *redis.Client
}

func NewStatusHandler(service *services.StatusService, rdb *redis.Client) *StatusHandler {
	return &StatusHandler{
		StatusService: service,
		RDB:           rdb,
	}
}

// GetStatusMahasiswa mendapatkan  status mahasiswa
// @Summary      Get status mahasiswa
// @Description  Mendapatkan data status mahasiswa
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200           {object}  models.ListDetailResponse{datas=[]models.Status}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /status-mhs [get]
func (h *StatusHandler) GetStatusMahasiswa(c *gin.Context) {
	ctx := context.Background()
	key := "status_mahasiswa"

	cached, err := h.RDB.Get(ctx, key).Bytes()
	if err == nil {
		var data []models.Status
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

	mongoData, svcErr := h.StatusService.GetStatusMahasiswa(ctx)
	if svcErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": svcErr.Error()})
		return
	}

	if len(mongoData) == 0 {
		c.JSON(500, gin.H{
			"error": "MongoDB returned empty data",
		})
		return
	}

	jsonBytes, _ := json.Marshal(mongoData)
	h.RDB.Set(ctx, key, jsonBytes, 0)

	c.JSON(http.StatusOK, gin.H{
		"datas":   mongoData,
		"message": "OK (from MongoDB, cached)",
	})
}

// GetStatusPegawai mendapatkan  status pegawai
// @Summary      Get status pegawai
// @Description  Mendapatkan data status pegawai
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200           {object}  models.ListDetailResponse{datas=[]models.Status}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /status-pegawai [get]
func (h *StatusHandler) GetStatusPegawai(c *gin.Context) {
	ctx := context.Background()
	key := "status_pegawai"

	cached, err := h.RDB.Get(ctx, key).Bytes()
	if err == nil {
		var data []models.Status
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

	mongoData, svcErr := h.StatusService.GetStatusPegawai(ctx)
	if svcErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": svcErr.Error()})
		return
	}

	if len(mongoData) == 0 {
		c.JSON(500, gin.H{
			"error": "MongoDB returned empty data",
		})
		return
	}

	jsonBytes, _ := json.Marshal(mongoData)
	h.RDB.Set(ctx, key, jsonBytes, 0)

	c.JSON(http.StatusOK, gin.H{
		"datas":   mongoData,
		"message": "OK (from MongoDB, cached)",
	})
}

// GetStatusKeaktifanPegawai mendapatkan  status keaktifan pegawai
// @Summary      Get status keaktifan pegawai
// @Description  Mendapatkan data status keaktifan pegawai
// @Tags         Status
// @Accept       json
// @Produce      json
// @Success      200           {object}  models.ListDetailResponse{datas=[]models.Status}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /status-keaktifan-pegawai [get]
func (h *StatusHandler) GetStatusKeaktifanPegawai(c *gin.Context) {
	ctx := context.Background()
	key := "status_keaktifan_pegawai"

	cached, err := h.RDB.Get(ctx, key).Bytes()
	if err == nil {
		var data []models.Status
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

	mongoData, svcErr := h.StatusService.GetStatusKeaktifanPegawai(ctx)
	if svcErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": svcErr.Error()})
		return
	}

	if len(mongoData) == 0 {
		c.JSON(500, gin.H{
			"error": "MongoDB returned empty data",
		})
		return
	}

	jsonBytes, _ := json.Marshal(mongoData)
	h.RDB.Set(ctx, key, jsonBytes, 0)

	c.JSON(http.StatusOK, gin.H{
		"datas":   mongoData,
		"message": "OK (from MongoDB, cached)",
	})
}
