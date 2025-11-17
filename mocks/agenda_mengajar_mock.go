package mocks

import (
	"fmt"
	"log"
)

// loadAgendaMengajarMock memuat data mock dari file YAML docs/agenda-mengajar.yml
// dan hanya mengambil bagian tertentu berdasarkan key agar fleksibel
func loadAgendaMengajarMock(key string) any {
	data, err := LoadYAMLSection("docs/agenda-mengajar.yml", key)
	if err != nil {
		log.Printf("⚠️ Gagal memuat mock YAML untuk agenda mengajar [%s]: %v", key, err)
		return fallbackAgendaMengajarMock(key)
	}
	return data
}

// fallbackMock memberikan data cadangan ketika file YAML gagal dimuat
func fallbackAgendaMengajarMock(key string) map[string]any {
	return map[string]any{
		"message": fmt.Sprintf("Fallback mock for %s", key),
		"datas": []map[string]any{
			{
				"id":       "01JDXK1A2B3C4D5E6F7G8H9J0K",
				"key":      "theme",
				"value":    "dark",
				"semester": "1",
			},
		},
		"status": "success",
	}
}
