package models

type KaryaAkhir struct {
	ID              int    `bson:"_id" json:"_id"`
	CurrentState    string `bson:"current_state" json:"current_state"`
	Judul           string `bson:"judul" json:"judul"`
	MainStage       string `bson:"main_stage" json:"main_stage"`
	NamaLengkap     string `bson:"nama_lengkap" json:"nama_lengkap"`
	NamaPA          string `bson:"nama_pa" json:"nama_pa"`
	NamaPembimbing1 string `bson:"nama_pembimbing_1" json:"nama_pembimbing_1"`
	NamaPembimbing2 string `bson:"nama_pembimbing_2" json:"nama_pembimbing_2"`
	NamaPembimbing3 string `bson:"nama_pembimbing_3,omitempty" json:"nama_pembimbing_3,omitempty"`
	NamaPenguji1    string `bson:"nama_penguji_1" json:"nama_penguji_1"`
	NamaPenguji2    string `bson:"nama_penguji_2" json:"nama_penguji_2"`
	NamaPenguji3    string `bson:"nama_penguji_3,omitempty" json:"nama_penguji_3,omitempty"`
	NamaPenguji4    string `bson:"nama_penguji_4,omitempty" json:"nama_penguji_4,omitempty"`
	NamaPenguji5    string `bson:"nama_penguji_5,omitempty" json:"nama_penguji_5,omitempty"`
	NamaPenguji6    string `bson:"nama_penguji_6,omitempty" json:"nama_penguji_6,omitempty"`
	NilaiAkhir      string `bson:"nilai_akhir" json:"nilai_akhir"`
	NIM             string `bson:"nim" json:"nim"`
	StatusJudul     string `bson:"status_judul" json:"status_judul"`
	TahunMasuk      int    `bson:"tahun_masuk" json:"tahun_masuk"`
	Unit            Unit   `bson:"unit" json:"unit"`
}
