package constants

import (
	"os"
	"restapi-golang/models"
)

var PrefillRedisKeys = []models.PrefillEntry{
	{
		Key:   "struktur_unit",
		Value: os.Getenv("UNIT_KERJA"),
		TTL:   0,
	},
}
