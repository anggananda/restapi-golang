package mocks

import "sync"

var once sync.Once

// MockRegistry menyimpan mapping endpoint dan method ke response mock
var MockRegistry map[string]map[string]any

// Initialize memuat semua mock dari file YAML atau variabel lokal
func Initialize() {
	once.Do(func() {
		MockRegistry = map[string]map[string]any{
			//agenda mengajar
			"/api/v1/agenda-mengajar": {
				"GET": loadAgendaMengajarMock("agenda_mengajar_filtered"),
			},

			//angket mahasiswa
			"/api/v1/angket-mhs": {
				"GET": loadAngketMhs("angket_mhs_filtered"),
			},
		}
	})
}

// GetMockResponse mencari mock berdasarkan path dan method
func GetMockResponse(path, method string) (any, bool) {
	if MockRegistry == nil {
		Initialize()
	}
	if methods, exists := MockRegistry[path]; exists {
		if resp, ok := methods[method]; ok {
			return resp, true
		}
	}
	return nil, false
}
