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

func setupMockPenelitianData() []models.Penelitian {
	now := time.Now().Format("2006-01-02 15:04:05")
	return []models.Penelitian{
		{
			ID:                  1,
			JudulTulisan:        "Pengembangan Model Prediksi Curah Hujan Berbasis Machine Learning",
			Tanggal:             "2023-08-01",
			CreatedAt:           now,
			IsValid:             "Y",
			Semester:            "Ganjil",
			TahunAjaran:         "2023/2024",
			TahunData:           "2023",
			TbName:              "penelitian",
			PrimaryKey:          "1",
			WaktuPelaksanaan:    "2023-08-01 s/d 2024-07-31",
			NamaDosen:           "Prof. Dr. Anton Yudha",
			SumberDana:          "Riset Dikti",
			IdSkim:              "SKM01",
			Skema:               "Penelitian Dasar Unggulan Perguruan Tinggi",
			Dana:                "75000000",
			JenisPenelitian:     "Dasar",
			InstitusiSumberDana: "Kementerian Pendidikan, Kebudayaan, Riset, dan Teknologi",
			Deskripsi:           "Penelitian untuk memprediksi curah hujan menggunakan algoritma Deep Learning.",
			Posisi:              "Ketua Peneliti",
			StatusLengkap:       "Selesai",
			RumpunIlmu:          "Ilmu Komputer",
			ValidIpk:            "Y",
			BidangPenelitian:    "Kecerdasan Buatan",
			// LevelCapaian:        "Nasional",
			Periode:             "2023/2024",
			SemesterType:        "Ganjil",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:                  2,
			JudulTulisan:        "Analisis Dampak Regulasi Fiskal terhadap Perekonomian Lokal",
			Tanggal:             "2024-02-15",
			CreatedAt:           now,
			IsValid:             "Y",
			Semester:            "Genap",
			TahunAjaran:         "2023/2024",
			TahunData:           "2024",
			TbName:              "penelitian",
			PrimaryKey:          "2",
			WaktuPelaksanaan:    "2024-03-01 s/d 2025-02-28",
			NamaDosen:           "Dr. Rina Wulandari",
			SumberDana:          "Internal",
			IdSkim:              "SKM02",
			Skema:               "Penelitian Mandiri Dosen",
			Dana:                "15000000",
			JenisPenelitian:     "Terapan",
			InstitusiSumberDana: "Universitas X",
			Deskripsi:           "Kajian mengenai kebijakan fiskal dan dampaknya pada UMKM di daerah.",
			Posisi:              "Peneliti Anggota",
			StatusLengkap:       "Berjalan",
			RumpunIlmu:          "Ekonomi",
			ValidIpk:            "Y",
			BidangPenelitian:    "Ekonomi Moneter",
			// LevelCapaian:        "Lokal",
			Periode:             "2023/2024",
			SemesterType:        "Genap",
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

func TestPenelitianFiltered_Success(t *testing.T) {

	mockData := setupMockPenelitianData()

	mockRepo := mocks.NewPenelitianMockRepository(mockData, nil)
	service := services.NewPenelitianService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetPenelitianFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetPenelitianFiltered(
		ctx,
		"", "", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Pengembangan Model Prediksi Curah Hujan Berbasis Machine Learning", result[0].JudulTulisan, "Verifikasi judul entri pertama benar")
	assert.Equal(t, "Dr. Rina Wulandari", result[1].NamaDosen, "Verifikasi nama dosen entri kedua benar")
	assert.Equal(t, "Riset Dikti", result[0].SumberDana, "Verifikasi sumber dana entri pertama benar")
}
