package models

import "time"

type PegawaiHistory struct {
	ID                string                `json:"_id" bson:"_id"`
	Alamat            string                `json:"alamat" bson:"alamat"`
	CreatedAt         time.Time             `json:"created_at" bson:"created_at"`
	EmailAdmin        string                `json:"email_admin" bson:"email_admin"`
	EmailSSO          string                `json:"email_sso" bson:"email_sso"`
	ExpiryAt          *time.Time            `json:"expiry_at" bson:"expiry_at"`
	FirstGolongan     string                `json:"first_golongan" bson:"first_golongan"`
	FirstTglSK        string                `json:"first_tgl_sk" bson:"first_tgl_sk"`
	FirstTglTerimaSK  string                `json:"first_tgl_terima_sk" bson:"first_tgl_terima_sk"`
	FirstTMT          string                `json:"first_tmt" bson:"first_tmt"`
	History           []HistoryPegawaiEntry `json:"history" bson:"history"`
	HP                string                `json:"hp" bson:"hp"`
	PrimaryID         int                   `json:"id" bson:"id"`
	IDSDMSister       string                `json:"id_sdm_sister" bson:"id_sdm_sister"`
	JabatanFungsional JabFungsional         `json:"jabatan_fungsional" bson:"jabatan_fungsional"`
	JabatanStruktural JabStruktural         `json:"jabatan_struktural" bson:"jabatan_struktural"`
	JK                string                `json:"jk" bson:"jk"`
	LastGolongan      string                `json:"last_golongan" bson:"last_golongan"`
	LastPangkat       string                `json:"last_pangkat" bson:"last_pangkat"`
	LastStrata        string                `json:"last_strata" bson:"last_strata"`
	Nama              string                `json:"nama" bson:"nama"`
	NamaTanpaGelar    string                `json:"nama_tanpa_gelar" bson:"nama_tanpa_gelar"`
	NIK               string                `json:"nik" bson:"nik"`
	NIP               string                `json:"nip" bson:"nip"`
	NIPLama           string                `json:"nip_lama" bson:"nip_lama"`
	NoIndukUndiksha   string                `json:"no_induk_undiksha" bson:"no_induk_undiksha"`
	NoKarpeg          string                `json:"no_karpeg" bson:"no_karpeg"`
	PegawaiKey        int                   `json:"pegawai_key" bson:"pegawai_key"`
	Photo             interface{}           `json:"photo" bson:"photo"`
	Status            StatusPegawai         `json:"status" bson:"status"`
	StrukturalFungsi  string                `json:"struktural_fungsional" bson:"struktural_fungsional"`
	TempatLahir       string                `json:"tempat_lahir" bson:"tempat_lahir"`
	TglLahir          string                `json:"tgl_lahir" bson:"tgl_lahir"`
	TMTCPNS           string                `json:"tmt_cpns" bson:"tmt_cpns"`
	Unit              Unit                  `json:"unit" bson:"unit"`
	UpdatedAt         time.Time             `json:"updated_at" bson:"updated_at"`
}

type HistoryPegawaiEntry struct {
	Tahun           int    `json:"tahun" bson:"tahun"`
	StatusPegawai   string `json:"status_pegawai" bson:"status_pegawai"`
	StatusKeaktifan string `json:"status_keaktifan" bson:"status_keaktifan"`
	IDStatusPegawai int    `json:"id_status_pegawai" bson:"id_status_pegawai"`
	IDStatusAktif   int    `json:"id_status_keaktifan" bson:"id_status_keaktifan"`
}

type JabFungsional struct {
	Nama    string `json:"nama" bson:"nama"`
	ID      int    `json:"id" bson:"id"`
	GradeID int    `json:"grade_id" bson:"grade_id"`
	Unit    string `json:"unit" bson:"unit"`
	UnitID  int    `json:"unit_id" bson:"unit_id"`
}

type JabStruktural struct {
	Nama    string `json:"nama" bson:"nama"`
	ID      int    `json:"id" bson:"id"`
	GradeID int    `json:"grade_id" bson:"grade_id"`
	Unit    string `json:"unit" bson:"unit"`
	UnitID  int    `json:"unit_id" bson:"unit_id"`
}

type StatusPegawai struct {
	PKey              int    `json:"p_key" bson:"p_key"`
	StatusKey         int    `json:"status_key" bson:"status_key"`
	StatusKepKey      int    `json:"statuskep_key" bson:"statuskep_key"`
	StatusPegawaiNama string `json:"status_pegawai_nama" bson:"status_pegawai_nama"`
	StatusKeaktifan   string `json:"status_keaktifan" bson:"status_keaktifan"`
	TMTStatusPegawai  string `json:"tmt_status_pegawai" bson:"tmt_status_pegawai"`
}

type PegawaiHistoryResponse struct {
	NIP             string `json:"nip" bson:"nip"`
	NoIndukUndiksha string `bson:"no_induk_undiksha" json:"no_induk_undiksha"`
	Nama            string `json:"nama" bson:"nama"`
	Fakultas        string `json:"fakultas" bson:"fakultas"`
	Jurusan         string `json:"jurusan" bson:"jurusan"`
	Prodi           string `json:"prodi" bson:"prodi"`
	Tahun           int    `json:"tahun" bson:"tahun"`
	StatusPegawai   string `json:"status_pegawai" bson:"status_pegawai"`
	StatusKeaktifan string `json:"status_keaktifan" bson:"status_keaktifan"`
	LastStrata      string `bson:"last_strata" json:"last_strata"`
}
