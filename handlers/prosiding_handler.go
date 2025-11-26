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

type ProsidingHandler struct {
	ProsidingService *services.ProsidingService
}

func NewProsidingHandler(service *services.ProsidingService) *ProsidingHandler {
	return &ProsidingHandler{
		ProsidingService: service,
	}
}

// GetProsidingFiltered mendapatkan data prosiding dengan filter dan pagination
// @Summary      Get filtered prosiding
// @Description  Mendapatkan data prosiding berdasarkan filter dengan pagination
// @Tags         Prosiding
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        semester         query     string  false  "Filter berdasarkan semester"
// @Param        indexer         query     string  false  "Filter berdasarkan indexer"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Prosiding}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /prosiding [get]
func (h *ProsidingHandler) GetProsidingFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	indexer := c.Query("indexer")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	prosiding, total, err := h.ProsidingService.GetProsidingFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search, page, limit)
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
		"datas":  prosiding,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *ProsidingHandler) ExportProsidingCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	indexer := c.Query("indexer")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	prosiding, _, err := h.ProsidingService.GetProsidingFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, indexer, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{

		"ID", "Tahun Ajaran", "Semester", "Semester Type", "Tahun Data", "Periode",
		"Waktu Pelaksanaan", "Tanggal Data", "Created At", "Updated At", "Deleted At",
		"Cron Tahun", "Cron Semester", "Keterangan",

		"Judul Artikel", "Nama Seminar", "Tanggal Awal Acara", "Tanggal Akhir Acara",
		"Tempat Pelaksanaan", "Penyelenggara", "Jenis Pembicara", "Tipe Prosiding",
		"Kode Tipe Prosiding", "Penerbit", "P ISSN", "E ISSN", "ISBN",
		"Scope", "Kode Scope", "Status Publish",

		"Nama Dosen Ketua", "Create Dosen ID", "Posisi", "Jumlah Penulis",
		"Anggota Penelitian (Gabungan)", "Anggota Non Dosen (Gabungan)",
		"Mahasiswa Penelitian (Gabungan)", "Sinta", "ID Sinta",

		"IsValid", "Valid IPK", "Valid IPK Komentar", "Komentar Validasi",
		"Bereputasi", "Satuan", "Volume Kegiatan", "Indexer (Gabungan)",

		"Is Produk", "Sumber Produk",
		"Produk Penelitian Judul", "Produk Penelitian ID", "Produk Penelitian",
		"Produk Pengabdian Judul", "Produk Pengabdian ID", "Produk Pengabdian",

		"File Upload (Dokumen)", "URL Dokumen", "URL Peer Review",
		"File Penilaian Reviewer", "File Hasil Uji Plagiarisme",
		"Tanggal Submit", "Tb Name", "Primary Key",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range prosiding {

		idStr := strconv.Itoa(item.ID)

		indexerList := strings.Join(item.Indexer, " | ")

		row := []string{

			idStr, item.TahunAjaran, item.Semester, item.SemesterType, item.TahunData, item.Periode,
			item.WaktuPelaksanaan, item.Tanggal, item.CreatedAt, item.UpdatedAt, item.DeletedAt,
			item.CronTahun, item.CronSemester, item.Keterangan,

			item.JudulArtikel, item.NamaSeminar, item.TglAwal, item.TglAkhir,
			item.TempatPelaksanaan, item.Penyelenggara, item.JenisPembicara, item.TipeProsiding,
			item.KodeTipeProsiding, item.Penerbit, item.PIssn, item.EIssn, item.ISBN,
			item.Scope, item.KodeScope, item.StatusPublish,

			item.NamaDosen, item.CreateDosenID, item.Posisi, item.JmlPenulis,
			item.AnggotaPenelitian,
			item.AnggotaNonDosen,
			item.MahasiswaPenelitian,
			item.Sinta, item.IdSinta,

			item.IsValid, item.ValidIpk, item.ValidIpkKomentar, item.Komentar,
			item.Bereputasi, item.Satuan, item.VolumeKegiatan, indexerList,

			item.IsProduk, item.SumberProduk,
			item.ProdukPenelitianJudul, item.ProdukPenelitianID, item.ProdukPenelitian,
			item.ProdukPengabdianJudul, item.ProdukPengabdianID, item.ProdukPengabdian,

			item.FileUpload, item.UrlDokumen, item.UrlPerReview,
			item.FilePenilaianReviewer, item.FileHasilUjiPlagiarim,
			item.Tanggal, item.TbName, item.PrimaryKey,

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
	filename := fmt.Sprintf("prosiding_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
