package models

type Penawaran struct {
	ID             int    `bson:"_id" json:"_id"`
	JmlMhsAmbil    string `bson:"jml_mhs_ambil" json:"jml_mhs_ambil"`
	KodeMatakuliah string `bson:"kode_matakuliah" json:"kode_matakuliah"`
	Kurikulum      string `bson:"kurikulum" json:"kurikulum"`
	NamaKelas      string `bson:"nama_kelas" json:"nama_kelas"`
	NamaMatakuliah string `bson:"nama_matakuliah" json:"nama_matakuliah"`
	NamaPengampu   string `bson:"nama_pengampu" json:"nama_pengampu"`
	NipPengampu    string `bson:"nip_pengampu" json:"nip_pengampu"`
	Pengampu       string `bson:"pengampu" json:"pengampu"`
	Semester       string `bson:"semester" json:"semester"`
	Tahun          string `bson:"tahun" json:"tahun"`
	Unit           Unit   `bson:"unit" json:"unit"`
}
