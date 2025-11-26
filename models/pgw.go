package models

import "time"

type Pegawai struct {
	UserType       string     `bson:"user_type" json:"user_type"`
	PegawaiID      int64      `bson:"pegawai_id" json:"pegawai_id"`
	PegawaiKey     int64      `bson:"pegawai_key" json:"pegawai_key"`
	NIP            string     `bson:"nip" json:"nip"`
	NIPLama        string     `bson:"nip_lama" json:"nip_lama"`
	NIK            string     `bson:"nik" json:"nik"`
	NoKarpeg       string     `bson:"no_karpeg" json:"no_karpeg"`
	NamaLengkap    string     `bson:"nama_lengkap" json:"nama_lengkap"`
	NamaTanpaGelar string     `bson:"nama_tanpa_gelar" json:"nama_tanpa_gelar"`
	Email          string     `bson:"email" json:"email"`
	Alamat         string     `bson:"alamat" json:"alamat"`
	JenisKelamin   string     `bson:"jenis_kelamin" json:"jenis_kelamin"`
	NoHP           string     `bson:"no_hp" json:"no_hp"`
	TempatLahir    string     `bson:"tempat_lahir" json:"tempat_lahir"`
	TanggalLahir   *time.Time `bson:"tanggal_lahir,omitempty" json:"tanggal_lahir,omitempty"`
	Photo          string     `bson:"photo" json:"photo"`
	TMTCPNS        *time.Time `bson:"tmt_cpns,omitempty" json:"tmt_cpns,omitempty"`

	UnitKerja        UnitKerjaPegawai `bson:"unit_kerja" json:"unit_kerja"`
	StrukturalFungsi string           `bson:"struktural_fungsional" json:"struktural_fungsional"`

	StatusPegawai     StatusPegawaiPegawai `bson:"status_pegawai" json:"status_pegawai"`
	JabatanFungsional *JabatanFungsional   `bson:"jabatan_fungsional,omitempty" json:"jabatan_fungsional,omitempty"`
	JabatanStruktural *JabatanFungsional   `bson:"jabatan_struktural,omitempty" json:"jabatan_struktural,omitempty"`

	Golongan Golongan `bson:"golongan" json:"golongan"`
	Strata   string   `bson:"strata" json:"strata"`

	IDSDMSister string     `bson:"id_sdm_sister" json:"id_sdm_sister"`
	ExpiryAt    *time.Time `bson:"expiry_at,omitempty" json:"expiry_at,omitempty"`
	CreatedAt   time.Time  `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at" json:"updated_at"`
}

type UnitKerjaPegawai struct {
	Kode string `bson:"kode" json:"kode"`
	ID   int64  `bson:"id" json:"id"`
	Nama string `bson:"nama" json:"nama"`
}

type StatusPegawaiPegawai struct {
	Status         string `bson:"status" json:"status"`
	StatusSekarang string `bson:"status_sekarang" json:"status_sekarang"`
	TMTStatus      string `bson:"tmt_status" json:"tmt_status"`
}

type JabatanFungsional struct {
	Nama    string `bson:"nama" json:"nama"`
	ID      int64  `bson:"id" json:"id"`
	GradeID int64  `bson:"grade_id" json:"grade_id"`
	Unit    string `bson:"unit" json:"unit"`
	UnitID  int64  `bson:"unit_id" json:"unit_id"`
}

type Golongan struct {
	Last          string     `bson:"last" json:"last"`
	Pangkat       string     `bson:"pangkat" json:"pangkat"`
	First         string     `bson:"first" json:"first"`
	FirstTMT      *time.Time `bson:"first_tmt,omitempty" json:"first_tmt,omitempty"`
	FirstTglSK    *time.Time `bson:"first_tgl_sk,omitempty" json:"first_tgl_sk,omitempty"`
	FirstTerimaSK *time.Time `bson:"first_tgl_terima_sk,omitempty" json:"first_tgl_terima_sk,omitempty"`
}
