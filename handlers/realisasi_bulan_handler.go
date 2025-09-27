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

type RealisasiBulanHandler struct {
	RealisasiBulanService *services.RealisasiBulanService
}

func NewRealisasiBulanHandler(service *services.RealisasiBulanService) *RealisasiBulanHandler {
	return &RealisasiBulanHandler{
		RealisasiBulanService: service,
	}
}

// GetRealisasiBulanFiltered mendapatkan data realisasi bulan dengan filter dan pagination
// @Summary      Get filtered realisasi bulan
// @Description  Mendapatkan data realisasi bulan berdasarkan filter dengan pagination
// @Tags         Realisasi Bulan
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.RealisasiBulan}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Router       /realisasi-bulan [get]
func (h *RealisasiBulanHandler) GetRealisasiBulanFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	tahun := c.Query("tahun")
	search := c.Query("search")

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	realisasiBulan, total, err := h.RealisasiBulanService.GetRealisasiBulanFiltered(ctx, tahun, search, page, limit)
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
		"datas":  realisasiBulan,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
