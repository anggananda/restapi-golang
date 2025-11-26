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

type KerjasamaHandler struct {
	KerjasamaService *services.KerjasamaService
}

func NewKerjasamaHandler(service *services.KerjasamaService) *KerjasamaHandler {
	return &KerjasamaHandler{
		KerjasamaService: service,
	}
}

// GetKerjasamaFiltered mendapatkan data kerjasama dengan filter dan pagination
// @Summary      Get filtered kerjasama
// @Description  Mendapatkan data kerjasama berdasarkan filter dengan pagination
// @Tags         Kerjasama
// @Accept       json
// @Produce      json
// @Param        page          query     int     false  "Nomor halaman (default: 1)"  default(1)  minimum(1)
// @Param        limit         query     int     false  "Jumlah data per halaman (default: 10, max: 100)"  default(10)  minimum(1)  maximum(100)
// @Param        kodeFakultas  query     string  false  "Filter berdasarkan kode fakultas"
// @Param        kodeJurusan   query     string  false  "Filter berdasarkan kode jurusan"
// @Param        kodeProdi     query     string  false  "Filter berdasarkan kode program studi"
// @Param        tahun         query     string  false  "Filter berdasarkan tahun akademik"
// @Param        search        query     string  false  "Pencarian global"
// @Success      200           {object}  models.ListResponse{datas=[]models.Kerjasama}
// @Failure      400           {object}  models.ErrorResponse
// @Failure      500           {object}  models.ErrorResponse
// @Security     BearerAuth
// @Router       /kerjasama [get]
func (h *KerjasamaHandler) GetKerjasamaFiltered(c *gin.Context) {
	page := utils.StringToInt(c.Query("page"), 1)
	limit := utils.StringToInt(c.Query("limit"), 10)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	kerjasama, total, err := h.KerjasamaService.GetKerjasamaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, search, page, limit)
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
		"datas":  kerjasama,
		"pagination": gin.H{
			"page":  page,
			"limit": limit,
			"total": total,
			"pages": pages,
		},
	})
}

func (h *KerjasamaHandler) ExportKerjasamaCSV(c *gin.Context) {
	limit := utils.StringToInt(c.Query("limit"), 0)
	kodeFakultas := c.Query("kodeFakultas")
	kodeJurusan := c.Query("kodeJurusan")
	kodeProdi := c.Query("kodeProdi")
	tahun := c.Query("tahun")
	search := c.Query("search")

	if tahun == "" {
		tahun = time.Now().Format("2006")
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 60*time.Second)
	defer cancel()

	kerjasama, _, err := h.KerjasamaService.GetKerjasamaFiltered(ctx, kodeFakultas, kodeJurusan, kodeProdi, tahun, search, 1, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	csvHeaders := []string{
		"ID", "KS Registrasi", "Tahun Dokumen", "Tahun Kerjasama", "Periode ID", "Created By ID", "Created By Name",
		"Updated By ID", "Updated By Name", "Waktu Mulai", "Waktu Selesai", "Tgl Laporan",
		"Created At", "Updated At", "Is Active", "Sedang Berjalan",

		"Deskripsi Singkat", "Jenis Asal Mitra ID", "Jenis Asal Mitra Nama", "Negara ID", "Negara Nama",
		"Bentuk Kerjasama ID", "Bentuk Kerjasama Nama", "Ruang Lingkup ID", "Ruang Lingkup Nama",
		"Ruang Lingkup Other", "Sasaran ID", "Sasaran Nama", "Indikator Kinerja ID",
		"Indikator Kinerja Nama", "Hasil Pelaksanaan", "Lap Singkat Pelaksanaan", "Luaran Volume",

		"Sumber Dana ID", "Sumber Dana Nama", "Alokasi Anggaran", "Nilai Kontrak", "Nilai Pendapatan",
		"Jenis Dokumen ID", "Jenis Dokumen Nama", "Nomor Dokumen Mitra", "Nomor Surat Undiksha",
		"Kategori Dokumen ID", "Kategori Dokumen Nama", "Ref KS Registrasi",

		"Durasi Hari", "Durasi Minggu", "Durasi Bulan",

		"Status Kerjasama Nama", "Status Mitra Nama", "ID Status Mitra", "Is Dokumen Valid",
		"Keterangan Validasi", "Validasi Nama", "Is Kampus QS", "Is JDIH", "Is Upload", "Is Upload Pusat",
		"Ref KS ID", "Ref KSRegistrasi",

		"Partner ID", "Partner Nama", "Partner Is Active", "Pihak Nama Mitra", "Alamat Mitra", "Telp Mitra",
		"Email Mitra", "Partner PJ Nama", "Partner PJ Jabatan", "Partner PJ Email",
		"Partner TTd Nama", "Partner TTd Jabatan",

		"UK ID Kerjasama", "UK Kerjasama Nama", "UK ID Pelaksana", "UK Pelaksana Nama",
		"PIC Dosen ID", "PIC Dosen Nama", "PIC Dosen Jabatan",
		"Undiksha PJ Nama", "Undiksha PJ Jabatan", "Undiksha PJ Email", "Undiksha PJ Kontak",
		"Undiksha TTd Nama", "Undiksha TTd Jabatan", "Undiksha TTd Email", "Undiksha TTd Kontak",

		"Dokumen URL", "Laporan URL", "Tautan Web",

		"Unit ID", "Unit UK Kode", "Unit Fakultas Kode", "Unit Jurusan Kode", "Unit Prodi Kode",
		"Fakultas Unit", "Jurusan Unit", "Prodi Unit",
	}
	var csvData [][]string

	for _, item := range kerjasama {
		idStr := strconv.FormatInt(item.ID, 10)

		createdAtStr := item.CreatedAt.Format("2006-01-02 15:04:05")
		updatedAtStr := item.UpdatedAt.Format("2006-01-02 15:04:05")

		row := []string{

			idStr, item.KSRegistrasi, item.TahunDokumen, item.Tahun, item.PeriodeID, item.CreatedBy, item.CreatedName,
			item.UpdatedBy, item.UpdatedName, item.TanggalAwal, item.TanggalAkhir, item.TglLaporan,
			createdAtStr, updatedAtStr, item.IsActive, item.SedangBerjalan,

			item.DeskripsiSingkat, item.JnsAsalMitraID, item.JnsAsalMitraNama, item.NegaraID, item.NegaraNama,
			item.BntkrjsmaID, item.BntkrjsmaNama, item.RuangLingkupID, item.RuangLingkupNama,
			item.RuangLingkupOther, item.SasaranID, item.SasaranNama, item.IndikatorKinerjaID,
			item.IndikatorKinerjaNama, item.HasilPelaksanaan, item.LapSingkatPelaksanaan, item.LuaranVolume,

			item.SumberdanaID, item.SumberdanaNama, item.AlokasiAnggaran, item.NilaiKontrak, item.NilaiPendapatan,
			item.JnsdokID, item.JnsdokNama, item.NomorDokumenMitra, item.NomorSuratUndiksha,
			item.KategoriDokID, item.KategoriDokNama, item.RefKSID,

			item.DurasiHari, item.DurasiMinggu, item.DurasiBulan,

			item.StskrjsmaNama, item.StatusMitraNama, item.IDStatusMitra, item.IsDokumenValid,
			item.KeteranganValidasi, item.ValidNama, item.IsKampusQS, item.IsJDIH, item.IsUpload, item.IsUploadPusat,
			item.RefKSID, item.RefKSRegistrasi,

			item.PartnerID, item.PartnerNama, item.PartnerIsActive, item.PihakNamaMitra, item.Alamat, item.Telp,
			item.Email, item.PartnerPenanggungjawabNama, item.PartnerPenanggungjawabJabatan, item.PartnerPenanggungjawabEmail,
			item.PartnerPenandatanganNama, item.PartnerPenandatanganJabatan,

			item.UKIDKerjasama, item.UKKerjasama, item.UKIDPelaksana, item.UKPelaksana,
			item.UndikshaPICDosenID, item.UndikshaPICDosenNama, item.UndikshaPICDosenJabatan,
			item.UndikshaPenanggungjawabNama, item.UndikshaPenanggungjawabJabatan, item.UndikshaPenanggungjawabEmail, item.UndikshaPenanggungjawabKontakTelpHp,
			item.UndikshaPenandatanganNama, item.UndikshaPenandatanganJabatan, item.UndikshaPenandatanganEmail, item.UndikshaPenandatanganKontakTelpHp,

			item.DokumenURL, item.LaporanURL, item.Tautan,

			item.Unit.UKID,
			item.Unit.UKKode,
			item.Unit.FKTKode,
			item.Unit.JRSKode,
			item.Unit.PRDKode,
			item.Unit.Fakultas,
			item.Unit.Jurusan,
			item.Unit.Prodi,
		}

		csvData = append(csvData, row)
	}

	currentTime := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("kerjasama_%s_%s", tahun, currentTime)
	utils.SendCSV(c, filename, csvHeaders, csvData)
}
