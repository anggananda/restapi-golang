package constants

import (
	"os"
	"restapi-golang/models"
)

func GetPrefillRedisKeys() []models.PrefillEntry {
	return []models.PrefillEntry{
		{
			Key:   "unit_kerja",
			Value: os.Getenv("UNIT_KERJA"),
			TTL:   0,
		},
		{
			Key:   "status_mahasiswa",
			Value: os.Getenv("STATUS_MAHASISWA"),
			TTL:   0,
		},
		{
			Key:   "status_pegawai",
			Value: os.Getenv("STATUS_PEGAWAI"),
			TTL:   0,
		},
		{
			Key:   "status_keaktifan_pegawai",
			Value: os.Getenv("STATUS_KEAKTIFAN_PEGAWAI"),
			TTL:   0,
		},
	}
}
