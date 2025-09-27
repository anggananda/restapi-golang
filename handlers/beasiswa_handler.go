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

type BeasiswaHandler struct {
	BeasiswaService *services.BeasiswaService
}

func NewBeasiswaHandler(service *services.BeasiswaService) *BeasiswaHandler {
	return &BeasiswaHandler{
		BeasiswaService: service,
	}
}

// GetBeasiswaFiltered mendapatkan data beasiswa dengan filter dan pagination
// @Summary      Get filtered Beasiswa
// @Description  Mendapatkan data beasiswa berdasarkan filter dengan pagination
// @Tags         Beasiswa
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        semester      query     string  false  "Filter berdasarkan semester"
// @Param        jenisBeasiswa      query     string  false  "Filter berdasarkan jenis beasiswa"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Beasiswa}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Router       /beasiswa [get]
func (h *BeasiswaHandler) GetBeasiswaFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	tahun := utils.StringToInt(c.Query("tahun"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	semester := c.Query("semester")
	jenisBeasiswa := c.Query("jenisBeasiswa")
	search := c.Query("search")

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	beasiswa, total, err := h.BeasiswaService.GetBeasiswaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, semester, jenisBeasiswa, search, tahun, page, limit)
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
		"datas":  beasiswa,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})

}
