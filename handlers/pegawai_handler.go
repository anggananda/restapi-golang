package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type PegawaiHandler struct {
	PegawaiService *services.PegawaiService
}

func NewPegawaiHandler(service *services.PegawaiService) *PegawaiHandler {
	return &PegawaiHandler{
		PegawaiService: service,
	}
}

// GetDetailPegawai mendapatkan  detail Pegawai
// @Summary      Get detail Pegawai
// @Description  Mendapatkan detail data Pegawai
// @Tags         Pegawai
// @Accept       json
// @Produce      json
// @Param        niu        query     string  false  "niu"
// @Success      200           {object}  models.ListDetailResponse{datas=models.PegawaiHistory}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /pegawai/{niu} [get]
func (h *PegawaiHandler) GetDetailPegawai(c *gin.Context) {
	niu := c.Param("niu")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	pegawai, err := h.PegawaiService.GetDetailPegawai(ctx, niu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": pegawai, "message": "OK"})
}

// GetPegawaiHistoryFiltered mendapatkan  Pegawai  dengan filter dan pagination
// @Summary      Get filtered Pegawai
// @Description  Mendapatkan data Pegawai berdasarkan filter dengan pagination
// @Tags         Pegawai
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        statusPegawai         query     int  false  "Filter berdasarkan statusPegawai"
// @Param        statusKeaktifan         query     int  false  "Filter berdasarkan statusKeaktifan"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.PegawaiHistoryResponse}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /pegawai/history [get]
func (h *PegawaiHandler) GetPegawaiHistoryFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahunStr := c.Query("tahun")
	search := c.Query("search")
	statusPegawai := utils.StringToInt(c.Query("statusPegawai"), 0)
	statusKeaktifan := utils.StringToInt(c.Query("statusKeaktifan"), 0)

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	pegawai, total, err := h.PegawaiService.GetPegawaiHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, statusPegawai, statusKeaktifan, page, limit)
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
		"datas":  pegawai,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *PegawaiHandler) ExportPegawaiCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahunStr := c.Query("tahun")
	search := c.Query("search")
	statusPegawai := utils.StringToInt(c.Query("statusPegawai"), 0)
	statusKeaktifan := utils.StringToInt(c.Query("statusKeaktifan"), 0)

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	pegawai, _, err := h.PegawaiService.GetPegawaiHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, statusPegawai, statusKeaktifan, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"NIP",
		"NIU",
		"Nama",
		"Fakultas",
		"Jurusan",
		"Prodi",
		"Tahun",
		"Status Pegawai",
		"Status Keaktifan",
		"Last Strata",
	}

	var csvData [][]string

	for _, item := range pegawai {
		tahunStr := strconv.Itoa(item.Tahun)

		row := []string{
			item.NIP,
			item.NoIndukUndiksha,
			item.Nama,
			item.Fakultas,
			item.Jurusan,
			item.Prodi,
			tahunStr,
			item.StatusPegawai,
			item.StatusKeaktifan,
			item.LastStrata,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("pegawai_history_%d_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)

}
