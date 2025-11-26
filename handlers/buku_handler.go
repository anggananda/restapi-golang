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

type BukuHandler struct {
	BukuService *services.BukuService
}

func NewBukuHandler(service *services.BukuService) *BukuHandler {
	return &BukuHandler{
		BukuService: service,
	}
}

// GetBukuFiltered mendapatkan data buku dengan filter dan pagination
// @Summary      Get filtered Buku
// @Description  Mendapatkan data buku berdasarkan filter dengan pagination
// @Tags         Buku
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
// @Success      200           {object}  models.ListResponse{datas=[]models.Buku}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /buku [get]
func (h *BukuHandler) GetBukuFiltered(c *gin.Context) {
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

	buku, total, err := h.BukuService.GetBukuFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
		"datas":  buku,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *BukuHandler) ExportBukuCSV(c *gin.Context) {
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

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	buku, _, err := h.BukuService.GetBukuFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "Tahun Ajaran", "Semester", "Tahun Data", "Scope", "Kode Scope",
		"Periode", "Semester Type", "Waktu Pelaksanaan", "Tanggal Data", "Created At",
		"Updated At", "Deleted At", "Komentar Validasi",

		"Judul Buku", "Penerbit", "ISBN", "Kategori Buku", "Kode Kategori Buku", "Satuan",
		"Jumlah Halaman", "Volume Kegiatan", "Jml Negara Pengedaran", "Jml Penulis",

		"Nama Dosen Pembuat", "Nama Fakultas", "Nama Jurusan", "Kode Prodi", "Nama Prodi Unit",
		"Level Capaian",

		"IsValid", "Valid IPK", "Valid IPK Komentar", "Create Dosen ID",

		"Is Produk", "Sumber Produk",
		"Produk Penelitian Judul", "Produk Penelitian ID", "Produk Penelitian", "Mahasiswa Penelitian",
		"Produk Pengabdian Judul", "Produk Pengabdian ID", "Produk Pengabdian", "Anggota Penelitian",

		"URL Dokumen", "URL Per Review", "File Upload", "File Pendahuluan",
		"File Isi Buku", "File Penutup/Referensi", "File Persetujuan Penerbit",
		"File Selesai Dicetak", "File Hasil Uji Plagiarisme", "File Penilaian Reviewer",

		"Keterangan", "Posisi", "Tb Name", "Primary Key", "Cron Tahun", "Cron Semester",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range buku {
		idStr := strconv.Itoa(item.ID)

		row := []string{
			idStr,
			item.TahunAjaran,
			item.Semester,
			item.TahunData,
			item.Scope,
			item.KodeScope,
			item.Periode,
			item.SemesterType,
			item.WaktuPelaksanaan,
			item.Tanggal,
			item.CreatedAt,
			item.UpdatedAt,
			item.DeletedAt,
			item.Komentar,

			item.JudulBuku,
			item.Penerbit,
			item.ISBN,
			item.KategoriBuku,
			item.KodeKategoriBuku,
			item.Satuan,
			item.JumlahHalaman,
			item.VolumeKegiatan,
			item.JmlNegaraPengedaran,
			item.JmlPenulis,

			item.NamaDosen,
			item.NamaFakultas,
			item.NamaJurusan,
			item.KodeProdi,
			item.LevelCapaian,

			item.IsValid,
			item.ValidIpk,
			item.ValidIpkKomentar,
			item.CreateDosenID,

			item.IsProduk,
			item.SumberProduk,
			item.ProdukPenelitianJudul,
			item.ProdukPenelitianID,
			item.ProdukPenelitian,
			item.MahasiswaPenelitian,
			item.ProdukPengabdianJudul,
			item.ProdukPengabdianID,
			item.ProdukPengabdian,
			item.AnggotaPenelitian,

			item.UrlDokumen,
			item.UrlPerReview,
			item.FileUpload,
			item.FilePendahuluan,
			item.FileIsiBuku,
			item.FilePenutupReferensi,
			item.FilePersetujuanPenerbit,
			item.FileSelesaiDicetak,
			item.FileHasilUjiPlagiarim,
			item.FilePenilaianReviewer,

			item.Keterangan,
			item.Posisi,
			item.TbName,
			item.PrimaryKey,
			item.CronTahun,
			item.CronSemester,

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
	filename := fmt.Sprintf("buku_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
