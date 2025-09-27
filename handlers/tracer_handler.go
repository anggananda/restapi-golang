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

type TracerHandler struct {
	TracerService *services.TracerService
}

func NewTracerHandler(service *services.TracerService) *TracerHandler {
	return &TracerHandler{
		TracerService: service,
	}
}

// GetTracerFiltered mendapatkan data tracer dengan filter dan pagination
// @Summary      Get filtered tracer
// @Description  Mendapatkan data tracer berdasarkan filter dengan pagination
// @Tags         Tracer
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        bulan         query     int  false  "Filter berdasarkan bulan"
// @Param        statusTracer         query     string  false  "Filter berdasarkan statusTracer"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Tracer}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Router       /tracer [get]
func (h *TracerHandler) GetTracerFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	tahun := utils.StringToInt(c.Query("tahun"), 0)
	bulan := utils.StringToInt(c.Query("bulan"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	statusTracer := c.Query("statusTracer")
	search := c.Query("search")

	ctx, cancle := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancle()

	tracer, total, err := h.TracerService.GetTracerFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search, tahun, bulan, page, limit)
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
		"datas":  tracer,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
