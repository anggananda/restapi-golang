package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockRealisasiUnitData() []models.RealisasiUnit {
	return []models.RealisasiUnit{
		{
			ID:               1,
			Kode:             "RU2024F01",
			KodeUnit:         "F01",
			NamaUnit:         "Fakultas Teknik",
			PaguPNBP:         "800000000",
			PaguRM:           "1500000000",
			PaguRMBOPTN:      "2300000000",
			RealisasiPNBP:    "750000000",
			RealisasiRM:      "1400000000",
			RealisasiRMBOPTN: "2150000000",
			TahunAnggaran:    "2024",
		},
		{
			ID:               2,
			Kode:             "RU2024F02",
			KodeUnit:         "F02",
			NamaUnit:         "Fakultas Ekonomi dan Bisnis",
			PaguPNBP:         "600000000",
			PaguRM:           "1200000000",
			PaguRMBOPTN:      "1800000000",
			RealisasiPNBP:    "550000000",
			RealisasiRM:      "1100000000",
			RealisasiRMBOPTN: "1650000000",
			TahunAnggaran:    "2024",
		},
	}
}

func TestRealisasiUnitFiltered_Success(t *testing.T) {
	mockData := setupMockRealisasiUnitData()
	mockRepo := mocks.NewRealisasiUnitMockRepository(mockData, nil)
	service := services.NewRealisasiUnitService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetRealisasiUnitFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetRealisasiUnitFiltered(
		ctx,
		"", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Fakultas Teknik", result[0].NamaUnit, "Verifikasi nama unit entri pertama benar")
	assert.Equal(t, "750000000", result[0].RealisasiPNBP, "Verifikasi realisasi PNBP entri pertama benar")
	assert.Equal(t, "1650000000", result[1].RealisasiRMBOPTN, "Verifikasi realisasi RM BOPTN entri kedua benar")
}
