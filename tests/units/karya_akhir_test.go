package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockKaryaAkhirData() []models.KaryaAkhir {
	return []models.KaryaAkhir{
		{
			ID:              1,
			CurrentState:    "Lulus Sidang",
			Judul:           "Sistem Pendukung Keputusan Pemilihan Dosen Pembimbing Skripsi",
			MainStage:       "Selesai",
			NamaLengkap:     "Budi Setiawan",
			NamaPA:          "Dr. Ani Susanti",
			NamaPembimbing1: "Prof. Dr. Anton Yudha",
			NamaPembimbing2: "Dr. Liana Dewi",
			NamaPembimbing3: "",
			NamaPenguji1:    "Dr. Rina Wulandari",
			NamaPenguji2:    "Ir. Surya Permana",
			NamaPenguji3:    "Dr. Dian Kusuma",
			NamaPenguji4:    "",
			NamaPenguji5:    "",
			NamaPenguji6:    "",
			NilaiAkhir:      "A",
			NIM:             "1905551001",
			StatusJudul:     "Diterima",
			TahunMasuk:      2019,
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:              2,
			CurrentState:    "Proses Bimbingan",
			Judul:           "Analisis Sentimen Ulasan Pengguna Aplikasi E-commerce X",
			MainStage:       "Bimbingan",
			NamaLengkap:     "Citra Ayu",
			NamaPA:          "Dr. Budi Santoso",
			NamaPembimbing1: "Dr. Santi Melati",
			NamaPembimbing2: "Ir. Made Wijaya",
			NamaPembimbing3: "Dr. Putu Eka",
			NamaPenguji1:    "",
			NamaPenguji2:    "",
			NamaPenguji3:    "",
			NamaPenguji4:    "",
			NamaPenguji5:    "",
			NamaPenguji6:    "",
			NilaiAkhir:      "",
			NIM:             "2005552002",
			StatusJudul:     "Diterima",
			TahunMasuk:      2020,
			Unit: models.Unit{
				FktKode:  "F02",
				JrsKose:  "J02",
				PrdKode:  "P02",
				Fakultas: "Ekonomi",
				Jurusan:  "Manajemen",
			},
		},
	}
}

func TestKaryaAkhirFiltered_Success(t *testing.T) {

	mockData := setupMockKaryaAkhirData()

	mockRepo := mocks.NewKaryaAkhirMockRepository(mockData, nil)
	service := services.NewKaryaAkhirService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetKaryaAkhirFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetKaryaAkhirFiltered(
		ctx,
		"", "", "", "", 0,
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Budi Setiawan", result[0].NamaLengkap, "Verifikasi nama lengkap entri pertama benar")
	assert.Equal(t, "Sistem Pendukung Keputusan Pemilihan Dosen Pembimbing Skripsi", result[0].Judul, "Verifikasi judul entri pertama benar")
	assert.Equal(t, "Citra Ayu", result[1].NamaLengkap, "Verifikasi nama lengkap entri kedua benar")
}
