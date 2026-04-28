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

func setupMockPengabdianData() []models.Pengabdian {
	now := time.Now().Format("2006-01-02 15:04:05")
	return []models.Pengabdian{
		{
			ID:                  1,
			NamaKegiatan:        "Pelatihan Literasi Digital untuk UMKM Desa Sukamaju",
			Tanggal:             "2023-11-05",
			CreatedAt:           now,
			Semester:            "Ganjil",
			TahunAjaran:         "2023/2024",
			TahunData:           "2023",
			IsValid:             "Y",
			TbName:              "pengabdian",
			PrimaryKey:          "1",
			WaktuPelaksanaan:    "2023-11-05 s/d 2023-11-20",
			NamaDosen:           "Dr. Siti Rahayu",
			SumberDana:          "Dana Internal Fakultas",
			IDSkim:              "PKM01",
			Skema:               "Skema Pemberdayaan Masyarakat",
			Dana:                "10000000",
			JenisPengabdian:     "Pemberdayaan",
			InstitusiSumberDana: "Fakultas Ekonomi dan Bisnis",
			Deskripsi:           "Memberikan bimbingan kepada pelaku UMKM dalam memanfaatkan media digital untuk pemasaran.",
			Posisi:              "Ketua Pelaksana",
			StatusLengkap:       "Selesai",
			RumpunIlmu:          "Sosial",
			ValidIpk:            "Y",
			BidangPenelitian:    "Pengembangan Masyarakat",
			// LevelCapaian:        "Lokal",
			Periode:      "2023/2024",
			SemesterType: "Ganjil",
			Unit: models.Unit{
				FktKode:  "F02",
				JrsKose:  "J02",
				PrdKode:  "P02",
				Fakultas: "Ekonomi",
				Jurusan:  "Manajemen",
			},
		},
		{
			ID:                  2,
			NamaKegiatan:        "Sosialisasi Bahaya Sampah Plastik di Sekolah Dasar",
			Tanggal:             "2024-04-10",
			CreatedAt:           now,
			Semester:            "Genap",
			TahunAjaran:         "2023/2024",
			TahunData:           "2024",
			IsValid:             "Y",
			TbName:              "pengabdian",
			PrimaryKey:          "2",
			WaktuPelaksanaan:    "2024-04-10 s/d 2024-04-10",
			NamaDosen:           "Prof. Dr. Anton Yudha",
			SumberDana:          "Hibah Ristekdikti",
			IDSkim:              "PKM02",
			Skema:               "Skema Penerapan Teknologi Tepat Guna",
			Dana:                "50000000",
			JenisPengabdian:     "Penerapan",
			InstitusiSumberDana: "Kementerian Pendidikan, Kebudayaan, Riset, dan Teknologi",
			Deskripsi:           "Edukasi tentang pengelolaan sampah plastik dan daur ulang kepada siswa SD.",
			Posisi:              "Anggota Pelaksana",
			StatusLengkap:       "Selesai",
			RumpunIlmu:          "Teknik",
			ValidIpk:            "Y",
			BidangPenelitian:    "Lingkungan",
			// LevelCapaian:        "Nasional",
			Periode:      "2023/2024",
			SemesterType: "Genap",
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

func TestPengabdianFiltered_Success(t *testing.T) {
	mockData := setupMockPengabdianData()
	mockRepo := mocks.NewPengabdianMockRepository(mockData, nil)
	service := services.NewPengabdianService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetPengabdianFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetPengabdianFiltered(
		ctx,
		"", "", "", "", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Pelatihan Literasi Digital untuk UMKM Desa Sukamaju", result[0].NamaKegiatan, "Verifikasi nama kegiatan entri pertama benar")
	assert.Equal(t, "Prof. Dr. Anton Yudha", result[1].NamaDosen, "Verifikasi nama dosen entri kedua benar")
	assert.Equal(t, "Dana Internal Fakultas", result[0].SumberDana, "Verifikasi sumber dana entri pertama benar")
}
