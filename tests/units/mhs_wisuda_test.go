package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockMhsWisudaData() []models.MhsWisuda {
	return []models.MhsWisuda{
		{
			ID:          1,
			NIM:         "1805551001",
			NamaLengkap: "Dewa Made Suta",
			TahunWisuda: 2023,
			BulanWisuda: 10,
			NamaBulan:   "Oktober",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:          2,
			NIM:         "1905552002",
			NamaLengkap: "Ni Luh Putu Ayu",
			TahunWisuda: 2024,
			BulanWisuda: 4,
			NamaBulan:   "April",
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

func TestMhsWisudaFiltered_Success(t *testing.T) {

	mockData := setupMockMhsWisudaData()

	mockRepo := mocks.NewMhsWisudaMockRepository(mockData, nil)
	service := services.NewMhsWisudaService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetMhsWisudaFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetMhsWisudaFiltered(
		ctx,
		"", "", "", "",
		0, 0, page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Dewa Made Suta", result[0].NamaLengkap, "Verifikasi nama lengkap entri pertama benar")
	assert.Equal(t, 2024, result[1].TahunWisuda, "Verifikasi tahun wisuda entri kedua benar")
}
