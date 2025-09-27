package models

type RekapPMB struct {
	ID         int     `bson:"_id" json:"_id"`
	Tahun      int     `bson:"tahun" json:"tahun"`
	Kode       string  `bson:"kode" json:"kode"`
	NamaProdi  string  `bson:"nama_prodi" json:"nama_prodi"`
	SNBP       StatPMB `bson:"snbp" json:"snbp"`
	SNBT       StatPMB `bson:"snbt" json:"snbt"`
	SMBJM_CBT  StatPMB `bson:"smbjm_cbt" json:"smbjm_cbt"`
	SMBJM_Rpt  StatPMB `bson:"smbjm_raport" json:"smbjm_raport"`
	SMBJM_Tlt  StatPMB `bson:"smbjm_talent" json:"smbjm_talent"`
	SMBJM_UTBK StatPMB `bson:"smbjm_utbk" json:"smbjm_utbk"`
	Profesi    StatPMB `bson:"profesi" json:"profesi"`
	Internas   StatPMB `bson:"internasional" json:"internasional"`
	Pasca      StatPMB `bson:"pasca" json:"pasca"`
	AdikPapua  StatPMB `bson:"adikpapua" json:"adikpapua"`
	Jumlah     StatPMB `bson:"jumlah" json:"jumlah"`
	Unit       Unit    `bson:"unit" json:"unit"`
}

type StatPMB struct {
	Peminat string `bson:"peminat" json:"peminat"`
	Lulus   string `bson:"lulus" json:"lulus"`
	Daftar  string `bson:"daftar" json:"daftar"`
}
