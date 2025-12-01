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

type AngketMhsHandler struct {
	AngketMhsService *services.AngketMhsService
}

func NewAngketMhsHandler(service *services.AngketMhsService) *AngketMhsHandler {
	return &AngketMhsHandler{
		AngketMhsService: service,
	}
}

// GetAngketMhsFiltered mendapatkan data angket mahasiswa dengan filter dan pagination
// @Summary      Get filtered Angket Mahasiswa
// @Description  Mendapatkan data angket mahasiswa berdasarkan filter dengan pagination
// @Tags         Angket Mahasiswa
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
// @Success      200           {object}  models.ListResponse{datas=[]models.AngketMhs}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /angket-mhs [get]
func (h *AngketMhsHandler) GetAngketMhsFiltered(c *gin.Context) {
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

	angketMhs, total, err := h.AngketMhsService.GetAngketMhsFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
		"datas":  angketMhs,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

// ExportAngketMhsCSV mengekspor data angket mahasiswa ke format CSV
// @Summary      Export Angket Mahasiswa ke CSV
// @Description  Mengekspor daftar angket mahasiswa yang telah difilter ke dalam file CSV.
// @Tags         Angket Mahasiswa
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        kodeFakultas    query    string false    "Filter berdasarkan Kode Fakultas"
// @Param        kodeJurusan     query    string false    "Filter berdasarkan Kode Jurusan"
// @Param        kodeProdi       query    string false    "Filter berdasarkan Kode Program Studi"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        semester        query    string false    "Filter berdasarkan Semester "
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /angket-mhs/export-csv [get]
func (h *AngketMhsHandler) ExportAngketMhsCSV(c *gin.Context) {
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

	angketMhs, _, err := h.AngketMhsService.GetAngketMhsFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID (_id)",
		"Mata Kuliah",
		"Kode MK",
		"ID Kelas",
		"ID Penawaran",
		"Dosen",
		"NIP Dosen",
		"Periode",
		"Semester",
		"Tahun",
		"Unit UK Kode",
		"Unit Fakultas Kode",
		"Unit Jurusan Kode",
		"Unit Prodi Kode",
		"Fakultas",
		"Jurusan",
		"Program Studi",
	}

	var csvData [][]string

	for _, item := range angketMhs {
		dosenList := strings.Join(item.Dosen, ", ")
		nipList := strings.Join(item.NipDosen, ", ")

		idStr := strconv.Itoa(item.ID)

		row := []string{
			idStr,
			item.Mk,
			item.Kode,
			item.IdKelas,
			item.IdPenawaran,
			dosenList,
			nipList,
			item.Periode,
			item.Semester,
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
	filename := fmt.Sprintf("angket_mahasiswa_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
