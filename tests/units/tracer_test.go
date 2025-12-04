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

func setupMockTracerData() []models.Tracer {
	tglLahir1 := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
	tglLulus1 := time.Date(2023, time.September, 15, 0, 0, 0, 0, time.UTC)
	tglLahir2 := time.Date(2001, time.March, 10, 0, 0, 0, 0, time.UTC)
	tglLulus2 := time.Date(2024, time.March, 20, 0, 0, 0, 0, time.UTC)

	return []models.Tracer{
		{
			ID:                     1,
			IDMahasiswa:            101,
			NIMMahasiswa:           "1234567890",
			NamaMahasiswa:          "Dian Anggraini",
			EmailMahasiswa:         "dian@example.com",
			NoTelp:                 "08123456789",
			TglLahirMahasiswa:      tglLahir1,
			JenisKelaminMahasiswa:  "P",
			BulanLulusMahasiswa:    9,
			TahunLulusMahasiswa:    2023,
			IPKMahasiswa:           3.75,
			TglLulusMahasiswa:      tglLulus1,
			StatusMahasiswa:        "Alumni",
			StatusPengisian:        "Lengkap",
			PersentasePengisian:    100,
			StatusSaatIni:          "Bekerja",
			MasaTungguSetelahLulus: "0-6 bulan",
			Gaji:                   "5000000 - 7000000",
			JenisPerusahaan:        "Swasta Nasional",
			NamaPerusahaan:         "PT. Solusi Digital Indonesia",
			TingkatTempatKerja:     "Nasional",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:                     2,
			IDMahasiswa:            102,
			NIMMahasiswa:           "0987654321",
			NamaMahasiswa:          "Fajar Nugroho",
			EmailMahasiswa:         "fajar@example.com",
			NoTelp:                 "08987654321",
			TglLahirMahasiswa:      tglLahir2,
			JenisKelaminMahasiswa:  "L",
			BulanLulusMahasiswa:    3,
			TahunLulusMahasiswa:    2024,
			IPKMahasiswa:           3.25,
			TglLulusMahasiswa:      tglLulus2,
			StatusMahasiswa:        "Alumni",
			StatusPengisian:        "Lengkap",
			PersentasePengisian:    100,
			StatusSaatIni:          "Studi Lanjut",
			MasaTungguSetelahLulus: "N/A",
			Gaji:                   "0",
			PerguruanTinggiLanjut:  "Universitas Gadjah Mada",
			ProdiMasukStudiLanjut:  "Magister Manajemen",
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

func TestTracerFiltered_Success(t *testing.T) {
	mockData := setupMockTracerData()
	mockRepo := mocks.NewTracerMockRepository(mockData, nil)
	service := services.NewTracerService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetTracerFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetTracerFiltered(
		ctx,
		"", "", "", "", "", 0, 0,
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Dian Anggraini", result[0].NamaMahasiswa, "Verifikasi nama mahasiswa entri pertama benar")
	assert.Equal(t, float64(3.25), result[1].IPKMahasiswa, "Verifikasi IPK mahasiswa entri kedua benar")
	assert.Equal(t, "Bekerja", result[0].StatusSaatIni, "Verifikasi status saat ini entri pertama benar")
	assert.Equal(t, "Universitas Gadjah Mada", result[1].PerguruanTinggiLanjut, "Verifikasi PT Studi Lanjut entri kedua benar")
}
