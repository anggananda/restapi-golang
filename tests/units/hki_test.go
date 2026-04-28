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

func setupMockHkiData() []models.Hki {
	now := time.Now().Format(time.RFC3339)

	return []models.Hki{
		{
			ID:                     1,
			NamaKarya:              "Sistem Otomasi Irigasi Berbasis IoT",
			Tanggal:                "2023-10-01",
			CreatedAt:              now,
			IsValid:                "Y",
			Semester:               "Ganjil",
			TahunAjaran:            "2023/2024",
			TahunData:              "2023",
			JenisPaten:             "Paten Sederhana",
			TbName:                 "hki",
			PrimaryKey:             "1",
			WaktuPelaksanaan:       "2023-01-01 s/d 2023-12-31",
			NamaDosen:              "Dr. Ahmad Riyadi",
			NoPendaftaran:          "S00202300001",
			NoPendatatanSertifikat: "IDP000012345",
			Scope:                  "Nasional",
			JmlNegaraPengaku:       "1",
			FileBuktiKinerja:       "bukti_kinerja_1.pdf",
			FileSertifikatPaten:    "sertifikat_paten_1.pdf",
			Deskripsi:              "Invensi di bidang pertanian cerdas.",
			Posisi:                 "Inventor Utama",
			JmlPenulis:             "3",
			IsProduk:               "Y",
			ProdukPenelitianJudul:  "Penelitian Pertanian Cerdas",
			ProdukPenelitianID:     "PNL001",
			ValidIpk:               "Y",
			CreateDosenID:          "DSN001",
			// LevelCapaian:           "1",
			SumberProduk:   "Penelitian",
			KodeScope:      "SC01",
			KodeJenisPaten: "JP01",
			Periode:        "2023/2024",
			SemesterType:   "Ganjil",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:                     2,
			NamaKarya:              "Aplikasi Pembelajaran Bahasa Jepang Interaktif",
			Tanggal:                "2024-03-15",
			CreatedAt:              now,
			IsValid:                "Y",
			Semester:               "Genap",
			TahunAjaran:            "2023/2024",
			TahunData:              "2024",
			JenisPaten:             "Hak Cipta",
			TbName:                 "hki",
			PrimaryKey:             "2",
			WaktuPelaksanaan:       "2024-01-01 s/d 2024-12-31",
			NamaDosen:              "Dr. Siti Rahayu",
			NoPendaftaran:          "C00202400002",
			NoPendatatanSertifikat: "IDH000056789",
			Scope:                  "Internasional",
			JmlNegaraPengaku:       "3",
			FileBuktiKinerja:       "bukti_kinerja_2.pdf",
			FileSertifikatPaten:    "sertifikat_hakcipta_2.pdf",
			Deskripsi:              "Program komputer sebagai alat bantu mengajar bahasa.",
			Posisi:                 "Pencipta Pertama",
			JmlPenulis:             "1",
			IsProduk:               "N",
			ValidIpk:               "Y",
			CreateDosenID:          "DSN002",
			// LevelCapaian:           "2",
			SumberProduk:   "Mandiri",
			KodeScope:      "SC02",
			KodeJenisPaten: "JP02",
			Periode:        "2023/2024",
			SemesterType:   "Genap",
			Unit: models.Unit{
				FktKode:  "F03",
				JrsKose:  "J03",
				PrdKode:  "P03",
				Fakultas: "Bahasa",
				Jurusan:  "Sastra Jepang",
			},
		},
	}
}

func TestHkiFiltered_Success(t *testing.T) {

	mockData := setupMockHkiData()

	mockRepo := mocks.NewHkiMockRepository(mockData, nil)
	service := services.NewHkiService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetHkiFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetHkiFiltered(
		ctx,
		"", "", "", "", "",
		"", page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Sistem Otomasi Irigasi Berbasis IoT", result[0].NamaKarya, "Verifikasi nama karya HKI entri pertama benar")
	assert.Equal(t, "Dr. Siti Rahayu", result[1].NamaDosen, "Verifikasi nama dosen HKI entri kedua benar")
}
