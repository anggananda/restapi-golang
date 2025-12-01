package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type PerpemHandler struct {
	PerpemService *services.PerpemService
}

func NewPerpemHandler(service *services.PerpemService) *PerpemHandler {
	return &PerpemHandler{PerpemService: service}
}

// GetPerpemFiltered mendapatkan data perangkat pembelajaran dengan filter dan pagination
// @Summary      Get filtered perangkat pembelajaran
// @Description  Mendapatkan data perangkat pembelajaran berdasarkan filter dengan pagination
// @Tags         Perangkat Pembelajaran
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
// @Success      200           {object}  models.ListResponse{datas=[]models.Perpem}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /perpem [get]
func (h *PerpemHandler) GetPerpemFiltered(c *gin.Context) {
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

	perpem, total, err := h.PerpemService.GetPerpemFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
		"datas":  perpem,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

// ExportPerpemCSV mengekspor data perpem ke format CSV
// @Summary      Export Perpem ke CSV
// @Description  Mengekspor daftar perpem yang telah difilter ke dalam file CSV.
// @Tags         Perangkat Pembelajaran
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        kodeFakultas    query    string false    "Filter berdasarkan Kode Fakultas"
// @Param        kodeJurusan     query    string false    "Filter berdasarkan Kode Jurusan"
// @Param        kodeProdi       query    string false    "Filter berdasarkan Kode Program Studi"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        semester          query    string false    "Filter berdasarkan semester"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /perpem/export-csv [get]
func (h *PerpemHandler) ExportPerpemCSV(c *gin.Context) {
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

	perpem, _, err := h.PerpemService.GetPerpemFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID Penawaran", "Kode Mata Kuliah", "ID Kelas", "Mata Kuliah",
		"Kurikulum", "Pertemuan Ke-", "Metode Pembelajaran",

		"Dosen Pengampu (Gabungan)",

		"Semester", "Tahun",

		"File Silabus", "File Kontrak Kuliah", "File RPS", "File RTM",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range perpem {
		dosenList := strings.Join(item.Dosen, " | ")
		row := []string{

			item.IdPenarawan,
			item.Kode,
			item.IdKelas,
			item.MK,
			item.Kurikulum,
			item.Pertemuan,
			item.Metode,

			dosenList,

			item.Semester,
			item.Tahun,

			item.Silabus,
			item.Kontrak,
			item.Rps,
			item.Rtm,

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
	filename := fmt.Sprintf("perangkat_pembelajaran_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
