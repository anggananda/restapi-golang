package models

type Beasiswa struct {
	ID            int     `bson:"_id" json:"_id"`
	NIM           string  `bson:"nim" json:"nim"`
	Nama          string  `bson:"nama" json:"nama"`
	JenisBeasiswa string  `bson:"jenis_beasiswa" json:"jenis_beasiswa"`
	IPK           float64 `bson:"ipk" json:"ipk"`
	Status        string  `bson:"status" json:"status"`
	Tahun         int     `bson:"tahun" json:"tahun"`
	Semester      string  `bson:"semester" json:"semester"`
	SemesterType  string  `bson:"semester_type" json:"semester_type"`
	Periode       string  `bson:"periode" json:"periode"`
	Unit          Unit    `bson:"unit" json:"unit"`
}
