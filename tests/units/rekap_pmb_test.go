package tests

import (
	"context"
	"restapi-golang/models"
	"restapi-golang/services"
	"restapi-golang/tests/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupStatPMB(peminat, lulus, daftar string) models.StatPMB {
	return models.StatPMB{
		Peminat: peminat,
		Lulus:   lulus,
		Daftar:  daftar,
	}
}

func setupMockRekapPMBData() []models.RekapPMB {
	return []models.RekapPMB{
		{
			ID:        1,
			Tahun:     2024,
			Kode:      "P01",
			NamaProdi: "Informatika",
			SNBP:      setupStatPMB("500", "50", "48"),
			SNBT:      setupStatPMB("1200", "150", "145"),
			SMBJM_CBT: setupStatPMB("300", "60", "55"),
			Jumlah:    setupStatPMB("2000", "260", "248"),
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
		},
		{
			ID:        2,
			Tahun:     2024,
			Kode:      "P02",
			NamaProdi: "Manajemen",
			SNBP:      setupStatPMB("700", "80", "75"),
			SNBT:      setupStatPMB("1500", "200", "190"),
			SMBJM_Rpt: setupStatPMB("400", "70", "68"),
			Pasca:     setupStatPMB("150", "30", "25"),
			Jumlah:    setupStatPMB("2750", "380", "358"),
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

func TestRekapPMBFiltered_Success(t *testing.T) {

	mockData := setupMockRekapPMBData()

	mockRepo := mocks.NewRekapPMBMockRepository(mockData, nil)
	service := services.NewRekapPMBService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetRekapPMBFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetRekapPMBFiltered(
		ctx,
		"", "", "",
		"",
		0, page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Informatika", result[0].NamaProdi, "Verifikasi nama prodi entri pertama benar")
	assert.Equal(t, "1200", result[0].SNBT.Peminat, "Verifikasi jumlah peminat SNBT entri pertama benar")
	assert.Equal(t, "358", result[1].Jumlah.Daftar, "Verifikasi total mahasiswa daftar ulang entri kedua benar")
}
