package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type AgendaMengajarHandler struct {
	AgendaMengajarService *services.AgendaMengajarService
}

func NewAgendaMengajarHandler(service *services.AgendaMengajarService) *AgendaMengajarHandler {
	return &AgendaMengajarHandler{
		AgendaMengajarService: service,
	}
}

// GetAgendaMengajarFiltered mendapatkan data agenda mengajar dengan filter dan pagination
// @Summary      Get filtered agenda mengajar
// @Description  Mendapatkan data agenda mengajar berdasarkan filter dengan pagination
// @Tags         Agenda Mengajar
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        semester      query     string  false  "Filter berdasarkan semester"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.AgendaMengajar}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /agenda-mengajar [get]
func (h *AgendaMengajarHandler) GetAgendaMengajarFiltered(c *gin.Context) {
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

	agendaMengajar, total, err := h.AgendaMengajarService.GetAgendaMengajarFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
		"datas":  agendaMengajar,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *AgendaMengajarHandler) ExportAgendaMengajarCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	agendaMengajar, _, err := h.AgendaMengajarService.GetAgendaMengajarFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID (_id)",
		"Dosen",
		"ID Kelas",
		"ID Penawaran",
		"Jenis Kelas",
		"Kode MK",
		"Kurikulum",
		"Mata Kuliah",
		"NIP Dosen",
		"Periode",
		"Pertemuan",
		"Semester",
		"Sumber",
		"Tahun",
		"Unit UK Kode",
		"Unit Fakultas Kode",
		"Unit Jurusan Kode",
		"Unit Prodi Kode",
		"Unit Fakultas Nama",
		"Unit Jurusan Nama",
		"Unit Prodi Nama",
	}

	var csvData [][]string

	for _, item := range agendaMengajar {
		dosenList := strings.Join(item.Dosen, ", ")
		nipList := strings.Join(item.NipDosen, ", ")

		idStr := strconv.Itoa(item.ID)

		row := []string{
			idStr,
			dosenList,
			item.IdKelas,
			item.IdPenawaran,
			item.JenisKelas,
			item.Kode,
			item.Kurikulum,
			item.Matakuliah,
			nipList,
			item.Periode,
			item.Pertemuan,
			item.Semester,
			item.Sumber,
			item.Tahun,
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

	filename := fmt.Sprintf("agenda_mengajar_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)

}
