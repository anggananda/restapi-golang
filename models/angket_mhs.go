package models

type AngketMhs struct {
	ID          int      `bson:"_id" json:"_id"`
	Dosen       []string `bson:"dosen" json:"dosen"`
	IdKelas     string   `bson:"id_kelas" json:"id_kelas"`
	IdPenawaran string   `bson:"id_penawaran" json:"id_penawaran"`
	Kode        string   `bson:"kode" json:"kode"`
	Mk          string   `bson:"mk" json:"mk"`
	NipDosen    []string `bson:"nip_dosen" json:"nip_dosen"`
	Periode     string   `bson:"periode" json:"periode"`
	Semester    string   `bson:"semester" json:"semester"`
	Tahun       string   `bson:"tahun" json:"tahun"`
	Unit        Unit     `bson:"unit" json:"unit"`
}
