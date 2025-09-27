package models

type RealisasiUnit struct {
	ID               int    `bson:"_id" json:"_id"`
	Kode             string `bson:"id" json:"id"`
	KodeUnit         string `bson:"kode_unit" json:"kode_unit"`
	NamaUnit         string `bson:"nama_unit" json:"nama_unit"`
	PaguPNBP         string `bson:"pagu_pnbp" json:"pagu_pnbp"`
	PaguRM           string `bson:"pagu_rm" json:"pagu_rm"`
	PaguRMBOPTN      string `bson:"pagu_rm_boptn" json:"pagu_rm_boptn"`
	RealisasiPNBP    string `bson:"realisasi_pnbp" json:"realisasi_pnbp"`
	RealisasiRM      string `bson:"realisasi_rm" json:"realisasi_rm"`
	RealisasiRMBOPTN string `bson:"realisasi_rm_boptn" json:"realisasi_rm_boptn"`
	TahunAnggaran    string `bson:"tahun_anggaran" json:"tahun_anggaran"`
}
