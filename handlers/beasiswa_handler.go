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
// @Security     BearerAuth
// @Router       /beasiswa [get]
func (h *BeasiswaHandler) GetBeasiswaFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	tahunStr := c.Query("tahun")
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	semester := c.Query("semester")
	jenisBeasiswa := c.Query("jenisBeasiswa")
	search := c.Query("search")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

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

func (h *BeasiswaHandler) ExportBeasiswaCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	tahunStr := c.Query("tahun")
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	semester := c.Query("semester")
	jenisBeasiswa := c.Query("jenisBeasiswa")
	search := c.Query("search")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	beasiswa, _, err := h.BeasiswaService.GetBeasiswaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, semester, jenisBeasiswa, search, tahun, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID (_id)",
		"NIM",
		"Nama",
		"Jenis Beasiswa",
		"IPK",
		"Status",
		"Tahun",
		"Semester",
		"Semester Type",
		"Periode",
		"Unit UK Kode",
		"Unit Fakultas Kode",
		"Unit Jurusan Kode",
		"Unit Prodi Kode",
		"Fakultas",
		"Jurusan",
		"Program Studi",
	}

	var csvData [][]string

	for _, item := range beasiswa {

		idStr := strconv.Itoa(item.ID)
		tahunStr := strconv.Itoa(item.Tahun)
		ipkStr := strconv.Itoa(int(item.IPK))

		row := []string{
			idStr,
			item.NIM,
			item.Nama,
			item.JenisBeasiswa,
			ipkStr,
			item.Status,
			tahunStr,
			item.Semester,
			item.SemesterType,
			item.Unit.UKKode,
			item.Unit.FktKode,
			item.Unit.JrsKose,
			item.Unit.PrdKode,
			item.Unit.Fakultas,
			item.Unit.Jurusan,
			item.Unit.Prodi,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("beasiswa_%d_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)

}
