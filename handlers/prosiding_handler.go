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

type ProsidingHandler struct {
	ProsidingService *services.ProsidingService
}

func NewProsidingHandler(service *services.ProsidingService) *ProsidingHandler {
	return &ProsidingHandler{
		ProsidingService: service,
	}
}

// GetProsidingFiltered mendapatkan data prosiding dengan filter dan pagination
// @Summary      Get filtered prosiding
// @Description  Mendapatkan data prosiding berdasarkan filter dengan pagination
// @Tags         Prosiding
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        semester         query     string  false  "Filter berdasarkan semester"
// @Param        indexer         query     string  false  "Filter berdasarkan indexer"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Prosiding}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Router       /prosiding [get]
func (h *ProsidingHandler) GetProsidingFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	indexer := c.Query("indexer")
	search := c.Query("search")

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	prosiding, total, err := h.ProsidingService.GetProsidingFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search, page, limit)
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
		"datas":  prosiding,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
