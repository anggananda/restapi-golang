package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupMockKritikSaranData() []models.KritikSaran {
	return []models.KritikSaran{
		{
			ID:       1,
			NIP:      "198001012005011001",
			Nama:     "Dr. Eko Prasetyo",
			Saran:    []string{"Perluasan akses jurnal internasional.", "Peningkatan fasilitas ruang rapat."},
			Tahun:    "2023",
			Semester: "Ganjil",
			Periode:  "2023/2024",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:       2,
			NIP:      "199005152018022002",
			Nama:     "Fitriani Dewi, M.Si.",
			Saran:    []string{"Sosialisasi peraturan baru lebih ditingkatkan."},
			Tahun:    "2024",
			Semester: "Genap",
			Periode:  "2023/2024",
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

func TestKritikSaranFiltered_Success(t *testing.T) {

	mockData := setupMockKritikSaranData()

	mockRepo := mocks.NewKritikSaranMockRepository(mockData, nil)
	service := services.NewKritikSaranService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetKritikSaranFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetKritikSaranFiltered(
		ctx,
		"", "", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Dr. Eko Prasetyo", result[0].Nama, "Verifikasi nama entri pertama benar")
	assert.Equal(t, 2, len(result[0].Saran), "Verifikasi jumlah saran entri pertama benar")
	assert.Equal(t, "Fitriani Dewi, M.Si.", result[1].Nama, "Verifikasi nama entri kedua benar")
}
