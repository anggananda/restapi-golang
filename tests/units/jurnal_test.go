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

func setupMockJurnalData() []models.Jurnal {
	now := time.Now().Format(time.RFC3339)
	deletedAt := new(string)
	*deletedAt = ""

	return []models.Jurnal{
		{
			ID:                    1,
			JudulArtikel:          "Aplikasi Deep Learning untuk Deteksi Penyakit Tanaman Padi",
			NamaJurnal:            "Jurnal Informatika Pertanian",
			Tanggal:               "2023-11-20",
			Semester:              "Ganjil",
			TahunAjaran:           "2023/2024",
			TahunData:             "2023",
			Akreditasi:            "Sinta 2",
			CreatedAt:             now,
			IsValid:               "Y",
			TbName:                "jurnal",
			PrimaryKey:            "1",
			WaktuPelaksanaan:      "2023-01-01 s/d 2023-12-31",
			NamaDosen:             "Prof. Dr. Anton Yudha",
			KodeProdi:             "P01",
			NamaJurusan:           "Informatika",
			NamaFakultas:          "Teknik",
			IdSinta:               "60133",
			Authors:               "Anton Yudha, Budi Hartono",
			Sitasi:                "15",
			VolumeJurnal:          "10",
			NomorJurnal:           "2",
			HalamanAwal:           "100",
			HalamanAkhir:          "115",
			PIssn:                 "2087-5555",
			EIssn:                 "2087-6666",
			Doi:                   "10.1234/jip.v10i2.501",
			Penerbit:              "Universitas Teknologi",
			AlamatWebJurnal:       "http://jurnal-teknologi.ac.id",
			BahasaID:              "Indonesia",
			Sinta:                 "2",
			Scope:                 "Nasional",
			JenisJurnal:           "Jurnal Ilmiah",
			AggregationType:       "Journal",
			TahunPublish:          "2023",
			Posisi:                "Penulis Pertama",
			JmlPenulis:            "2",
			UpdatedAt:             now,
			DeletedAt:             nil,
			IsProduk:              "Y",
			ProdukPenelitianJudul: "Riset Unggulan AI Pertanian",
			ProdukPenelitianID:    "PNT001",
			StatusLengkap:         "Tersimpan",
			ValidIpk:              "Y",
			CreateDosenID:         "DSN001",
			// Indexer:               []string{"Sinta", "Google Scholar"},
			SumberProduk:          "Penelitian",
			KodeAkreditasi:        "S2",
			KodeScope:             "SC01",
			KodeJenisJurnal:       "JJ01",
			Periode:               "2023/2024",
			SemesterType:          "Ganjil",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:               2,
			JudulArtikel:     "The Role of ESG in Sustainable Investment Strategies",
			NamaJurnal:       "Journal of Global Finance",
			Tanggal:          "2024-01-10",
			Semester:         "Genap",
			TahunAjaran:      "2023/2024",
			TahunData:        "2024",
			Akreditasi:       "Scopus Q1",
			CreatedAt:        now,
			IsValid:          "Y",
			TbName:           "jurnal",
			PrimaryKey:       "2",
			WaktuPelaksanaan: "2024-01-01 s/d 2024-12-31",
			NamaDosen:        "Dr. Rina Wulandari",
			KodeProdi:        "P02",
			NamaJurusan:      "Akuntansi",
			NamaFakultas:     "Ekonomi",
			IdSinta:          "",
			Authors:          "Rina Wulandari, J. Smith",
			Sitasi:           "42",
			VolumeJurnal:     "5",
			NomorJurnal:      "1",
			HalamanAwal:      "50",
			HalamanAkhir:     "65",
			PIssn:            "1234-5678",
			EIssn:            "8765-4321",
			Doi:              "10.5678/jgf.v5i1.2024.005",
			Penerbit:         "Global Academic Press",
			AlamatWebJurnal:  "http://journal-global-finance.com",
			BahasaID:         "English",
			Sinta:            "",
			Scope:            "Internasional",
			JenisJurnal:      "Jurnal Internasional",
			AggregationType:  "Journal",
			TahunPublish:     "2024",
			Posisi:           "Penulis Korespondensi",
			JmlPenulis:       "2",
			UpdatedAt:        now,
			DeletedAt:        deletedAt,
			IsProduk:         "N",
			StatusLengkap:    "Terkirim",
			ValidIpk:         "Y",
			CreateDosenID:    "DSN002",
			// Indexer:          []string{"Scopus", "Web of Science"},
			SumberProduk:     "Mandiri",
			KodeAkreditasi:   "Q1",
			KodeScope:        "SC02",
			KodeJenisJurnal:  "JJ02",
			Periode:          "2023/2024",
			SemesterType:     "Genap",
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

func TestJurnalFiltered_Success(t *testing.T) {

	mockData := setupMockJurnalData()

	mockRepo := mocks.NewJurnalMockRepository(mockData, nil)
	service := services.NewJurnalService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetJurnalFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetJurnalFiltered(
		ctx,
		"", "", "", "", "",
		"", "", "",
		page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Aplikasi Deep Learning untuk Deteksi Penyakit Tanaman Padi", result[0].JudulArtikel, "Verifikasi judul artikel entri pertama benar")
	assert.Equal(t, "Dr. Rina Wulandari", result[1].NamaDosen, "Verifikasi nama dosen entri kedua benar")
}
