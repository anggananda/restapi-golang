package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockRealisasiBulanData() []models.RealisasiBulan {
	return []models.RealisasiBulan{
		{
			ID:                    1,
			Kode:                  "RB202401",
			Bulan:                 "01",
			RealisasiTotalPNBP:    "500000000",
			RealisasiTotalRM:      "1000000000",
			RealisasiTotalRMBOPTN: "1500000000",
			TahunAnggaran:         "2024",
		},
		{
			ID:                    2,
			Kode:                  "RB202402",
			Bulan:                 "02",
			RealisasiTotalPNBP:    "650000000",
			RealisasiTotalRM:      "900000000",
			RealisasiTotalRMBOPTN: "1450000000",
			TahunAnggaran:         "2024",
		},
	}
}

func TestRealisasiBulanFiltered_Success(t *testing.T) {
	mockData := setupMockRealisasiBulanData()
	mockRepo := mocks.NewRealisasiBulanMockRepository(mockData, nil)
	service := services.NewRealisasiBulanService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetRealisasiBulanFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetRealisasiBulanFiltered(
		ctx,
		"", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "500000000", result[0].RealisasiTotalPNBP, "Verifikasi realisasi PNBP entri pertama benar")
	assert.Equal(t, "1450000000", result[1].RealisasiTotalRMBOPTN, "Verifikasi realisasi RM BOPTN entri kedua benar")
	assert.Equal(t, "2024", result[0].TahunAnggaran, "Verifikasi Tahun Anggaran entri pertama benar")
}
