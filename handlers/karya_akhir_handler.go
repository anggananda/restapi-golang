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

type KaryaAkhirHandler struct {
	KaryaAkhirService *services.KaryaAkhirService
}

func NewKaryaAkhirHandler(service *services.KaryaAkhirService) *KaryaAkhirHandler {
	return &KaryaAkhirHandler{
		KaryaAkhirService: service,
	}
}

// GetKaryaAkhirFiltered mendapatkan data karya akhir dengan filter dan pagination
// @Summary      Get filtered karya akhir
// @Description  Mendapatkan data karya akhir berdasarkan filter dengan pagination
// @Tags         Karya Akhir
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.KaryaAkhir}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Router       /karya-akhir [get]
func (h *KaryaAkhirHandler) GetKaryaAkhirFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	search := c.Query("search")
	tahun := utils.StringToInt(c.Query("tahun"), 0)

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	karyaAkhir, total, err := h.KaryaAkhirService.GetKaryaAkhirFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, page, limit)
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
		"datas":  karyaAkhir,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
