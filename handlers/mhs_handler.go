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

type MhsHandler struct {
	MhsService *services.MhsService
}

func NewMhsHandler(service *services.MhsService) *MhsHandler {
	return &MhsHandler{
		MhsService: service,
	}
}

// GetDetailMhs mendapatkan  detail mahasiswa
// @Summary      Get detail mahasiswa
// @Description  Mendapatkan detail data mahasiswa
// @Tags         Mahasiswa
// @Accept       json
// @Produce      json
// @Param        nim        query     string  false  "nim"
// @Success      200           {object}  models.ListDetailResponse{datas=models.Mahasiswa}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /mhs/{nim} [get]
func (h *MhsHandler) GetDetailMhs(c *gin.Context) {
	nim := c.Param("nim")
	contentType := c.DefaultQuery("contentType", "json")
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	mh, err := h.MhsService.GetDetailMhs(ctx, nim)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if contentType == "msgpack" {
		utils.Render(c, http.StatusOK, gin.H{"datas": mh, "message": "OK"})
	} else {
		c.JSON(http.StatusOK, gin.H{"datas": mh, "message": "OK"})
	}
}

// GetMahasiswaHistoryFiltered mendapatkan  mahasiswa  dengan filter dan pagination
// @Summary      Get filtered mahasiswa
// @Description  Mendapatkan data mahasiswa berdasarkan filter dengan pagination
// @Tags         Mahasiswa
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        semester         query     int  false  "Filter berdasarkan semester"
// @Param        angkatan         query     int  false  "Filter berdasarkan angkatan"
// @Param        status         query     int  false  "Filter berdasarkan status"
// @Param        kewarganegaraan         query     string  false  "Filter berdasarkan kewarganegaraan"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.MahasiswaHistoryResponse}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /mhs/history [get]
func (h *MhsHandler) GetMahasiswaHistoryFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahunStr := c.Query("tahun")
	semester := utils.StringToInt(c.Query("semester"), 0)
	angkatan := utils.StringToInt(c.Query("angkatan"), 0)
	search := c.Query("search")
	status := utils.StringToInt(c.Query("status"), 0)
	kewarganegaraan := c.Query("kewarganegaraan")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	results, total, err := h.MhsService.GetMahasiswaHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, kewarganegaraan, search, tahun, semester, angkatan, status, page, limit)
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
			"datas":  results,
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
			"datas":  results,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	}
}

// ExportMhsCSV mengekspor data mahasiswa ke format CSV
// @Summary      Export Mahasiswa ke CSV
// @Description  Mengekspor daftar mahasiswa yang telah difilter ke dalam file CSV.
// @Tags         Mahasiswa
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        kodeFakultas    query    string false    "Filter berdasarkan Kode Fakultas"
// @Param        kodeJurusan     query    string false    "Filter berdasarkan Kode Jurusan"
// @Param        kodeProdi       query    string false    "Filter berdasarkan Kode Program Studi"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        semester        query    int false    "Filter berdasarkan semester"
// @Param        angkatan        query    int false    "Filter berdasarkan angkatan"
// @Param        status        query    int false    "Filter berdasarkan status"
// @Param        kewarganegaraan        query    string false    "Filter berdasarkan kewarganegaraan"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /mhs/history/export-csv [get]
func (h *MhsHandler) ExportMhsCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahunStr := c.Query("tahun")
	semester := utils.StringToInt(c.Query("semester"), 0)
	angkatan := utils.StringToInt(c.Query("angkatan"), 0)
	search := c.Query("search")
	status := utils.StringToInt(c.Query("status"), 0)
	kewarganegaraan := c.Query("kewarganegaraan")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	mahasiswa, _, err := h.MhsService.GetMahasiswaHistoryFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, kewarganegaraan, search, tahun, semester, angkatan, status, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"NIM",
		"Nama Lengkap",
		"Fakultas",
		"Jurusan",
		"Prodi",
		"Angkatan",
		"Tahun",
		"Semester",
		"Status",
		"Periode",
		"Kewarganegaraan",
		"Telepon",
		"Email SSO",
		"Nama PA",
	}
	var csvData [][]string

	for _, item := range mahasiswa {
		tahunStr := strconv.Itoa(item.Tahun)
		semesterStr := strconv.Itoa(item.Semester)

		row := []string{
			item.NIM,
			item.Nama,
			item.Fakultas,
			item.Jurusan,
			item.Prodi,
			item.TahunMasuk,
			tahunStr,
			semesterStr,
			item.Status,
			item.Periode,
			item.Kewarganegaraan,
			item.Telp,
			item.EmailSSO,
			item.NamaPA,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("mahasiswa_history_%d_%d_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)

}
