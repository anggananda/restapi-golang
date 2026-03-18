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
// @Security     BearerAuth
// @Router       /karya-akhir [get]
func (h *KaryaAkhirHandler) GetKaryaAkhirFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	search := c.Query("search")
	tahunStr := c.Query("tahun")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	karyaAkhir, total, err := h.KaryaAkhirService.GetKaryaAkhirFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pages := int64(0)
	if limit > 0 {
		pages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{
			"status": "success",
			"datas":  karyaAkhir,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	} else {
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
}

// ExportKaryaAkhirCSV mengekspor data karya akhir ke format CSV
// @Summary      Export Karya Akhir ke CSV
// @Description  Mengekspor daftar karya akhir yang telah difilter ke dalam file CSV.
// @Tags         Karya Akhir
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        kodeFakultas    query    string false    "Filter berdasarkan Kode Fakultas"
// @Param        kodeJurusan     query    string false    "Filter berdasarkan Kode Jurusan"
// @Param        kodeProdi       query    string false    "Filter berdasarkan Kode Program Studi"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /karya-akhir/export-csv [get]
func (h *KaryaAkhirHandler) ExportKaryaAkhirCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	search := c.Query("search")
	tahunStr := c.Query("tahun")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	karyaAkhir, _, err := h.KaryaAkhirService.GetKaryaAkhirFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "NIM", "Nama Lengkap Mahasiswa", "Tahun Masuk", "Judul Karya Akhir",
		"Current State", "Main Stage", "Status Judul", "Nama PA", "Nilai Akhir",

		"Pembimbing 1", "Pembimbing 2", "Pembimbing 3",

		"Penguji 1", "Penguji 2", "Penguji 3", "Penguji 4", "Penguji 5", "Penguji 6",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range karyaAkhir {

		idStr := strconv.Itoa(item.ID)
		tahunMasukStr := strconv.Itoa(item.TahunMasuk)

		row := []string{

			idStr,
			item.NIM,
			item.NamaLengkap,
			tahunMasukStr,
			item.Judul,
			item.CurrentState,
			item.MainStage,
			item.StatusJudul,
			item.NamaPA,
			item.NilaiAkhir,

			item.NamaPembimbing1,
			item.NamaPembimbing2,
			item.NamaPembimbing3,

			item.NamaPenguji1,
			item.NamaPenguji2,
			item.NamaPenguji3,
			item.NamaPenguji4,
			item.NamaPenguji5,
			item.NamaPenguji6,

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
	filename := fmt.Sprintf("karya_akhi_%d_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
