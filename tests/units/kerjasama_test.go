package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setupMockKerjasamaData() []models.Kerjasama {
	now := time.Now()
	tanggalAkhirDomestik := time.Now().AddDate(1, 0, 0).Format("2006-01-02")
	tanggalAkhirInternasional := time.Now().AddDate(3, 0, 0).Format("2006-01-02")

	return []models.Kerjasama{

		{
			ID:                   1,
			Alamat:               "Jl. Sudirman No. 12, Jakarta",
			AlokasiAnggaran:      "50000000",
			BntkrjsmaID:          "1",
			BntkrjsmaNama:        "Memorandum of Understanding (MoU)",
			CreatedBy:            "NIP123",
			CreatedName:          "Dr. Budi Santoso",
			DeskripsiSingkat:     "Kerja sama di bidang penelitian dan pengabdian masyarakat.",
			DurasiBulan:          "12",
			Email:                "info@mitra-a.co.id",
			IDStatusMitra:        "1",
			IndikatorKinerjaID:   "1",
			IndikatorKinerjaNama: "Penelitian Bersama",
			IsActive:             "Y",
			IsDokumenValid:       "Y",
			IsJDIH:               "N",
			IsKampusQS:           "N",
			JnsAsalMitraID:       "1",
			JnsAsalMitraNama:     "Perguruan Tinggi Domestik",
			JnsdokID:             "1",
			JnsdokNama:           "Nota Kesepahaman (MoU)",
			KategoriDokID:        "1",
			KategoriDokNama:      "Pendidikan",
			KSRegistrasi:         "KS/2023/001",
			NegaraID:             "ID",
			NegaraNama:           "Indonesia",
			NilaiKontrak:         "0",
			NomorDokumenMitra:    "123/MITRA-A/KS/2023",
			PartnerID:            "P001",
			PartnerNama:          "Universitas Mitra Jaya",
			RuangLingkupID:       "1",
			RuangLingkupNama:     "Pendidikan dan Penelitian",
			SasaranID:            "1",
			SasaranNama:          "Dosen dan Mahasiswa",
			SedangBerjalan:       "Y",
			StatusMitraNama:      "Aktif",
			StskrjsmaNama:        "Berjalan",
			SumberdanaID:         "1",
			SumberdanaNama:       "Internal",
			Tahun:                "2023",
			TanggalAkhir:         tanggalAkhirDomestik,
			TanggalAwal:          now.Format("2006-01-02"),
			Unit: models.UnitDetail{
				FKTKode:  "F01",
				Fakultas: "Teknik",
				Prodi:    "Informatika",
			},
			CreatedAt: now,
			UpdatedAt: now,
		},

		{
			ID:                   2,
			Alamat:               "10 Downing Street, London",
			AlokasiAnggaran:      "10000",
			BntkrjsmaID:          "2",
			BntkrjsmaNama:        "Memorandum of Agreement (MoA)",
			CreatedBy:            "NIP456",
			CreatedName:          "Dr. Siti Rahayu",
			DeskripsiSingkat:     "Program pertukaran mahasiswa dan staf akademik.",
			DurasiBulan:          "36",
			Email:                "contact@univ-global.uk",
			IDStatusMitra:        "2",
			IndikatorKinerjaID:   "2",
			IndikatorKinerjaNama: "Pertukaran Mahasiswa",
			IsActive:             "Y",
			IsDokumenValid:       "Y",
			IsJDIH:               "Y",
			IsKampusQS:           "Y",
			JnsAsalMitraID:       "2",
			JnsAsalMitraNama:     "Perguruan Tinggi Internasional",
			JnsdokID:             "2",
			JnsdokNama:           "Perjanjian Kerja Sama (PKS)",
			KategoriDokID:        "2",
			KategoriDokNama:      "Penyelenggaraan Pendidikan",
			KSRegistrasi:         "KS/2023/002/INT",
			NegaraID:             "GB",
			NegaraNama:           "Inggris",
			NilaiKontrak:         "50000",
			NomorDokumenMitra:    "GLOBAL/2023/MoA/100",
			PartnerID:            "P002",
			PartnerNama:          "Global University of Technology",
			RuangLingkupID:       "2",
			RuangLingkupNama:     "Pengembangan SDM",
			SasaranID:            "2",
			SasaranNama:          "Mahasiswa",
			SedangBerjalan:       "Y",
			StatusMitraNama:      "Aktif",
			StskrjsmaNama:        "Berjalan",
			SumberdanaID:         "2",
			SumberdanaNama:       "Luar Negeri",
			Tahun:                "2023",
			TanggalAkhir:         tanggalAkhirInternasional,
			TanggalAwal:          now.Format("2006-01-02"),
			Unit: models.UnitDetail{
				FKTKode:  "F03",
				Fakultas: "Bahasa",
				Prodi:    "Sastra Inggris",
			},
			CreatedAt: now,
			UpdatedAt: now,
		},
	}
}

func TestKerjasamaFiltered_Success(t *testing.T) {

	mockData := setupMockKerjasamaData()

	mockRepo := mocks.NewKerjasamaMockRepository(mockData, nil)
	service := services.NewKerjasamaService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetKerjasamaFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetKerjasamaFiltered(
		ctx,
		"", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Universitas Mitra Jaya", result[0].PartnerNama, "Verifikasi nama mitra entri pertama benar")
	assert.Equal(t, "Global University of Technology", result[1].PartnerNama, "Verifikasi nama mitra entri kedua benar")
	assert.Equal(t, "Indonesia", result[0].NegaraNama, "Verifikasi negara entri pertama benar")
	assert.Equal(t, "Inggris", result[1].NegaraNama, "Verifikasi negara entri kedua benar")
}
