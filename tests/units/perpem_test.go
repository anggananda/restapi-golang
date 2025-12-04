package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockPerpemData() []models.Perpem {
	return []models.Perpem{
		{
			IdPenarawan: "PNW001",
			Kode:        "TI101",
			IdKelas:     "A",
			MK:          "Algoritma dan Struktur Data",
			Kurikulum:   "K2020",
			Pertemuan:   "16",
			Dosen:       []string{"197001011995011001", "198005052008021002"},
			Metode:      "Tatap Muka dan Daring",
			Silabus:     "FileSilabus_TI101.pdf",
			Kontrak:     "FileKontrak_TI101.pdf",
			Rps:         "FileRPS_TI101.pdf",
			Rtm:         "FileRTM_TI101.pdf",
			Semester:    "Ganjil",
			Tahun:       "2023",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			IdPenarawan: "PNW002",
			Kode:        "EK205",
			IdKelas:     "B",
			MK:          "Ekonomi Makro",
			Kurikulum:   "K2016",
			Pertemuan:   "14",
			Dosen:       []string{"198505152010022002"},
			Metode:      "Full Daring",
			Silabus:     "FileSilabus_EK205.pdf",
			Kontrak:     "FileKontrak_EK205.pdf",
			Rps:         "FileRPS_EK205.pdf",
			Rtm:         "FileRTM_EK205.pdf",
			Semester:    "Genap",
			Tahun:       "2024",
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

func TestPerpemFiltered_Success(t *testing.T) {
	mockData := setupMockPerpemData()
	mockRepo := mocks.NewPerpemMockRepository(mockData, nil)
	service := services.NewPerpemService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetPerpemFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetPerpemFiltered(
		ctx,
		"", "", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "TI101", result[0].Kode, "Verifikasi kode matakuliah entri pertama benar")
	assert.Equal(t, "Ekonomi Makro", result[1].MK, "Verifikasi nama matakuliah entri kedua benar")
	assert.Len(t, result[0].Dosen, 2, "Verifikasi jumlah dosen pengampu entri pertama benar")
}
