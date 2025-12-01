package mocks

import "sync"

var once sync.Once

var MockRegistry map[string]map[string]any

func Initialize() {
	once.Do(func() {
		MockRegistry = map[string]map[string]any{

			"/api/v1/health-check": {
				"GET": loadMock("docs/public/health_check/health_check.yaml", "health"),
			},
			"/api/v1/user/details": {
				"GET": loadMock("docs/private/user/user.yaml", "user_profile"),
			},
			"/api/v1/agenda-mengajar": {
				"GET": loadMock("docs/private/agenda_mengajar/agenda_mengajar.yaml", "agenda_mengajar_filtered"),
			},

			"/api/v1/angket-mhs": {
				"GET": loadMock("docs/private/angket_mhs/angket_mhs.yaml", "angket_mhs_filtered"),
			},

			"/api/v1/beasiswa": {
				"GET": loadMock("docs/private/beasiswa/beasiswa.yaml", "beasiswa_filtered"),
			},

			"/api/v1/buku": {
				"GET": loadMock("docs/private/buku/buku.yaml", "buku_filtered"),
			},

			"/api/v1/hki": {
				"GET": loadMock("docs/private/hki/hki.yaml", "hki_filtered"),
			},

			"/api/v1/jurnal": {
				"GET": loadMock("docs/private/jurnal/jurnal.yaml", "jurnal_filtered"),
			},

			"/api/v1/karya-akhir": {
				"GET": loadMock("docs/private/karya_akhir/karya_akhir.yaml", "karya_akhir_filtered"),
			},

			"/api/v1/khs": {
				"GET": loadMock("docs/private/khs/khs.yaml", "khs_filtered"),
			},
			"/api/v1/kritik-saran": {
				"GET": loadMock("docs/private/kritik_saran/kritik_saran.yaml", "kritik_saran_filtered"),
			},
			"/api/v1/mhs-wisuda": {
				"GET": loadMock("docs/private/mhs_wisuda/mhs_wisuda.yaml", "mhs_wisuda_filtered"),
			},
			"/api/v1/penawaran": {
				"GET": loadMock("docs/private/penawaran/penawaran.yaml", "penawaran_filtered"),
			},
			"/api/v1/penelitian": {
				"GET": loadMock("docs/private/penelitian/penelitian.yaml", "penelitian_filtered"),
			},
			"/api/v1/pengabdian": {
				"GET": loadMock("docs/private/pengabdian/pengabdian.yaml", "pengabdian_filtered"),
			},
			"/api/v1/perpem": {
				"GET": loadMock("docs/private/perpem/perpem.yaml", "perpem_filtered"),
			},
			"/api/v1/prosiding": {
				"GET": loadMock("docs/private/prosiding/prosiding.yaml", "prosiding_filtered"),
			},
			"/api/v1/realisasi-bulan": {
				"GET": loadMock("docs/private/realisasi_bulan/realisasi_bulan.yaml", "realisasi_bulan_filtered"),
			},
			"/api/v1/realisasi-unit": {
				"GET": loadMock("docs/private/realisasi_unit/realisasi_unit.yaml", "realisasi_unit_filtered"),
			},
			"/api/v1/rekap-pmb": {
				"GET": loadMock("docs/private/rekap_pmb/rekap_pmb.yaml", "rekap_pmb_filtered"),
			},
			"/api/v1/tracer": {
				"GET": loadMock("docs/private/tracer/tracer.yaml", "tracer_filtered"),
			},
			"/api/v1/kerjasama": {
				"GET": loadMock("docs/private/kerjasama/kerjasama.yaml", "kerjasama_filtered"),
			},
			"/api/v1/mhs/history": {
				"GET": loadMock("docs/private/mahasiswa/mahasiswa_history.yaml", "mahasiswa_history_filtered"),
			},
			"/api/v1/mhs/:nim": {
				"GET": loadMock("docs/private/mahasiswa/detail_mahasiswa.yaml", "mahasiswa_filtered"),
			},
			"/api/v1/dosen/history": {
				"GET": loadMock("docs/private/dosen/dosen_history.yaml", "dosen_history_filtered"),
			},
			"/api/v1/dosen/:niu": {
				"GET": loadMock("docs/private/dosen/detail_dosen.yaml", "dosen_filtered"),
			},
			"/api/v1/pegawai/history": {
				"GET": loadMock("docs/private/pegawai/pegawai_history.yaml", "pegawai_history_filtered"),
			},
			"/api/v1/pegawai/:niu": {
				"GET": loadMock("docs/private/pegawai/detail_pegawai.yaml", "pegawai_filtered"),
			},
			"/api/v1/evaluasi-dosen": {
				"GET": loadMock("docs/private/evaluasi_dosen/evaluasi_dosen.yaml", "evaluasi_dosen_filtered"),
			},
			"/api/v1/unit-kerja": {
				"GET": loadMock("docs/private/unitkerja/unitkerja.yaml", "unit_kerja_filtered"),
			},
			"/api/v1/status-mhs": {
				"GET": loadMock("docs/private/status/status_mahasiswa.yaml", "status_filtered"),
			},
			"/api/v1/status-pegawai": {
				"GET": loadMock("docs/private/status/status_pegawai.yaml", "status_filtered"),
			},
			"/api/v1/status-keaktifan-pegawai": {
				"GET": loadMock("docs/private/status/status_keaktifan_pegawai.yaml", "status_filtered"),
			},

			"/api/v1/dashboard-mhs/overview": {
				"GET": loadMock("docs/private/dashboard_mahasiswa/overview_mahasiswa.yaml", "dashboard_card_filtered"),
			},
			"/api/v1/dashboard-mhs/fakultas": {
				"GET": loadMock("docs/private/dashboard_mahasiswa/drilldown_mahasiswa.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-mhs/jurusan": {
				"GET": loadMock("docs/private/dashboard_mahasiswa/drilldown_mahasiswa.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-mhs/prodi": {
				"GET": loadMock("docs/private/dashboard_mahasiswa/drilldown_mahasiswa.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-dosen/overview": {
				"GET": loadMock("docs/private/dashboard_dosen/overview_dosen.yaml", "dashboard_card_pegawai_filtered"),
			},
			"/api/v1/dashboard-dosen/fakultas": {
				"GET": loadMock("docs/private/dashboard_dosen/drilldown_dosen.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-dosen/jurusan": {
				"GET": loadMock("docs/private/dashboard_dosen/drilldown_dosen.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-dosen/prodi": {
				"GET": loadMock("docs/private/dashboard_dosen/drilldown_dosen.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-pegawai/overview": {
				"GET": loadMock("docs/private/dashboard_pegawai/overview_pegawai.yaml", "dashboard_card_pegawai_filtered"),
			},
			"/api/v1/dashboard-pegawai/fakultas": {
				"GET": loadMock("docs/private/dashboard_pegawai/drilldown_pegawai.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-pegawai/jurusan": {
				"GET": loadMock("docs/private/dashboard_pegawai/drilldown_pegawai.yaml", "drilldown_item_filtered"),
			},
			"/api/v1/dashboard-pegawai/prodi": {
				"GET": loadMock("docs/private/dashboard_pegawai/drilldown_pegawai.yaml", "drilldown_item_filtered"),
			},
		}
	})
}

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
