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

type TracerHandler struct {
	TracerService *services.TracerService
}

func NewTracerHandler(service *services.TracerService) *TracerHandler {
	return &TracerHandler{
		TracerService: service,
	}
}

// GetTracerFiltered mendapatkan data tracer dengan filter dan pagination
// @Summary      Get filtered tracer
// @Description  Mendapatkan data tracer berdasarkan filter dengan pagination
// @Tags         Tracer
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     int  false  "Filter berdasarkan tahun akademik"
// @Param        bulan         query     int  false  "Filter berdasarkan bulan"
// @Param        statusTracer         query     string  false  "Filter berdasarkan statusTracer"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Tracer}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /tracer [get]
func (h *TracerHandler) GetTracerFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	tahunStr := c.Query("tahun")
	bulan := utils.StringToInt(c.Query("bulan"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	statusTracer := c.Query("statusTracer")
	search := c.Query("search")
	contentType := c.DefaultQuery("contentType", "json")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	tracer, total, err := h.TracerService.GetTracerFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search, tahun, bulan, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	pages := int64(0)
	if limit > 0 {
		pages = int64(math.Ceil(float64(total) / float64(limit)))
	}

	if contentType == "mgspack" {
		utils.Render(c, http.StatusOK, gin.H{
			"status": "success",
			"datas":  tracer,
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
			"datas":  tracer,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	}
}

// ExportTracerCSV mengekspor data tracer ke format CSV
// @Summary      Export Tracer ke CSV
// @Description  Mengekspor daftar tracer yang telah difilter ke dalam file CSV.
// @Tags         Tracer
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
// @Router       /tracer/export-csv [get]
func (h *TracerHandler) ExportTracerCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	tahunStr := c.Query("tahun")
	bulan := utils.StringToInt(c.Query("bulan"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	statusTracer := c.Query("statusTracer")
	search := c.Query("search")

	var tahun int
	if tahunStr == "" {
		tahunStr = time.Now().Format("2006")
	}
	tahun = utils.StringToInt(tahunStr, 0)

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	tracer, _, err := h.TracerService.GetTracerFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, statusTracer, search, tahun, bulan, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "ID Mahasiswa", "User ID", "NIM", "Nama Lengkap", "Jenis Kelamin",
		"Tgl Lahir", "Email", "No Telp", "NIK", "NPWP", "Jenjang",
		"Bulan Lulus", "Tahun Lulus", "Tgl Lulus", "IPK", "Status Mahasiswa",
		"Bulan Wisuda", "Tahun Wisuda", "No Ijazah", "No SK Yudisium", "ID Jurusan",
		"Status Pengisian", "Persentase Pengisian", "Pengisian Terakhir", "Dikti Status",
		"Status Saat Ini (Kerja/Lanjut/dll)", "Masa Tunggu Sebelum Lulus",
		"Masa Tunggu Setelah Lulus", "Provinsi", "Kabupaten", "Deleted At",
		"Nama Ayah", "Nama Ibu", "Nama Saudara", "Nama Wali",
		"Nama Perusahaan", "Alamat Perusahaan", "Jenis Perusahaan", "Gaji",
		"Jabatan Berwirausaha", "Tingkat Tempat Kerja", "Sumber Biaya Studi Lanjut",
		"PT Lanjut", "Prodi Masuk Lanjut", "Tanggal Masuk Lanjut", "Sumber Biaya Lanjut",
		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range tracer {

		idStr := strconv.Itoa(item.ID)
		idMhsStr := strconv.Itoa(item.IDMahasiswa)
		userIdStr := strconv.Itoa(item.UserID)
		idJurusanStr := strconv.Itoa(item.IDJurusan)

		ipkStr := strconv.FormatFloat(item.IPKMahasiswa, 'f', 2, 64)
		persenStr := strconv.Itoa(item.PersentasePengisian)

		bulanLulusStr := strconv.Itoa(item.BulanLulusMahasiswa)
		tahunLulusStr := strconv.Itoa(item.TahunLulusMahasiswa)

		bulanWisudaStr := strconv.Itoa(item.BulanWisuda)
		tahunWisudaStr := strconv.Itoa(item.TahunWisuda)

		tglLahirStr := item.TglLahirMahasiswa.Format("02-01-2006")
		tglLulusStr := item.TglLulusMahasiswa.Format("02-01-2006")
		pengisianTerakhirStr := item.PengisianTerakhir.Format("02-01-2006 15:04:05")
		deletedAtStr := item.DeletedAt.Format("02-01-2006 15:04:05")

		row := []string{

			idStr, idMhsStr, userIdStr, item.NIMMahasiswa, item.NamaMahasiswa, item.JenisKelaminMahasiswa,
			tglLahirStr, item.EmailMahasiswa, item.NoTelp, item.NIKMahasiswa, item.NPWPMahasiswa, item.Jenjang,

			bulanLulusStr, tahunLulusStr, tglLulusStr, ipkStr, item.StatusMahasiswa,
			bulanWisudaStr, tahunWisudaStr, item.NoIjasah, item.NoSKYudisium, idJurusanStr,

			item.StatusPengisian, persenStr, pengisianTerakhirStr, item.Dikti,

			item.StatusSaatIni, item.MasaTungguSebelumLulus,
			item.MasaTungguSetelahLulus, item.Provinsi, item.Kabupaten, deletedAtStr,

			item.Ayah, item.Ibu, item.Saudara, item.Wali,

			item.NamaPerusahaan, item.AlamatPerusahaan, item.JenisPerusahaan, item.Gaji,
			item.JabatanDalamBerwirausaha, item.TingkatTempatKerja, item.SumberBiayaStudiLanjut,

			item.PerguruanTinggiLanjut, item.ProdiMasukStudiLanjut, item.TanggalMasukStudiLanjut, item.SumberBiayaStudiLanjut,

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
	filename := fmt.Sprintf("tracer_%d_%d_%s", tahun, bulan, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
