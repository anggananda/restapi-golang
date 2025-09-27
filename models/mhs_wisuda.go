package models

type MhsWisuda struct {
	ID          int    `bson:"_id" json:"_id"`
	NIM         string `bson:"nim" json:"nim"`
	NamaLengkap string `bson:"nama_lengkap" json:"nama_lengkap"`
	TahunWisuda int    `bson:"tahun_wisuda" json:"tahun_wisuda"`
	BulanWisuda int    `bson:"bulan_wisuda" json:"bulan_wisuda"`
	NamaBulan   string `bson:"nama_bulan" json:"nama_bulan"`
	Unit        Unit   `bson:"unit" json:"unit"`
}
