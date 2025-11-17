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

type PenelitianHandler struct {
	PenelitianService *services.PenelitianService
}

func NewPenelitianHandler(service *services.PenelitianService) *PenelitianHandler {
	return &PenelitianHandler{
		PenelitianService: service,
	}
}

// GetPenelitianFiltered mendapatkan data penelitian dengan filter dan pagination
// @Summary      Get filtered Penelitian
// @Description  Mendapatkan data penelitian berdasarkan filter dengan pagination
// @Tags         Penelitian
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        semester         query     string  false  "Filter berdasarkan semester"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Penelitian}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /penelitian [get]
func (h *PenelitianHandler) GetPenelitianFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	penelitian, total, err := h.PenelitianService.GetPenelitianFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
		"datas":  penelitian,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
