package mocks

import (
	"log"
)

func loadMock(path, key string) any {
	data, err := LoadYAMLSection(path, key)
	if err != nil {
		log.Printf("⚠️ Gagal memuat mock YAML [%s - %s]: %v", path, key, err)
		return fallbackMock(key)
	}
	return data
}
