package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type HkiHandler struct {
	HkiService *services.HkiService
}

func NewHkiHandler(service *services.HkiService) *HkiHandler {
	return &HkiHandler{
		HkiService: service,
	}
}

// GetHkiFiltered mendapatkan data hki dengan filter dan pagination
// @Summary      Get filtered hki
// @Description  Mendapatkan data hki berdasarkan filter dengan pagination
// @Tags         Hki
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
// @Success      200           {object}  models.ListResponse{datas=[]models.Hki}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /hki [get]
func (h *HkiHandler) GetHkiFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	search := c.Query("search")
	contentType := c.DefaultQuery("contentType", "json")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	hki, total, err := h.HkiService.GetHkiFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
			"datas":  hki,
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
			"datas":  hki,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	}
}

// ExportHkiCSV mengekspor data hki ke format CSV
// @Summary      Export Hki ke CSV
// @Description  Mengekspor daftar hki yang telah difilter ke dalam file CSV.
// @Tags         Hki
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        kodeFakultas    query    string false    "Filter berdasarkan Kode Fakultas"
// @Param        kodeJurusan     query    string false    "Filter berdasarkan Kode Jurusan"
// @Param        kodeProdi       query    string false    "Filter berdasarkan Kode Program Studi"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        semester        query    string false    "Filter berdasarkan semester"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /hki/export-csv [get]
func (h *HkiHandler) ExportHkiCSV(c *gin.Context) {
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

	hki, _, err := h.HkiService.GetHkiFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "Tahun Ajaran", "Semester", "Semester Type", "Tahun Data", "Periode", "Waktu Pelaksanaan",
		"Tanggal Pendaftaran", "Created At", "Updated At", "Deleted At", "Cron Tahun", "Cron Semester", "Deskripsi Karya",

		"Nama Karya", "Jenis Paten", "Kode Jenis Paten", "No Pendaftaran", "No Pendatatan Sertifikat",
		"Scope", "Kode Scope",

		"Nama Dosen Pencipta", "Posisi", "Create Dosen ID", "Level Capaian", "Jml Penulis", "Jml Negara Pengaku",

		"IsValid", "Valid IPK", "Valid IPK Komentar", "Komentar Validasi", "File Penilaian Reviewer",

		"Is Produk", "Sumber Produk",
		"Produk Penelitian Judul", "Produk Penelitian ID", "Produk Penelitian", "Mahasiswa Penelitian",
		"Produk Pengabdian Judul", "Produk Pengabdian ID", "Produk Pengabdian", "Anggota Penelitian",

		"File Bukti Kinerja", "File Sertifikat Paten", "File Pendaftaran",
		"File Pemeriksaan Substansi", "File Uji Publik", "File Sertifikasi",
		"File Hasil Uji Plagiarisme", "Tb Name", "Primary Key",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range hki {

		anggotaByte, _ := json.Marshal(item.AnggotaPenelitian)
		anggotaStr := string(anggotaByte)

		mhsByte, _ := json.Marshal(item.MahasiswaPenelitian)
		mhsStr := string(mhsByte)

		levelByte, _ := json.Marshal(item.LevelCapaian)
		levelStr := string(levelByte)

		prodPenelitianByte, _ := json.Marshal(item.ProdukPenelitian)
		prodPenelitianStr := string(prodPenelitianByte)

		prodPengabdianByte, _ := json.Marshal(item.ProdukPengabdian)
		prodPengabdianStr := string(prodPengabdianByte)
		idStr := strconv.Itoa(item.ID)

		row := []string{
			idStr,
			item.TahunAjaran,
			item.Semester,
			item.SemesterType,
			item.TahunData,
			item.Periode,
			item.WaktuPelaksanaan,
			item.Tanggal,
			item.CreatedAt,
			item.UpdatedAt,
			item.DeletedAt,
			item.CronTahun,
			item.CronSemester,
			item.Deskripsi,

			item.NamaKarya,
			item.JenisPaten,
			item.KodeJenisPaten,
			item.NoPendaftaran,
			item.NoPendatatanSertifikat,
			item.Scope,
			item.KodeScope,

			item.NamaDosen,
			item.Posisi,
			item.CreateDosenID,
			levelStr,
			item.JmlPenulis,
			item.JmlNegaraPengaku,

			item.IsValid,
			item.ValidIpk,
			item.ValidIpkKomentar,
			item.Komentar,
			item.FilePenilaianReviewer,

			item.IsProduk,
			item.SumberProduk,
			item.ProdukPenelitianJudul,
			item.ProdukPenelitianID,
			prodPenelitianStr,
			mhsStr,
			item.ProdukPengabdianJudul,
			item.ProdukPengabdianID,
			prodPengabdianStr,
			anggotaStr,

			item.FileBuktiKinerja,
			item.FileSertifikatPaten,
			item.FilePendaftaran,
			item.FilePemeriksaanSubstansi,
			item.FileUjiPublik,
			item.FileSertifikasi,
			item.FileHasilUjiPlagiarim,
			item.TbName,
			item.PrimaryKey,

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
	filename := fmt.Sprintf("hki_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
