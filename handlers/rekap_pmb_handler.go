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

type RekapPMBHandler struct {
	RekapPMBService *services.RekapPMBService
}

func NewRekapPMBHandler(service *services.RekapPMBService) *RekapPMBHandler {
	return &RekapPMBHandler{
		RekapPMBService: service,
	}
}

// GetRekapPMBFiltered mendapatkan data rekap pmb dengan filter dan pagination
// @Summary      Get filtered rekap pmb
// @Description  Mendapatkan data rekap pmb berdasarkan filter dengan pagination
// @Tags         Rekap PMB
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.RekapPMB}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /rekap-pmb [get]
func (h *RekapPMBHandler) GetRekapPMBFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahunStr := c.Query("tahun")
	search := c.Query("search")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	rekapPMB, total, err := h.RekapPMBService.GetRekapPMBFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, page, limit)
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
		"datas":  rekapPMB,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

// ExportRekapPMBCSV mengekspor data rekap pmb ke format CSV
// @Summary      Export Rekap PMB ke CSV
// @Description  Mengekspor daftar rekap pmb yang telah difilter ke dalam file CSV.
// @Tags         Rekap PMB
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
// @Router       /rekap-pmb/export-csv [get]
func (h *RekapPMBHandler) ExportRekapPMBCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahunStr := c.Query("tahun")
	search := c.Query("search")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	rekapPMB, _, err := h.RekapPMBService.GetRekapPMBFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, search, tahun, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{

		"ID", "Tahun", "Kode Prodi", "Nama Prodi",

		"SNBP Peminat", "SNBP Lulus", "SNBP Daftar",

		"SNBT Peminat", "SNBT Lulus", "SNBT Daftar",

		"SMBJM CBT Peminat", "SMBJM CBT Lulus", "SMBJM CBT Daftar",

		"SMBJM Raport Peminat", "SMBJM Raport Lulus", "SMBJM Raport Daftar",

		"SMBJM Talent Peminat", "SMBJM Talent Lulus", "SMBJM Talent Daftar",

		"SMBJM UTBK Peminat", "SMBJM UTBK Lulus", "SMBJM UTBK Daftar",

		"Profesi Peminat", "Profesi Lulus", "Profesi Daftar",

		"Internasional Peminat", "Internasional Lulus", "Internasional Daftar",

		"Pascasarjana Peminat", "Pascasarjana Lulus", "Pascasarjana Daftar",

		"ADIK Papua Peminat", "ADIK Papua Lulus", "ADIK Papua Daftar",

		"Jumlah Total Peminat", "Jumlah Total Lulus", "Jumlah Total Daftar",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range rekapPMB {

		idStr := strconv.Itoa(item.ID)
		tahunStr := strconv.Itoa(item.Tahun)

		row := []string{

			idStr, tahunStr, item.Kode, item.NamaProdi,

			item.SNBP.Peminat, item.SNBP.Lulus, item.SNBP.Daftar,

			item.SNBT.Peminat, item.SNBT.Lulus, item.SNBT.Daftar,

			item.SMBJM_CBT.Peminat, item.SMBJM_CBT.Lulus, item.SMBJM_CBT.Daftar,

			item.SMBJM_Rpt.Peminat, item.SMBJM_Rpt.Lulus, item.SMBJM_Rpt.Daftar,

			item.SMBJM_Tlt.Peminat, item.SMBJM_Tlt.Lulus, item.SMBJM_Tlt.Daftar,

			item.SMBJM_UTBK.Peminat, item.SMBJM_UTBK.Lulus, item.SMBJM_UTBK.Daftar,

			item.Profesi.Peminat, item.Profesi.Lulus, item.Profesi.Daftar,

			item.Internas.Peminat, item.Internas.Lulus, item.Internas.Daftar,

			item.Pasca.Peminat, item.Pasca.Lulus, item.Pasca.Daftar,

			item.AdikPapua.Peminat, item.AdikPapua.Lulus, item.AdikPapua.Daftar,

			item.Jumlah.Peminat, item.Jumlah.Lulus, item.Jumlah.Daftar,

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
	filename := fmt.Sprintf("rekap_pmb_%d_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
