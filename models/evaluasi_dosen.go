package models

import "time"

type Penilaian struct {
	SangatBaik int `bson:"sangat_baik" json:"sangat_baik"`
	Baik       int `bson:"baik" json:"baik"`
	Cukup      int `bson:"cukup" json:"cukup"`
	Kurang     int `bson:"kurang" json:"kurang"`
}

type EvaluasiDosen struct {
	ID              int64  `bson:"_id" json:"_id"`
	NIP             string `bson:"nip" json:"nip"`
	NoIndukUndiksha string `bson:"no_induk_undiksha" json:"no_induk_undiksha"`
	NamaLengkap     string `bson:"nama_lengkap" json:"nama_lengkap"`
	NamaKelas       string `bson:"nama_kelas" json:"nama_kelas"`
	KodeMatakuliah  string `bson:"kode_matakuliah" json:"kode_matakuliah"`
	NamaMatakuliah  string `bson:"nama_matakuliah" json:"nama_matakuliah"`

	Tahun        string `bson:"tahun" json:"tahun"`
	Semester     string `bson:"semester" json:"semester"`
	KodeProdi    string `bson:"kode_prodi" json:"kode_prodi"`
	KodeFakultas string `bson:"kode_fakultas" json:"kode_fakultas"`
	Unit         Unit   `bson:"unit" json:"unit"`

	CreatedAt time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time  `bson:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `bson:"deleted_at" json:"deleted_at,omitempty"`

	PerencanaanPerkuliahan                     Penilaian `bson:"Perencanaan_perkuliahan" json:"Perencanaan_perkuliahan"`
	RelevansiMateriDenganTujuanPembelajaran    Penilaian `bson:"Relevansi_materi_dengan_tujuan_pembelajaran" json:"Relevansi_materi_dengan_tujuan_pembelajaran"`
	PenguasaanMateriPerkuliahan                Penilaian `bson:"Penguasaan_materi_perkuliahan" json:"Penguasaan_materi_perkuliahan"`
	MetodeDanPendekatanPerkuliahan             Penilaian `bson:"Metode_dan_pendekatan_perkuliahan" json:"Metode_dan_pendekatan_perkuliahan"`
	InovasiDalamPerkuliahan                    Penilaian `bson:"Inovasi_dalam_perkuliahan" json:"Inovasi_dalam_perkuliahan"`
	KreatifitasDalamPerkuliahan                Penilaian `bson:"Kreatifitas_dalam_perkuliahan" json:"Kreatifitas_dalam_perkuliahan"`
	MediaPembelajaran                          Penilaian `bson:"Media_pembelajaran" json:"Media_pembelajaran"`
	SumberBelajar                              Penilaian `bson:"Sumber_belajar" json:"Sumber_belajar"`
	PenilaianHasilBelajar                      Penilaian `bson:"Penilaian_hasil_belajar" json:"Penilaian_hasil_belajar"`
	PenilaianProsesBelajar                     Penilaian `bson:"Penilaian_proses_belajar" json:"Penilaian_proses_belajar"`
	PemberianTugasPerkuliahan                  Penilaian `bson:"Pemberian_tugas_perkuliahan" json:"Pemberian_tugas_perkuliahan"`
	PengelolaanKelas                           Penilaian `bson:"Pengelolaan_kelas" json:"Pengelolaan_kelas"`
	MotivasiDanAntusiasmeMengajar              Penilaian `bson:"Motivasi_dan_antusiasme_mengajar" json:"Motivasi_dan_antusiasme_mengajar"`
	PenciptaanIklimBelajar                     Penilaian `bson:"Penciptaan_iklim_belajar" json:"Penciptaan_iklim_belajar"`
	Kedisiplinan                               Penilaian `bson:"Kedisiplinan" json:"Kedisiplinan"`
	PenegakanAturanPerkuliahan                 Penilaian `bson:"Penegakan_aturan_perkuliahan" json:"Penegakan_aturan_perkuliahan"`
	PengembanganKarakterMahasiswa              Penilaian `bson:"Pengembangan_karakter_mahasiswa" json:"Pengembangan_karakter_mahasiswa"`
	KeteladananDalamBersikapDanBertindak       Penilaian `bson:"Keteladanan_dalam_bersikap_dan_bertindak" json:"Keteladanan_dalam_bersikap_dan_bertindak"`
	KemampuanBerkomunikasi                     Penilaian `bson:"Kemampuan_berkomunikasi" json:"Kemampuan_berkomunikasi"`
	PenggunaanBahasaLisanDanTulisan            Penilaian `bson:"Penggunaan_bahasa_lisan_dan_tulisan" json:"Penggunaan_bahasa_lisan_dan_tulisan"`
	KemampuanBerinteraksiSosialDenganMahasiswa Penilaian `bson:"Kemampuan_berinteraksi_sosial_dengan_mahasiswa" json:"Kemampuan_berinteraksi_sosial_dengan_mahasiswa"`
}
