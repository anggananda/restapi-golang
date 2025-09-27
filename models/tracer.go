package models

import "time"

type Tracer struct {
	ID                       int       `bson:"_id" json:"_id"`
	IDMahasiswa              int       `bson:"id_mahasiswa" json:"id_mahasiswa"`
	NIMMahasiswa             string    `bson:"nim_mahasiswa" json:"nim_mahasiswa"`
	NamaMahasiswa            string    `bson:"nama_mahasiswa" json:"nama_mahasiswa"`
	EmailMahasiswa           string    `bson:"email_mahasiswa" json:"email_mahasiswa"`
	NoTelp                   string    `bson:"no_telp" json:"no_telp"`
	NIKMahasiswa             string    `bson:"nik_mahasiswa" json:"nik_mahasiswa"`
	NPWPMahasiswa            string    `bson:"npwp_mahasiswa" json:"npwp_mahasiswa"`
	TglLahirMahasiswa        time.Time `bson:"tgl_lahir_mahasiswa" json:"tgl_lahir_mahasiswa"`
	JenisKelaminMahasiswa    string    `bson:"jenis_kelamin_mahasiswa" json:"jenis_kelamin_mahasiswa"`
	BulanLulusMahasiswa      int       `bson:"bulan_lulus_mahasiswa" json:"bulan_lulus_mahasiswa"`
	TahunLulusMahasiswa      int       `bson:"tahun_lulus_mahasiswa" json:"tahun_lulus_mahasiswa"`
	IPKMahasiswa             float64   `bson:"ipk_mahasiswa" json:"ipk_mahasiswa"`
	TglLulusMahasiswa        time.Time `bson:"tgl_lulus_mahasiswa" json:"tgl_lulus_mahasiswa"`
	StatusMahasiswa          string    `bson:"status_mahasiswa" json:"status_mahasiswa"`
	IDJurusan                int       `bson:"id_jurusan" json:"id_jurusan"`
	UserID                   int       `bson:"user_id" json:"user_id"`
	DeletedAt                time.Time `bson:"deleted_at" json:"deleted_at"`
	StatusPengisian          string    `bson:"status_pengisian" json:"status_pengisian"`
	PersentasePengisian      int       `bson:"persentase_pengisian" json:"persentase_pengisian"`
	PengisianTerakhir        time.Time `bson:"pengisian_terakhir" json:"pengisian_terakhir"`
	Dikti                    string    `bson:"dikti" json:"dikti"`
	BulanWisuda              int       `bson:"bulan_wisuda" json:"bulan_wisuda"`
	TahunWisuda              int       `bson:"tahun_wisuda" json:"tahun_wisuda"`
	NoIjasah                 string    `bson:"no_ijasah" json:"no_ijasah"`
	NoSKYudisium             string    `bson:"no_sk_yudisium" json:"no_sk_yudisium"`
	Ayah                     string    `bson:"ayah" json:"ayah"`
	Ibu                      string    `bson:"ibu" json:"ibu"`
	Saudara                  string    `bson:"saudara" json:"saudara"`
	Wali                     string    `bson:"wali" json:"wali"`
	Jenjang                  string    `bson:"jenjang" json:"jenjang"`
	StatusSaatIni            string    `bson:"status_saat_ini" json:"status_saat_ini"`
	MasaTungguSebelumLulus   string    `bson:"masa_tunggu_sebelum_lulus" json:"masa_tunggu_sebelum_lulus"`
	MasaTungguSetelahLulus   string    `bson:"masa_tunggu_setelah_lulus" json:"masa_tunggu_setelah_lulus"`
	Provinsi                 string    `bson:"provinsi" json:"provinsi"`
	Kabupaten                string    `bson:"kabupaten" json:"kabupaten"`
	AlamatPerusahaan         string    `bson:"alamat_perusahaan" json:"alamat_perusahaan"`
	Gaji                     string    `bson:"gaji" json:"gaji"`
	JenisPerusahaan          string    `bson:"jenis_perusahaan" json:"jenis_perusahaan"`
	NamaPerusahaan           string    `bson:"nama_perusahaan" json:"nama_perusahaan"`
	JabatanDalamBerwirausaha string    `bson:"jabatan_dalam_berwirausaha" json:"jabatan_dalam_berwirausaha"`
	TingkatTempatKerja       string    `bson:"tingkat_tempat_kerja" json:"tingkat_tempat_kerja"`
	SumberBiayaStudiLanjut   string    `bson:"sumber_biaya_studi_lanjut" json:"sumber_biaya_studi_lanjut"`
	PerguruanTinggiLanjut    string    `bson:"perguruan_tinggi_studi_lanjut" json:"perguruan_tinggi_studi_lanjut"`
	ProdiMasukStudiLanjut    string    `bson:"prodi_masuk_studi_lanjut" json:"prodi_masuk_studi_lanjut"`
	TanggalMasukStudiLanjut  string    `bson:"tanggal_masuk_studi_lanjut" json:"tanggal_masuk_studi_lanjut"`
	Unit                     Unit      `bson:"unit" json:"unit"`
}
