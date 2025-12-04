package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockBeasiswaData() []models.Beasiswa {
	return []models.Beasiswa{
		{
			ID:            1,
			NIM:           "B901001",
			Nama:          "Rani Puspita",
			JenisBeasiswa: "PPA",
			IPK:           3.85,
			Status:        "Diterima",
			Tahun:         2023,
			Semester:      "7",
			SemesterType:  "Ganjil",
			Periode:       "2023/2024",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:            2,
			NIM:           "B902005",
			Nama:          "Dion Saputra",
			JenisBeasiswa: "Unggulan",
			IPK:           3.21,
			Status:        "Pending",
			Tahun:         2024,
			Semester:      "3",
			SemesterType:  "Genap",
			Periode:       "2023/2024",
			Unit: models.Unit{
				FktKode:  "F03",
				JrsKose:  "J03",
				PrdKode:  "P03",
				Fakultas: "Hukum",
				Jurusan:  "Hukum Bisnis",
			},
		},
	}
}

func TestBeasiswaFiltered_Success(t *testing.T) {
	mockData := setupMockBeasiswaData()
	mockRepo := mocks.NewBeasiswaMockRepository(mockData, nil)
	service := services.NewBeasiswaService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Preparing %d mock records for successful retrieval.", len(mockData))
	t.Logf("Call: GetAngketMhsFiltered with page=%d, limit=%d.", page, limit)

	result, total, err := service.GetBeasiswaFiltered(ctx, "", "", "", "", "", "", 0, page, limit)

	t.Logf("Result: Received %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Should be no error on success")
	assert.Equal(t, int64(2), total, "Total data should be 3")
	assert.Len(t, result, 2, "Result length should match mock data length")
	assert.Equal(t, "Rani Puspita", result[0].Nama, "Verify the first data entry is correct")
}
