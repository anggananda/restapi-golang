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

func setupMockEvaluasiDosenData() []models.EvaluasiDosen {
	now := time.Now()
	goodScore := models.Penilaian{SangatBaik: 45, Baik: 5, Cukup: 0, Kurang: 0}
	mixedScore := models.Penilaian{SangatBaik: 15, Baik: 20, Cukup: 10, Kurang: 5}
	return []models.EvaluasiDosen{
		{
			ID:              101,
			NIP:             "198001012005011001",
			NoIndukUndiksha: "IND001",
			NamaLengkap:     "Dr. Dewi Kartika, M.Kom.",
			NamaKelas:       "TI-A",
			KodeMatakuliah:  "TIN501",
			NamaMatakuliah:  "Pemrograman Lanjut",
			Tahun:           "2023",
			Semester:        "Ganjil",
			KodeProdi:       "P01",
			KodeFakultas:    "F01",
			Unit: models.Unit{
				FktKode:  "F01",
				JrsKose:  "J01",
				PrdKode:  "P01",
				Fakultas: "Teknik",
				Jurusan:  "Informatika",
			},
			CreatedAt: now,
			UpdatedAt: now,

			PerencanaanPerkuliahan:                     goodScore,
			RelevansiMateriDenganTujuanPembelajaran:    goodScore,
			PenguasaanMateriPerkuliahan:                goodScore,
			MetodeDanPendekatanPerkuliahan:             goodScore,
			InovasiDalamPerkuliahan:                    goodScore,
			KreatifitasDalamPerkuliahan:                goodScore,
			MediaPembelajaran:                          goodScore,
			SumberBelajar:                              goodScore,
			PenilaianHasilBelajar:                      goodScore,
			PenilaianProsesBelajar:                     goodScore,
			PemberianTugasPerkuliahan:                  goodScore,
			PengelolaanKelas:                           goodScore,
			MotivasiDanAntusiasmeMengajar:              goodScore,
			PenciptaanIklimBelajar:                     goodScore,
			Kedisiplinan:                               goodScore,
			PenegakanAturanPerkuliahan:                 goodScore,
			PengembanganKarakterMahasiswa:              goodScore,
			KeteladananDalamBersikapDanBertindak:       goodScore,
			KemampuanBerkomunikasi:                     goodScore,
			PenggunaanBahasaLisanDanTulisan:            goodScore,
			KemampuanBerinteraksiSosialDenganMahasiswa: goodScore,
		},
		{
			ID:              102,
			NIP:             "197505102000032002",
			NoIndukUndiksha: "IND002",
			NamaLengkap:     "Ir. Surya Permana, M.M.",
			NamaKelas:       "EK-B",
			KodeMatakuliah:  "EKO402",
			NamaMatakuliah:  "Manajemen Keuangan",
			Tahun:           "2023",
			Semester:        "Ganjil",
			KodeProdi:       "P02",
			KodeFakultas:    "F02",
			Unit: models.Unit{
				FktKode:  "F02",
				JrsKose:  "J02",
				PrdKode:  "P02",
				Fakultas: "Ekonomi",
				Jurusan:  "Akuntansi",
			},
			CreatedAt: now,
			UpdatedAt: now,

			PerencanaanPerkuliahan:                     mixedScore,
			RelevansiMateriDenganTujuanPembelajaran:    mixedScore,
			PenguasaanMateriPerkuliahan:                mixedScore,
			MetodeDanPendekatanPerkuliahan:             mixedScore,
			InovasiDalamPerkuliahan:                    mixedScore,
			KreatifitasDalamPerkuliahan:                mixedScore,
			MediaPembelajaran:                          mixedScore,
			SumberBelajar:                              mixedScore,
			PenilaianHasilBelajar:                      mixedScore,
			PenilaianProsesBelajar:                     mixedScore,
			PemberianTugasPerkuliahan:                  mixedScore,
			PengelolaanKelas:                           mixedScore,
			MotivasiDanAntusiasmeMengajar:              mixedScore,
			PenciptaanIklimBelajar:                     mixedScore,
			Kedisiplinan:                               mixedScore,
			PenegakanAturanPerkuliahan:                 mixedScore,
			PengembanganKarakterMahasiswa:              mixedScore,
			KeteladananDalamBersikapDanBertindak:       mixedScore,
			KemampuanBerkomunikasi:                     mixedScore,
			PenggunaanBahasaLisanDanTulisan:            mixedScore,
			KemampuanBerinteraksiSosialDenganMahasiswa: mixedScore,
		},
	}
}

func TestEvaluasiDosenFiltered_Success(t *testing.T) {
	mockData := setupMockEvaluasiDosenData()
	mockRepo := mocks.NewEvaluasiDosenMockRepository(mockData, nil)
	service := services.NewEvaluasiDosenService(mockRepo)
	ctx := context.Background()

	page := 1
	limit := 10

	t.Logf("Setup: Mempersiapkan %d mock records untuk pengambilan data berhasil.", len(mockData))
	t.Logf("Call: Memanggil GetEvaluasiDosenFiltered dengan page=%d, limit=%d.", page, limit)

	result, total, err := service.GetEvaluasiDosenFiltered(
		ctx,
		"", "", "", "", "", "",
		"", page, limit,
	)

	t.Logf("Result: Diterima %d records, Total Count: %d.", len(result), total)

	assert.NoError(t, err, "Seharusnya tidak ada error saat sukses")
	assert.Equal(t, int64(2), total, "Total data seharusnya 2")
	assert.Len(t, result, 2, "Panjang hasil seharusnya sama dengan panjang data mock")
	assert.Equal(t, "Dr. Dewi Kartika, M.Kom.", result[0].NamaLengkap, "Verifikasi nama dosen entri pertama benar")
	assert.Equal(t, "Ir. Surya Permana, M.M.", result[1].NamaLengkap, "Verifikasi nama dosen entri kedua benar")
}
