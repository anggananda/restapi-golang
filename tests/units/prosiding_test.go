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

func setupMockProsidingData() []models.Prosiding {
	now := time.Now().Format("2006-01-02 15:04:05")
	return []models.Prosiding{
		{
			ID:                1,
			JudulArtikel:      "Optimasi Algoritma Penjadwalan Cloud Computing",
			NamaSeminar:       "ICON-IT International Conference on Information Technology",
			TglAwal:           "2023-10-25",
			CreatedAt:         now,
			IsValid:           "Y",
			Semester:          "Ganjil",
			TahunAjaran:       "2023/2024",
			TahunData:         "2023",
			TbName:            "prosiding",
			PrimaryKey:        "1",
			WaktuPelaksanaan:  "2023-10-25 s/d 2023-10-26",
			NamaDosen:         "Dr. Budi Santoso",
			Scope:             "Internasional",
			TipeProsiding:     "Full Paper",
			Penyelenggara:     "Universitas X",
			Penerbit:          "IEEE Xplore",
			EIssn:             "2355-5085",
			Bereputasi:        "Y",
			Posisi:            "Penulis Utama",
			JmlPenulis:        "4",
			KodeTipeProsiding: "TP01",
			KodeScope:         "SC01",
			Periode:           "2023/2024",
			SemesterType:      "Ganjil",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:                2,
			JudulArtikel:      "Pengaruh Inflasi terhadap Daya Beli Masyarakat di Pedesaan",
			NamaSeminar:       "Seminar Nasional Ekonomi dan Bisnis (SNEB) 2024",
			TglAwal:           "2024-05-10",
			CreatedAt:         now,
			IsValid:           "Y",
			Semester:          "Genap",
			TahunAjaran:       "2023/2024",
			TahunData:         "2024",
			TbName:            "prosiding",
			PrimaryKey:        "2",
			WaktuPelaksanaan:  "2024-05-10 s/d 2024-05-10",
			NamaDosen:         "Dr. Mira Kirana",
			Scope:             "Nasional",
			TipeProsiding:     "Poster",
			Penyelenggara:     "Asosiasi Dosen Ekonomi Indonesia",
			Penerbit:          "Saintek Press",
			PIssn:             "1412-2580",
			Bereputasi:        "N",
			Posisi:            "Penulis Anggota",
			JmlPenulis:        "2",
			KodeTipeProsiding: "TP02",
			KodeScope:         "SC02",
			Periode:           "2023/2024",
			SemesterType:      "Genap",
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

func TestProsidingFiltered_Success(t *testing.T) {
	mockData := setupMockProsidingData()
	mockRepo := mocks.NewProsidingMockRepository(mockData, nil)
	service := services.NewProsidingService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetProsidingFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetProsidingFiltered(
		ctx,
		"", "", "", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Optimasi Algoritma Penjadwalan Cloud Computing", result[0].JudulArtikel, "Verifikasi judul artikel entri pertama benar")
	assert.Equal(t, "Dr. Mira Kirana", result[1].NamaDosen, "Verifikasi nama dosen entri kedua benar")
	assert.Equal(t, "IEEE Xplore", result[0].Penerbit, "Verifikasi penerbit entri pertama benar")
	assert.Equal(t, "Y", result[0].Bereputasi, "Verifikasi status bereputasi entri pertama benar")
}
