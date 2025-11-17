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
	}
}
