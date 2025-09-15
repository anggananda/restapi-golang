package models

type Unit struct {
	UKKode   string `json:"uk_kode" bson:"uk_kode"`
	FktKode  string `json:"fkt_kode" bson:"fkt_kode"`
	JrsKose  string `json:"jrs_kode" bson:"jrs_kode"`
	PrdKode  string `json:"prd_kode" bson:"prd_kode"`
	Fakultas string `json:"fakultas" bson:"fakultas"`
	Jurusan  string `json:"jurusan" bson:"jurusan"`
	Prodi    string `json:"prodi" bson:"prodi"`
}
