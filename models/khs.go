package models

import "time"

type Khs struct {
	ID       int       `bson:"_id" json:"_id"`
	NIM      string    `bson:"nim" json:"nim"`
	NamaMHS  string    `bson:"nama_mhs" json:"nama_mhs"`
	Semester string    `bson:"semester" json:"semester"`
	Tahun    string    `bson:"tahun" json:"tahun"`
	Dilihat  time.Time `bson:"dilihat" json:"dilihat"`
	Foto     string    `bson:"foto" json:"foto"`
	Unit     Unit      `bson:"unit" json:"unit"`
}
