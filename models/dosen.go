package models

import "time"

type Dosen struct {
	UserType          string     `bson:"user_type" json:"user_type"`
	DosenID           int64     `bson:"dosen_id" json:"dosen_id"`
	DosenKey          string     `bson:"dosen_key" json:"dosen_key"`
	NIDN              string     `bson:"nidn" json:"nidn"`
	NIP               string     `bson:"nip" json:"nip"`
	NIK               string     `bson:"nik" json:"nik"`
	NoKarpeg          string     `bson:"no_karpeg" json:"no_karpeg"`
	NamaLengkap       string     `bson:"nama_lengkap" json:"nama_lengkap"`
	NamaTanpaGelar    string     `bson:"nama_tanpa_gelar" json:"nama_tanpa_gelar"`
	Email             string     `bson:"email" json:"email"`
	Alamat            string     `bson:"alamat" json:"alamat"`
	JenisKelamin      string     `bson:"jenis_kelamin" json:"jenis_kelamin"`
	NoHP              string     `bson:"no_hp" json:"no_hp"`
	TempatLahir       string     `bson:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir      *time.Time `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Photo             string     `bson:"photo" json:"photo"`
	JabatanFungsional string     `bson:"jabatan_fungsional" json:"jabatan_fungsional"`
	Grade             int64     `bson:"grade" json:"grade"`
	GradeFungsional   int64     `bson:"grade_fungsional" json:"grade_fungsional"`
	Golongan          string     `bson:"golongan" json:"golongan"`
	Pangkat           string     `bson:"pangkat" json:"pangkat"`
	Strata            string     `bson:"strata" json:"strata"`
	StatusSerdos      int64     `bson:"status_serdos" json:"status_serdos"`

	UnitKerjaTugas      UnitKerjaTugas    `bson:"unit_kerja_tugas" json:"unit_kerja_tugas"`
	UnitKerjaHomebase   UnitKerjaHomebase `bson:"unit_kerja_homebase" json:"unit_kerja_homebase"`
	UnitKerjaAkreditasi UnitKerjaHomebase `bson:"unit_kerja_akreditasi" json:"unit_kerja_akreditasi"`

	StatusPegawai     StatusPegawaiDosen `bson:"status_pegawai" json:"status_pegawai"`
	JabatanStruktural *JabatanStruktural `bson:"jabatan_struktural,omitempty" json:"jabatan_struktural,omitempty"`

	ProfilSingkat string `bson:"profil_singkat" json:"profil_singkat"`
	NUPTK         string `bson:"nuptk" json:"nuptk"`
	IDSDMSister   string `bson:"id_sdm_sister" json:"id_sdm_sister"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

type UnitKerjaTugas struct {
	IDJurusan   int64 `bson:"id_jurusan" json:"id_jurusan"`
	KodeProdi   string `bson:"kode_prodi" json:"kode_prodi"`
	KodeJurusan string `bson:"kode_jurusan" json:"kode_jurusan"`
}

type UnitKerjaHomebase struct {
	UkID        string `bson:"uk_id" json:"uk_id"`
	UkKode      string `bson:"uk_kode" json:"uk_kode"`
	UkNama      string `bson:"uk_nama" json:"uk_nama"`
	KodeProdi   string `bson:"kode_prodi" json:"kode_prodi"`
	KodeJurusan string `bson:"kode_jurusan" json:"kode_jurusan"`
}

type StatusPegawaiDosen struct {
	Status           string `bson:"status" json:"status"`
	StatusSekarang   string `bson:"status_sekarang" json:"status_sekarang"`
	TMTStatus        string `bson:"tmt_status" json:"tmt_status"`
	IDStatusSekarang int64 `bson:"id_status_sekarang" json:"id_status_sekarang"`
	IDStatusPegawai  int64 `bson:"id_status_pegawai" json:"id_status_pegawai"`
}

type JabatanStruktural struct {
	JbtID    int64 `bson:"jbt_id" json:"jbt_id"`
	Nama     string `bson:"nama" json:"nama"`
	IDUnit   int64 `bson:"id_unit" json:"id_unit"`
	KodeUnit string `bson:"kode_unit" json:"kode_unit"`
}
