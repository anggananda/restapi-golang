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

type MhsWisudaHandler struct {
	MhsWisudaService *services.MhsWisudaService
}

func NewMhsWisudaHandler(service *services.MhsWisudaService) *MhsWisudaHandler {
	return &MhsWisudaHandler{
		MhsWisudaService: service,
	}
}

// GetMhsWisudaFiltered mendapatkan  mahasiswa wisuda dengan filter dan pagination
// @Summary      Get filtered mahasiswa wisuda
// @Description  Mendapatkan data mahasiswa wisuda berdasarkan filter dengan pagination
// @Tags         Mahasiswa Wisuda
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        semester         query     int  false  "Filter berdasarkan semester"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.MhsWisuda}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /mhs-wisuda [get]
func (h *MhsWisudaHandler) GetMhsWisudaFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	search := c.Query("search")
	tahunStr := c.Query("tahun")
	bulan := utils.StringToInt(c.Query("bulan"), 0)

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	mhsWisuda, total, err := h.MhsWisudaService.GetMhsWisudaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, bulan, page, limit)
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
		"datas":  mhsWisuda,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}
