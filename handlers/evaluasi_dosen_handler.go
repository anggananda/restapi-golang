package handlers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type EvaluasiDosenHandler struct {
	EvaluasiDosenService *services.EvaluasiDosenService
}

func NewEvaluasiDosenHandler(service *services.EvaluasiDosenService) *EvaluasiDosenHandler {
	return &EvaluasiDosenHandler{
		EvaluasiDosenService: service,
	}
}

// GetEvaluasiDosenFiltered mendapatkan data Evaluasi Dosen dengan filter dan pagination
// @Summary      Get filtered Evaluasi Dosen
// @Description  Mendapatkan data Evaluasi Dosen berdasarkan filter dengan pagination
// @Tags         Evaluasi Dosen
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        semester      query     string  false  "Filter berdasarkan semester"
// @Param        namaDosen        query     string  false  "Filter berdasarkan namaDosen"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.EvaluasiDosen}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /evaluasi-dosen [get]
func (h *EvaluasiDosenHandler) GetEvaluasiDosenFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	namaDosen := c.Query("namaDosen")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	evaluasiDosen, total, err := h.EvaluasiDosenService.GetEvaluasiDosenFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, namaDosen, search, page, limit)
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
		"datas":  evaluasiDosen,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *EvaluasiDosenHandler) ExportEvaluasiDosenCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	semester := c.Query("semester")
	namaDosen := c.Query("namaDosen")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	evaluasiDosen, _, err := h.EvaluasiDosenService.GetEvaluasiDosenFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, semester, namaDosen, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var aspekPenilaian = []string{
		"Perencanaan Perkuliahan",
		"Relevansi Materi",
		"Penguasaan Materi",
		"Metode dan Pendekatan",
		"Inovasi",
		"Kreatifitas",
		"Media Pembelajaran",
		"Sumber Belajar",
		"Penilaian Hasil Belajar",
		"Penilaian Proses Belajar",
		"Pemberian Tugas",
		"Pengelolaan Kelas",
		"Motivasi dan Antusiasme",
		"Penciptaan Iklim Belajar",
		"Kedisiplinan",
		"Penegakan Aturan",
		"Pengembangan Karakter",
		"Keteladanan Sikap",
		"Kemampuan Berkomunikasi",
		"Penggunaan Bahasa",
		"Kemampuan Berinteraksi Sosial",
	}

	csvHeaders := []string{
		"ID", "NIP", "No Induk Undiksha", "Nama Dosen", "Nama Kelas",
		"Kode Matakuliah", "Nama Matakuliah", "Tahun", "Semester",
		"Kode Prodi", "Kode Fakultas",
		"Created At", "Updated At", "Deleted At",

		"Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}

	for _, aspek := range aspekPenilaian {
		csvHeaders = append(csvHeaders,
			fmt.Sprintf("%s - Sangat Baik", aspek),
			fmt.Sprintf("%s - Baik", aspek),
			fmt.Sprintf("%s - Cukup", aspek),
			fmt.Sprintf("%s - Kurang", aspek),
		)
	}

	var csvData [][]string

	for _, item := range evaluasiDosen {
		idStr := strconv.FormatInt(item.ID, 10)

		createdAtStr := item.CreatedAt.Format("2006-01-02 15:04:05")
		updatedAtStr := item.UpdatedAt.Format("2006-01-02 15:04:05")
		deletedAtStr := ""
		if item.DeletedAt != nil {
			deletedAtStr = item.DeletedAt.Format("2006-01-02 15:04:05")
		}

		row := []string{

			idStr,
			item.NIP,
			item.NoIndukUndiksha,
			item.NamaLengkap,
			item.NamaKelas,
			item.KodeMatakuliah,
			item.NamaMatakuliah,
			item.Tahun,
			item.Semester,
			item.KodeProdi,
			item.KodeFakultas,
			createdAtStr,
			updatedAtStr,
			deletedAtStr,

			item.Unit.UKKode,
			item.Unit.FktKode,
			item.Unit.JrsKose,
			item.Unit.PrdKode,
			item.Unit.Fakultas,
			item.Unit.Jurusan,
			item.Unit.Prodi,
		}

		penilaianList := []models.Penilaian{
			item.PerencanaanPerkuliahan,
			item.RelevansiMateriDenganTujuanPembelajaran,
			item.PenguasaanMateriPerkuliahan,
			item.MetodeDanPendekatanPerkuliahan,
			item.InovasiDalamPerkuliahan,
			item.KreatifitasDalamPerkuliahan,
			item.MediaPembelajaran,
			item.SumberBelajar,
			item.PenilaianHasilBelajar,
			item.PenilaianProsesBelajar,
			item.PemberianTugasPerkuliahan,
			item.PengelolaanKelas,
			item.MotivasiDanAntusiasmeMengajar,
			item.PenciptaanIklimBelajar,
			item.Kedisiplinan,
			item.PenegakanAturanPerkuliahan,
			item.PengembanganKarakterMahasiswa,
			item.KeteladananDalamBersikapDanBertindak,
			item.KemampuanBerkomunikasi,
			item.PenggunaanBahasaLisanDanTulisan,
			item.KemampuanBerinteraksiSosialDenganMahasiswa,
		}

		for _, p := range penilaianList {

			row = append(row, strconv.Itoa(p.SangatBaik))
			row = append(row, strconv.Itoa(p.Baik))
			row = append(row, strconv.Itoa(p.Cukup))
			row = append(row, strconv.Itoa(p.Kurang))
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("evaluasi_dosen_%s_%s_%s", tahun, semester, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
