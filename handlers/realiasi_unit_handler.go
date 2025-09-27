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

type RealisasiUnitHandler struct {
	RealisasiUnitService *services.RealisasiUnitService
}

func NewRealisasiUnitHandler(service *services.RealisasiUnitService) *RealisasiUnitHandler {
	return &RealisasiUnitHandler{
		RealisasiUnitService: service,
	}
}

// GetRealisasiUnitFiltered mendapatkan data ralisasi unit dengan filter dan pagination
// @Summary      Get filtered ralisasi unit
// @Description  Mendapatkan data ralisasi unit berdasarkan filter dengan pagination
// @Tags         Realisasi Unit
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.RealisasiUnit}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Router       /realisasi-unit [get]
func (h *RealisasiUnitHandler) GetRealisasiUnitFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	search := c.Query("search")

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	realisasiUnit, total, err := h.RealisasiUnitService.GetRealisasiUnitFiltered(ctx, search, page, limit)
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
		"datas":  realisasiUnit,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
