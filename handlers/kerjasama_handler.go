package handlers

import (
	"context"
	"math"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"time"

	"github.com/gin-gonic/gin"
)

type KerjasamaHandler struct {
	KerjasamaService *services.KerjasamaService
}

func NewKerjasamaHandler(service *services.KerjasamaService) *KerjasamaHandler {
	return &KerjasamaHandler{
		KerjasamaService: service,
	}
}

func (h *KerjasamaHandler) GetKerjasamaFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	kerjasama, total, err := h.KerjasamaService.GetKerjasamaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, search, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pages := int64(0)
	if limit > 0 {
		pages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"datas":  kerjasama,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
