package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockAgendaMengajarData() []models.AgendaMengajar {
	return []models.AgendaMengajar{
		{
			ID:          1,
			Dosen:       []string{"Dr. Wira S.", "M.Sc."},
			IdKelas:     "TI-A",
			IdPenawaran: "OFR-TI-A01",
			JenisKelas:  "Teori",
			Kode:        "TIN601",
			Kurikulum:   "2021",
			Matakuliah:  "Pengembangan Aplikasi Mobile",
			NipDosen:    []string{"198501202010011003"},
			Periode:     "Ganjil",
			Pertemuan:   "3", // Pertemuan ke-3
			Semester:    "5",
			Sumber:      "SIMAK",
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
			Dosen:       []string{"Dr. Maya P.", "Ak., M.M."},
			IdKelas:     "EK-B",
			IdPenawaran: "OFR-EK-B05",
			JenisKelas:  "Praktikum",
			Kode:        "EKO302",
			Kurikulum:   "2018",
			Matakuliah:  "Akuntansi Biaya Lanjutan",
			NipDosen:    []string{"197805152000032001"},
			Periode:     "Genap",
			Pertemuan:   "10", // Pertemuan ke-10
			Semester:    "4",
			Sumber:      "SIMAK",
			Tahun:       "2024",
			Unit: models.Unit{
				FktKode:  "F02",
				JrsKose:  "J02",
				PrdKode:  "P02",
				Fakultas: "Ekonomi",
				Jurusan:  "Akuntansi",
			},
		},
	}
}

func TestAgendaMengajarFiltered_Success(t *testing.T) {
	mockData := setupMockAgendaMengajarData()
	mockRepo := mocks.NewAgendaMengajarMockRepository(mockData, nil)
	service := services.NewAgendaMengajarService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Preparing %d mock records for successful retrieval.", len(mockData))
	t.Logf("Call: GetAngketMhsFiltered with page=%d, limit=%d.", page, limit)

	result, total, err := service.GetAgendaMengajarFiltered(ctx, "", "", "", "", "", "", page, limit)

	t.Logf("Result: Received %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Should be no error on success")
	assert.Equal(t, int64(2), total, "Total data should be 3")
	assert.Len(t, result, 2, "Result length should match mock data length")
	assert.Equal(t, "TIN601", result[0].Kode, "Verify the first data entry is correct")
}
