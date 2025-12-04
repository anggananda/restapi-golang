package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockPenawaranData() []models.Penawaran {
	return []models.Penawaran{
		{
			ID:             1,
			JmlMhsAmbil:    "45",
			KodeMatakuliah: "TI101",
			Kurikulum:      "K2020",
			NamaKelas:      "A",
			NamaMatakuliah: "Algoritma dan Struktur Data",
			NamaPengampu:   "Prof. Dr. Anton Yudha",
			NipPengampu:    "197001011995011001",
			Pengampu:       "Anton Yudha",
			Semester:       "Ganjil",
			Tahun:          "2023",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:             2,
			JmlMhsAmbil:    "30",
			KodeMatakuliah: "EK205",
			Kurikulum:      "K2016",
			NamaKelas:      "B",
			NamaMatakuliah: "Ekonomi Makro",
			NamaPengampu:   "Dr. Rina Wulandari",
			NipPengampu:    "198505152010022002",
			Pengampu:       "Rina Wulandari",
			Semester:       "Genap",
			Tahun:          "2024",
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

func TestPenawaranFiltered_Success(t *testing.T) {

	mockData := setupMockPenawaranData()

	mockRepo := mocks.NewPenawaranMockRepository(mockData, nil)
	service := services.NewPenawaranService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetPenawaranFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetPenawaranFiltered(
		ctx,
		"", "", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "TI101", result[0].KodeMatakuliah, "Verifikasi kode matakuliah entri pertama benar")
	assert.Equal(t, "Ekonomi Makro", result[1].NamaMatakuliah, "Verifikasi nama matakuliah entri kedua benar")
	assert.Equal(t, "Prof. Dr. Anton Yudha", result[0].NamaPengampu, "Verifikasi nama pengampu entri pertama benar")
}
