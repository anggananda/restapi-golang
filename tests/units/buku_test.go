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

func setupMockBukuData() []models.Buku {
	now := time.Now().Format(time.RFC3339)

	return []models.Buku{
		{
			ID:                  1,
			JudulBuku:           "Dasar-Dasar Pemrograman Go",
			Penerbit:            "Penerbit Informatika",
			Tanggal:             "2023-01-15",
			CreatedAt:           now,
			IsValid:             "Y",
			Semester:            "Ganjil",
			TahunAjaran:         "2023/2024",
			TahunData:           "2023",
			Scope:               "Nasional",
			TbName:              "buku",
			PrimaryKey:          "1",
			WaktuPelaksanaan:    "2023-01-01 s/d 2023-12-31",
			NamaDosen:           "Dr. Budi Santoso",
			KodeProdi:           "P01",
			NamaJurusan:         "Informatika",
			NamaFakultas:        "Teknik",
			Keterangan:          "Buku ajar untuk mata kuliah Dasar Pemrograman",
			KategoriBuku:        "Textbook",
			ISBN:                "978-602-03-3456-7",
			UrlDokumen:          "http://repo.com/dok/buku_go_1.pdf",
			UrlPerReview:        "http://repo.com/review/buku_go_1.pdf",
			Satuan:              "Eksemplar",
			JumlahHalaman:       "350",
			VolumeKegiatan:      "1",
			FileUpload:          "file_go_1.zip",
			Posisi:              "Penulis Utama",
			JmlNegaraPengedaran: "1",
			JmlPenulis:          "2",
			IsProduk:            "N",
			LevelCapaian:        "1",
			SumberProduk:        "Mandiri",
			KodeKategoriBuku:    "KB01",
			KodeScope:           "SC01",
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
			ID:                    2,
			JudulBuku:             "Akuntansi Lanjutan untuk Bisnis Modern",
			Penerbit:              "Erlangga",
			Tanggal:               "2024-03-10",
			CreatedAt:             now,
			IsValid:               "Y",
			Semester:              "Genap",
			TahunAjaran:           "2023/2024",
			TahunData:             "2024",
			Scope:                 "Internasional",
			TbName:                "buku",
			PrimaryKey:            "2",
			WaktuPelaksanaan:      "2024-01-01 s/d 2024-12-31",
			NamaDosen:             "Prof. Ani Lestari",
			KodeProdi:             "P02",
			NamaJurusan:           "Akuntansi",
			NamaFakultas:          "Ekonomi",
			Keterangan:            "Buku referensi Akuntansi Keuangan.",
			KategoriBuku:          "Buku Referensi",
			ISBN:                  "978-979-03-9876-5",
			UrlDokumen:            "http://repo.com/dok/buku_ak_2.pdf",
			UrlPerReview:          "http://repo.com/review/buku_ak_2.pdf",
			Satuan:                "Eksemplar",
			JumlahHalaman:         "520",
			VolumeKegiatan:        "1",
			FileUpload:            "file_ak_2.zip",
			Posisi:                "Penulis Pendamping",
			JmlNegaraPengedaran:   "3",
			JmlPenulis:            "3",
			IsProduk:              "Y",
			ProdukPenelitianJudul: "Laporan Penelitian A",
			ProdukPenelitianID:    "PRD001",
			LevelCapaian:          "3",
			SumberProduk:          "Penelitian",
			KodeKategoriBuku:      "KB02",
			KodeScope:             "SC02",
			Periode:               "2023/2024",
			SemesterType:          "Genap",
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

func TestBukuFiltered_Success(t *testing.T) {
	mockData := setupMockBukuData()

	mockRepo := mocks.NewBukuMockRepository(mockData, nil)
	service := services.NewBukuService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetBukuFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetBukuFiltered(
		ctx,
		"", "", "", "", "",
		"", page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Dasar-Dasar Pemrograman Go", result[0].JudulBuku, "Verifikasi judul buku entri pertama benar")
	assert.Equal(t, "Prof. Ani Lestari", result[1].NamaDosen, "Verifikasi nama dosen entri kedua benar")
}
