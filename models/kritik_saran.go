package models

type KritikSaran struct {
	ID       int      `bson:"_id" json:"_id"`
	NIP      string   `bson:"nip" json:"nip"`
	Nama     string   `bson:"nama" json:"nama"`
	Saran    []string `bson:"saran" json:"saran"`
	Tahun    string   `bson:"tahun" json:"tahun"`
	Semester string   `bson:"semester" json:"semester"`
	Periode  string   `bson:"periode" json:"periode"`
	Unit     Unit     `bson:"unit" json:"unit"`
}
