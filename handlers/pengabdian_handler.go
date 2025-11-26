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

type PengabdianHandler struct {
	PengabdianService *services.PengabdianService
}

func NewPengabdianHandler(service *services.PengabdianService) *PengabdianHandler {
	return &PengabdianHandler{
		PengabdianService: service,
	}
}

// GetPengabdianFiltered mendapatkan data pengabdian dengan filter dan pagination
// @Summary      Get filtered Pengabdian
// @Description  Mendapatkan data pengabdian berdasarkan filter dengan pagination
// @Tags         Pengabdian
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
// @Success      200           {object}  models.ListResponse{datas=[]models.Pengabdian}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /pengabdian [get]
func (h *PengabdianHandler) GetPengabdianFiltered(c *gin.Context) {
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

	pengabdian, total, err := h.PengabdianService.GetPengabdianFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, page, limit)
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
		"datas":  pengabdian,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *PengabdianHandler) ExportPengabdianCSV(c *gin.Context) {
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

	pengabdian, _, err := h.PengabdianService.GetPengabdianFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "Nama Kegiatan", "Nama Dosen Ketua", "Tahun Ajaran", "Semester",
		"Semester Type", "Tahun Data", "Periode", "Tanggal Data", "Waktu Pelaksanaan",
		"Created At", "Updated At", "Deleted At",

		"Jenis Pengabdian", "Skema", "Skim", "Id Skim", "Skema", "Rumpun Ilmu", "Bidang Penelitian",
		"ID Bidang Sub", "ID Sub Bidang", "Nama Sub Bidang", "ID Fokus", "ID Hibah", "ID Sosek Sub",

		"Sumber Dana", "Institusi Sumber Dana", "Dana (Nominal)", "Tahun Proposal", "Tahun Awal",
		"Tahun Implementasi", "Deskripsi", "Tujuan Ekonomi Sosial", "Sumber Data", "Posisi",

		"IsValid", "Status Lengkap", "Valid IPK", "Valid IPK Komentar", "Level Capaian",

		"Mahasiswa Pengabdian", "Anggota Pengabdian", "ID Sinta", "ID Silidia", "Tb Name", "Primary Key",

		"File Sampul", "File Daftar Isi", "File Lembar Pengesahan", "File Bukti Kerja", "File Proposal",
		"File Revisi", "File Laporan Kemajuan", "File Laporan Akhir", "File Mitra Awal",
		"Dokumen Pendukung", "Dokumen Pendukung Opsional",

		"Cron Tahun", "Cron Semester",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	var csvData [][]string

	for _, item := range pengabdian {

		idStr := strconv.Itoa(item.ID)

		row := []string{

			idStr, item.NamaKegiatan, item.NamaDosen, item.TahunAjaran, item.Semester,
			item.SemesterType, item.TahunData, item.Periode, item.Tanggal, item.WaktuPelaksanaan,
			item.CreatedAt, item.UpdatedAt, item.DeletedAt,

			item.JenisPengabdian, item.Skema, item.Skim, item.IDSkim, item.Skema, item.RumpunIlmu, item.BidangPenelitian,
			item.IDPenelitianBidangSub, item.IDSubBidang, item.NamaSubBidang, item.IDFokus, item.IDHibah, item.IDSosekSub,

			item.SumberDana, item.InstitusiSumberDana, item.Dana, item.TahunProposal, item.TahunAwal,
			item.TahunImplementasi, item.Deskripsi, item.TujuanEkonomiSosial, item.SumberData, item.Posisi,

			item.IsValid, item.StatusLengkap, item.ValidIpk, item.ValidIpkKomentar, item.LevelCapaian,

			item.MahasiswaPengabdian, item.AnggotaPengabdian, item.IDSinta, item.IDSilidia, item.TbName, item.PrimaryKey,

			item.FileSampul, item.FileDaftarIsi, item.FileLembarPengesahan, item.FileBuktiKerja, item.FileProposal,
			item.FileRevisi, item.FileLaporanKemajuan, item.FileLaporanAkhir, item.FileMitraAwal,
			item.DokumenPendukung, item.DokumenPendukungOpsional,

			item.CronTahun, item.CronSemester,

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
	filename := fmt.Sprintf("pengabdian_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
