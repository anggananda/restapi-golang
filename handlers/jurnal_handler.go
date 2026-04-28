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

type JurnalHandler struct {
	JurnalService *services.JurnalService
}

func NewJurnalHandler(service *services.JurnalService) *JurnalHandler {
	return &JurnalHandler{
		JurnalService: service,
	}
}

// GetJurnalFiltered mendapatkan data jurnal dengan filter dan pagination
// @Summary      Get filtered jurnal
// @Description  Mendapatkan data jurnal berdasarkan filter dengan pagination
// @Tags         Jurnal
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        semester      query     string  false  "Filter berdasarkan semester"
// @Param        indexer      query     string  false  "Filter berdasarkan indexer"
// @Param        akreditasi      query     string  false  "Filter berdasarkan akreditasi"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Jurnal}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /jurnal [get]
func (h *JurnalHandler) GetJurnalFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	indexer := c.Query("indexer")
	akreditasi := c.Query("akreditasi")
	search := c.Query("search")
	contentType := c.DefaultQuery("contentType", "json")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	jurnal, total, err := h.JurnalService.GetJurnalFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, akreditasi, search, page, limit)
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
			"datas":  jurnal,
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
			"datas":  jurnal,
			"pagination": gin.H{
				"page":  page,
				"limit": limit,
				"total": total,
				"pages": pages,
			},
		})
	}
}

// ExportJurnalCSV mengekspor data jurnal ke format CSV
// @Summary      Export Jurnal ke CSV
// @Description  Mengekspor daftar jurnal yang telah difilter ke dalam file CSV.
// @Tags         Jurnal
// @Accept       json
// @Produce      application/octet-stream
// @Param        limit           query    int    false    "Maksimal data yang akan diekspor"
// @Param        kodeFakultas    query    string false    "Filter berdasarkan Kode Fakultas"
// @Param        kodeJurusan     query    string false    "Filter berdasarkan Kode Jurusan"
// @Param        kodeProdi       query    string false    "Filter berdasarkan Kode Program Studi"
// @Param        tahun           query    string false    "Filter berdasarkan Tahun Ajaran (default: tahun sekarang)"
// @Param        semester        query    string false    "Filter berdasarkan semester"
// @Param        indexer        query    string false    "Filter berdasarkan indexer"
// @Param        akreditasi        query    string false    "Filter berdasarkan akreditasi"
// @Param        search          query    string false    "Pencarian bebas"
// @Success      200           {file}  string "File CSV berhasil diunduh"
// @Failure      500           {object}  models.ErrorResponse "Kesalahan pada server saat pengambilan data"
// @Security     BearerAuth
// @Router       /jurnal/export-csv [get]
func (h *JurnalHandler) ExportJurnalCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	indexer := c.Query("indexer")
	akreditasi := c.Query("akreditasi")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	jurnal, _, err := h.JurnalService.GetJurnalFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, akreditasi, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "Tahun Ajaran", "Semester", "Semester Type", "Tahun Data", "Periode", "Waktu Pelaksanaan",
		"Tanggal Submit", "Created At", "Updated At", "Cron Tahun", "Cron Semester", "Tb Name", "Primary Key", "Status Lengkap",

		"Judul Artikel", "Nama Jurnal", "Tahun Publish", "Akreditasi", "Kode Akreditasi", "Jenis Jurnal", "Kode Jenis Jurnal",
		"Penerbit", "Volume Jurnal", "Nomor Jurnal", "Halaman Awal", "Halaman Akhir",
		"P ISSN", "E ISSN", "DOI",

		"Nama Dosen Penulis", "Create Dosen ID", "Posisi", "Authors", "Jml Penulis", "Sitasi", "Sinta",

		"IsValid", "Valid IPK", "Valid IPK Komentar", "Komentar Validasi", "Bahasa ID", "Scope", "Kode Scope",

		"Aggregation Type", "Impact Factor", "Satuan", "Volume Kegiatan",
		"Is Produk", "Sumber Produk", "Produk Penelitian ID", "Produk Pengabdian ID", "Dari API Sinta",

		"Produk Penelitian Judul", "Produk Penelitian", "Mahasiswa Penelitian", "Anggota Penelitian",
		"Produk Pengabdian Judul", "Produk Pengabdian",

		"File Upload", "Alamat Web Jurnal", "URL Dokumen", "URL Peer Review", "File Submit",
		"File Revisi", "File Sudah Revisi", "File Diterima", "File Selesai Dicetak", "File Terpublikasi",
		"File Hasil Uji Plagiarisme", "File Penilaian Reviewer",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range jurnal {

		prodPenelitianByte, _ := json.Marshal(item.ProdukPenelitian)
		prodPenelitianStr := string(prodPenelitianByte)

		prodPengabdianByte, _ := json.Marshal(item.ProdukPengabdian)
		prodPengabdianStr := string(prodPengabdianByte)

		anggotaMhsByte, _ := json.Marshal(item.MahasiswaPenelitian)
		anggotaMhsStr := string(anggotaMhsByte)

		anggotaByte, _ := json.Marshal(item.AnggotaPenelitian)
		anggotaStr := string(anggotaByte)

		indexByte, _ := json.Marshal(item.Indexer)
		indexStr := string(indexByte)

		idStr := strconv.Itoa(item.ID)

		row := []string{
			idStr, item.TahunAjaran, item.Semester, item.SemesterType, item.TahunData, item.Periode, item.WaktuPelaksanaan,
			item.Tanggal, item.CreatedAt, item.UpdatedAt, item.CronTahun, item.CronSemester, item.TbName, item.PrimaryKey, item.StatusLengkap,
			item.JudulArtikel, item.NamaJurnal, item.TahunPublish, item.Akreditasi, item.KodeAkreditasi, item.JenisJurnal, item.KodeJenisJurnal,
			item.Penerbit, item.VolumeJurnal, item.NomorJurnal, item.HalamanAwal, item.HalamanAkhir,
			item.PIssn, item.EIssn, item.Doi,
			item.NamaDosen, item.CreateDosenID, item.Posisi, item.Authors, item.JmlPenulis, item.Sitasi, item.Sinta,
			item.IdSinta, indexStr,
			item.IsValid, item.ValidIpk, item.ValidIpkKomentar, item.Komentar, item.BahasaID, item.Scope, item.KodeScope,
			item.AggregationType, item.ImpactFactor, item.Satuan, item.VolumeKegiatan,
			item.IsProduk, item.SumberProduk, item.ProdukPenelitianID, item.ProdukPengabdianID, item.DariApiSinta,
			item.ProdukPenelitianJudul, prodPenelitianStr, anggotaMhsStr, anggotaStr,
			item.ProdukPengabdianJudul, prodPengabdianStr,
			item.FileUpload, item.AlamatWebJurnal, item.UrlDokumen, item.UrlPeerReview, item.FileSubmit,
			item.FileRevisi, item.FileSudahRevisi, item.FileDiterima, item.FileSelesaiDicetak, item.FileTerpublikasi,
			item.FileHasilUjiPlagiarim, item.FilePenilaianReviewer,
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
	filename := fmt.Sprintf("jurnal_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
