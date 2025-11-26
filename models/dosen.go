package models

import "time"

type Dosen struct {
	ID              string  `bson:"_id" json:"_id"`
	RecordID        int64   `bson:"id" json:"id"`
	DosenID         int64   `bson:"dosen_id" json:"dosen_id"`
	DosenKey        string  `bson:"dosen_key" json:"dosen_key"`
	NoIndukUndiksha string  `bson:"no_induk_undiksha" json:"no_induk_undiksha"`
	NIDN            string  `bson:"nidn" json:"nidn"`
	NIP             string  `bson:"nip" json:"nip"`
	NIK             string  `bson:"nik" json:"nik"`
	NUPTK           *string `bson:"nuptk" json:"nuptk"`
	IDSDMSister     string  `bson:"id_sdm_sister" json:"id_sdm_sister"`
	NoKarpeg        string  `bson:"no_karpeg" json:"no_karpeg"`
	StatusSerdos    int64   `bson:"status_serdos" json:"status_serdos"`

	Nama           string  `bson:"nama" json:"nama"`
	NamaTanpaGelar string  `bson:"nama_tanpa_gelar" json:"nama_tanpa_gelar"`
	Email          string  `bson:"email" json:"email"`
	EmailSSO       string  `bson:"email_sso" json:"email_sso"`
	Alamat         string  `bson:"alamat" json:"alamat"`
	JK             string  `bson:"jk" json:"jk"`
	HP             string  `bson:"hp" json:"hp"`
	TempatLahir    string  `bson:"tempat_lahir" json:"tempat_lahir"`
	TglLahir       *string `bson:"tgl_lahir" json:"tgl_lahir"`
	Photo          *string `bson:"photo" json:"photo"`
	ProfilSingkat  string  `bson:"profil_singkat" json:"profil_singkat"`

	IDStatusPegawai  int64  `bson:"id_status_pegawai" json:"id_status_pegawai"`
	StatusPegawai    string `bson:"status_pegawai" json:"status_pegawai"`
	IDStatusSekarang int64  `bson:"id_status_sekarang" json:"id_status_sekarang"`

	StatusKeaktifan  string `bson:"status_keaktifan" json:"status_keaktifan"`
	TMTStatusPegawai string `bson:"tmt_status_pegawai" json:"tmt_status_pegawai"`

	JabatanFungsional string `bson:"jabatan_fungsional" json:"jabatan_fungsional"`
	Pangkat           string `bson:"pangkat" json:"pangkat"`
	Golongan          string `bson:"golongan" json:"golongan"`
	Strata            string `bson:"strata" json:"strata"`
	Grade             int64  `bson:"grade" json:"grade"`
	GradeFungsional   int64  `bson:"grade_fungsional" json:"grade_fungsional"`

	GolonganTMT *string `bson:"golongan_tmt" json:"golongan_tmt"`

	FirstGolongan    string  `bson:"first_golongan" json:"first_golongan"`
	FirstTMT         *string `bson:"first_tmt" json:"first_tmt"`
	FirstTglSK       *string `bson:"first_tgl_sk" json:"first_tgl_sk"`
	FirstTglTerimaSK *string `bson:"first_tgl_terima_sk" json:"first_tgl_terima_sk"`

	FungsionalTMT            *string `bson:"fungsional_tmt" json:"fungsional_tmt"`
	FungsionalTglSK          *string `bson:"fungsional_tgl_sk" json:"fungsional_tgl_sk"`
	FungsionalTglTerimaSK    *string `bson:"fungsional_tgl_terima_sk" json:"fungsional_tgl_terima_sk"`
	FungsionalTMTPAK         *string `bson:"fungsional_tmt_pak" json:"fungsional_tmt_pak"`
	FungsionalTglSKPAK       *string `bson:"fungsional_tgl_sk_pak" json:"fungsional_tgl_sk_pak"`
	FungsionalTglTerimaSKPAK *string `bson:"fungsional_tgl_terima_sk_pak" json:"fungsional_tgl_terima_sk_pak"`

	UnitTugas         UnitKerjaTugas     `bson:"unit_tugas" json:"unit_tugas"`
	Homebase          UnitKerjaHomebase  `bson:"homebase" json:"homebase"`
	Akreditasi        UnitKerjaHomebase  `bson:"akreditasi" json:"akreditasi"`
	JabatanStruktural *JabatanStruktural `bson:"jabatan_struktural" json:"jabatan_struktural"`
	Unit              UnitLengkap        `bson:"unit" json:"unit"`

	History   []HistoryDosen `bson:"history" json:"history"`
	CreatedAt time.Time      `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time      `bson:"updated_at" json:"updated_at"`
}

type UnitKerjaTugas struct {
	IDJurusan   int64  `bson:"id_jurusan" json:"id_jurusan"`
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

type JabatanStruktural struct {
	Nama       string  `bson:"nama" json:"nama"`
	JbtID      int64   `bson:"jbt_id" json:"jbt_id"`
	IDUnit     int64   `bson:"id_unit" json:"id_unit"`
	KodeUnit   string  `bson:"kode_unit" json:"kode_unit"`
	TglSelesai *string `bson:"tgl_selesai" json:"tgl_selesai"`
}

type UnitLengkap struct {
	UkKode   string `bson:"uk_kode" json:"uk_kode"`
	FktKode  string `bson:"fkt_kode" json:"fkt_kode"`
	JrsKode  string `bson:"jrs_kode" json:"jrs_kode"`
	PrdKode  string `bson:"prd_kode" json:"prd_kode"`
	Fakultas string `bson:"fakultas" json:"fakultas"`
	Jurusan  string `bson:"jurusan" json:"jurusan"`
	Prodi    string `bson:"prodi" json:"prodi"`
}

type HistoryDosen struct {
	Tahun             int    `bson:"tahun" json:"tahun"`
	StatusPegawai     string `bson:"status_pegawai" json:"status_pegawai"`
	StatusKeaktifan   string `bson:"status_keaktifan" json:"status_keaktifan"`
	IDStatusPegawai   int64  `bson:"id_status_pegawai" json:"id_status_pegawai"`
	IDStatusKeaktifan int64  `bson:"id_status_keaktifan" json:"id_status_keaktifan"`
}

type DosenHistoryResponse struct {
	NIP               string `json:"nip" bson:"nip"`
	NoIndukUndiksha   string `bson:"no_induk_undiksha" json:"no_induk_undiksha"`
	Nama              string `json:"nama" bson:"nama"`
	Fakultas          string `json:"fakultas" bson:"fakultas"`
	Jurusan           string `json:"jurusan" bson:"jurusan"`
	Prodi             string `json:"prodi" bson:"prodi"`
	Tahun             int    `json:"tahun" bson:"tahun"`
	StatusPegawai     string `json:"status_pegawai" bson:"status_pegawai"`
	StatusKeaktifan   string `json:"status_keaktifan" bson:"status_keaktifan"`
	JabatanFungsional string `json:"jabatan_fungsional" bson:"jabatan_fungsional"`
	Strata            string `bson:"strata" json:"strata"`
}
