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

type DosenHandler struct {
	DosenService *services.DosenService
}

func NewDosenHandler(service *services.DosenService) *DosenHandler {
	return &DosenHandler{
		DosenService: service,
	}
}

// GetDetailDosen mendapatkan  detail dosen
// @Summary      Get detail dosen
// @Description  Mendapatkan detail data dosen
// @Tags         Dosen
// @Accept       json
// @Produce      json
// @Param        niu        query     string  false  "niu"
// @Success      200           {object}  models.ListDetailResponse{datas=models.Dosen}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dosen/{niu} [get]
func (h *DosenHandler) GetDetailDosen(c *gin.Context) {
	niu := c.Param("niu")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	dosen, err := h.DosenService.GetDetailDosen(ctx, niu)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"datas": dosen, "message": "OK"})
}

// GetDosenHistoryFiltered mendapatkan  dosen  dengan filter dan pagination
// @Summary      Get filtered dosen
// @Description  Mendapatkan data dosen berdasarkan filter dengan pagination
// @Tags         Dosen
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
// @Success      200           {object}  models.ListResponse{datas=[]models.DosenHistoryResponse}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /dosen/history [get]
func (h *DosenHandler) GetDosenHistoryFiltered(c *gin.Context) {
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

	results, total, err := h.DosenService.GetDosenHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, statusPegawai, statusKeaktifan, page, limit)
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
		"datas":  results,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *DosenHandler) ExportDosenCSV(c *gin.Context) {
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

	dosen, _, err := h.DosenService.GetDosenHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, statusPegawai, statusKeaktifan, 1, limit)
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
		"Jabatan Fungsional",
		"Strata",
	}

	var csvData [][]string

	for _, item := range dosen {
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
			item.JabatanFungsional,
			item.Strata,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("dosen_history_%d_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)

}
