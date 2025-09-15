package models

type Perpem struct {
	IdPenarawan string   `bson:"id_penawaran" json:"id_penawaran"`
	Kode        string   `bson:"kode" json:"kode"`
	IdKelas     string   `bson:"id_kelas" json:"id_kelas"`
	MK          string   `bson:"mk" json:"mk"`
	Kurikulum   string   `bson:"kurikulum" json:"kurikulum"`
	Pertemuan   string   `bson:"pertemuan" json:"pertemuan"`
	Dosen       []string `bson:"dosen" json:"dosen"`
	Metode      string   `bson:"metode" json:"metode"`
	Silabus     string   `bson:"silabus" json:"silabus"`
	Kontrak     string   `bson:"kontrak" json:"kontrak"`
	Rps         string   `bson:"rps" json:"rps"`
	Rtm         string   `bson:"rtm" json:"rtm"`
	Semester    string   `bson:"semester" json:"semester"`
	Tahun       string   `bson:"tahun" json:"tahun"`
	Unit        Unit     `bson:"unit" json:"unit"`
}
