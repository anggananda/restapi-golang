package models

// type Mahasiswa struct {
// 	ID                     string    `json:"id" bson:"_id"`
// 	SourceKey              string    `json:"source_key" bson:"source_key"`
// 	NIM                    string    `json:"nim" bson:"nim"`
// 	NISN                   string    `json:"nisn" bson:"nisn"`
// 	Nama                   string    `json:"nama" bson:"nama"`
// 	JK                     string    `json:"jk" bson:"jk"`
// 	TmpLahir               string    `json:"tmp_lahir" bson:"tmp_lahir"`
// 	TglLahir               string    `json:"tgl_lahir" bson:"tgl_lahir"`
// 	Alamat                 string    `json:"alamat" bson:"alamat"`
// 	RT                     string    `json:"rt" bson:"rt"`
// 	RW                     string    `json:"rw" bson:"rw"`
// 	KodePos                string    `json:"kode_pos" bson:"kode_pos"`
// 	Kelurahan              string    `json:"kelurahan" bson:"kelurahan"`
// 	HP                     string    `json:"hp" bson:"hp"`
// 	Telp                   string    `json:"telp" bson:"telp"`
// 	WA                     string    `json:"wa" bson:"wa"`
// 	Email2                 string    `json:"email2" bson:"email2"`
// 	Email                  string    `json:"email" bson:"email"`
// 	AgamaKey               int       `json:"agama_key" bson:"agama_key"`
// 	KodeProvinsi           string    `json:"kode_provinsi" bson:"kode_provinsi"`
// 	KodeKabupaten          string    `json:"kode_kabupaten" bson:"kode_kabupaten"`
// 	KodeKecamatan          string    `json:"kode_kecamatan" bson:"kode_kecamatan"`
// 	NPSNSekolah            string    `json:"npsn_sekolah" bson:"npsn_sekolah"`
// 	JurusanKey             int       `json:"jurusan_key" bson:"jurusan_key"`
// 	LastStatus             string    `json:"last_status" bson:"last_status"`
// 	NIK                    string    `json:"nik" bson:"nik"`
// 	DosenPAKey             string    `json:"dosen_pa_key" bson:"dosen_pa_key"`
// 	SPP                    string    `json:"spp" bson:"spp"`
// 	Kurikulum              string    `json:"kurikulum" bson:"kurikulum"`
// 	TempatKuliah           string    `json:"tempatkuliah" bson:"tempatkuliah"`
// 	Foto                   string    `json:"foto" bson:"foto"`
// 	NamaGadisIbuKandung    string    `json:"nama_gadis_ibu_kandung" bson:"nama_gadis_ibu_kandung"`
// 	Kewarganegaraan        string    `json:"kewarganegaraan" bson:"kewarganegaraan"`
// 	NoKPS                  string    `json:"no_kps" bson:"no_kps"`
// 	PenerimaBidikmisi      string    `json:"penerima_bidikmisi" bson:"penerima_bidikmisi"`
// 	TahunMasuk             string    `json:"tahun_masuk" bson:"tahun_masuk"`
// 	Permalink              string    `json:"permalink" bson:"permalink"`
// 	Validasi               string    `json:"validasi" bson:"validasi"`
// 	Keterangan             string    `json:"keterangan" bson:"keterangan"`
// 	JenisJalurMahasiswaKey int       `json:"jenis_jalur_mahasiswa_key" bson:"jenis_jalur_mahasiswa_key"`
// 	CreatedAt              time.Time `json:"created_at" bson:"created_at"`
// 	UpdatedAt              time.Time `json:"updated_at" bson:"updated_at"`
// 	ExpiryAt               time.Time `json:"expiry_at" bson:"expiry_at"`
// 	MetadataKey            int       `json:"metadata_key" bson:"metadata_key"`
// 	EmailSSO               string    `json:"email_sso" bson:"email_sso"`
// 	UKKode                 string    `json:"uk_kode" bson:"uk_kode"`
// 	UKID                   string    `json:"uk_id" bson:"uk_id"`
// 	UKProgram              string    `json:"uk_program" bson:"uk_program"`
// 	SemesterPosisi         int       `json:"semester_posisi" bson:"semester_posisi"`
// 	IDMahasiswa            string    `json:"id_mahasiswa" bson:"id_mahasiswa"`
// 	Status                 string    `json:"status" bson:"status"`
// 	IDStatusMahasiswa      string    `json:"id_status_mahasiswa" bson:"id_status_mahasiswa"`
// 	Sumber                 string    `json:"sumber" bson:"sumber"`
// 	TglMulaiKuliah         time.Time `json:"tgl_mulai_kuliah" bson:"tgl_mulai_kuliah"`
// }

// type MahasiswaHistoryResponse struct {
// 	NIM         string `json:"nim" bson:"nim"`
// 	Nama        string `json:"nama" bson:"nama"`
// 	Fakultas    string `json:"fakultas" bson:"fakultas"`
// 	Jurusan     string `json:"jurusan" bson:"jurusan"`
// 	Prodi       string `json:"prodi" bson:"prodi"`
// 	Angkatan    string `json:"angkatan" bson:"tahun_masuk"`
// 	Status      string `json:"status" bson:"status"`
// 	WargaNegara string `json:"warga_negara" bson:"kewarganegaraan"`
// 	NamaPA      string `json:"nama_pa" bson:"nama_pa"`
// }

type MahasiswaHistoryResponse struct {
	NIM             string `json:"nim" bson:"nim"`
	Nama            string `json:"nama" bson:"nama"`
	TahunMasuk      string `json:"tahun_masuk" bson:"tahun_masuk"`
	Kewarganegaraan string `json:"kewarganegaraan" bson:"kewarganegaraan"`
	Fakultas        string `json:"fakultas" bson:"fakultas"`                 // Field flat, bukan unit.fakultas
	Jurusan         string `json:"jurusan" bson:"jurusan"`                   // Field flat
	Prodi           string `json:"prodi" bson:"prodi"`                       // Field flat
	Tahun           int    `json:"tahun" bson:"tahun"`                       // Field flat, bukan history.tahun
	Semester        int    `json:"semester" bson:"semester"`                 // Field flat
	Status          string `json:"status" bson:"status"`                     // Field flat
	StatusSingkat   string `json:"status_singkatan" bson:"status_singkatan"` // Field flat
	NamaPA          string `json:"nama_pa" bson:"nama_pa"`                   // Field flat
}

type HistoryEntry struct {
	Tahun         int    `json:"tahun" bson:"tahun"`
	Semester      int    `json:"semester" bson:"semester"`
	Status        string `json:"status" bson:"status"`
	StatusSingkat string `json:"status_singkatan" bson:"status_singkatan"`
	NamaPA        string `json:"nama_pa" bson:"nama_pa"`
}

// Struct untuk unit


// Struct untuk mahasiswa
type Mahasiswa struct {
	NIM             string         `json:"nim" bson:"nim"`
	Nama            string         `json:"nama" bson:"nama"`
	TahunMasuk      string         `json:"tahun_masuk" bson:"tahun_masuk"`
	Kewarganegaraan string         `json:"kewarganegaraan" bson:"kewarganegaraan"`
	Unit            Unit           `json:"unit" bson:"unit"`
	History         []HistoryEntry `json:"history" bson:"history"`
}

type MahasiswaHistoryRequest struct {
	Nama            string `form:"nama"`
	Tahun           int    `form:"tahun"`
	Semester        string    `form:"semester"`
	Fakultas        string `form:"fakultas"`
	Jurusan         string `form:"jurusan"`
	Prodi           string `form:"prodi"`
	Status          string `form:"status"` // status_singkatan: A, C, D
	Kewarganegaraan string `form:"kewarganegaraan"`
	NIM             string `form:"nim"`
	Page            int    `form:"page" default:"1"`
	Limit           int    `form:"limit" default:"10"`
}
