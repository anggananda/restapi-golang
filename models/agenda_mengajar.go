package models

type AgendaMengajar struct {
	ID          int      `bson:"_id" json:"_id"`
	Dosen       []string `bson:"dosen" json:"dosen"`
	IdKelas     string   `bson:"id_kelas" json:"id_kelas"`
	IdPenawaran string   `bson:"id_penawaran" json:"id_penawaran"`
	JenisKelas  string   `bson:"jenis_kelas" json:"jenis_kelas"`
	Kode        string   `bson:"kode" json:"kode"`
	Kurikulum   string   `bson:"kurikulum" json:"kurikulum"`
	Matakuliah  string   `bson:"matakuliah" json:"matakuliah"`
	NipDosen    []string `bson:"nip_dosen" json:"nip_dosen"`
	Periode     string   `bson:"periode" json:"periode"`
	Pertemuan   string   `bson:"pertemuan" json:"pertemuan"`
	Semester    string   `bson:"semester" json:"semester"`
	Sumber      string   `bson:"sumber" json:"sumber"`
	Tahun       string   `bson:"tahun" json:"tahun"`
	Unit        Unit     `bson:"unit" json:"unit"`
}
