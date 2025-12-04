package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockAngketMhsData() []models.AngketMhs {
	return []models.AngketMhs{
		{
			ID:          1,
			Dosen:       []string{"Dr. Indra K.", "Prof. Siti A."},
			IdKelas:     "TI-A",
			IdPenawaran: "P001",
			Kode:        "TIN401",
			Mk:          "Algoritma dan Struktur Data",
			NipDosen:    []string{"198010102005011001", "197505052000032002"},
			Periode:     "Ganjil",
			Semester:    "5",
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
			ID:          2,
			Dosen:       []string{"Dr. Budi S."},
			IdKelas:     "EK-B",
			IdPenawaran: "P002",
			Kode:        "EKO305",
			Mk:          "Ekonomi Manajerial",
			NipDosen:    []string{"198501012010011005"},
			Periode:     "Ganjil",
			Semester:    "3",
			Tahun:       "2023",
			Unit: models.Unit{
				FktKode:  "F02",
				JrsKose:  "J02",
				PrdKode:  "P02",
				Fakultas: "Ekonomi",
				Jurusan:  "Akuntansi",
			},
		},
		{
			ID:          3,
			Dosen:       []string{"Ir. Rina D.", "Dr. Budi S."},
			IdKelas:     "TI-C",
			IdPenawaran: "P003",
			Kode:        "TIN102",
			Mk:          "Matematika Dasar",
			NipDosen:    []string{"197802022004022003", "198501012010011005"},
			Periode:     "Genap",
			Semester:    "1",
			Tahun:       "2022",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
	}
}

func TestGetAngketMhsFiltered__Success(t *testing.T) {
	mockData := setupMockAngketMhsData()
	mockRepo := mocks.NewAngketMhsMockRepository(mockData, nil)
	service := services.NewAngketMhsService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Preparing %d mock records for successful retrieval.", len(mockData))
	t.Logf("Call: GetAngketMhsFiltered with page=%d, limit=%d.", page, limit)

	result, total, err := service.GetAngketMhsFiltered(ctx, "", "", "", "", "", "", page, limit)

	t.Logf("Result: Received %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Should be no error on success")
	assert.Equal(t, int64(3), total, "Total data should be 3")
	assert.Len(t, result, 3, "Result length should match mock data length")
	assert.Equal(t, "TIN401", result[0].Kode, "Verify the first data entry is correct")
}
