package models

type RealisasiBulan struct {
	ID                    int    `bson:"_id" json:"_id"`
	Kode                  string `bson:"id" json:"id"`
	Bulan                 string `bson:"bulan" json:"bulan"`
	RealisasiTotalPNBP    string `bson:"realisasi_total_pnbp" json:"realisasi_total_pnbp"`
	RealisasiTotalRM      string `bson:"realisasi_total_rm" json:"realisasi_total_rm"`
	RealisasiTotalRMBOPTN string `bson:"realisasi_total_rm_boptn" json:"realisasi_total_rm_boptn"`
	TahunAnggaran         string `bson:"tahun_anggaran" json:"tahun_anggaran"`
}
