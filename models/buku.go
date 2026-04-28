package models

type Buku struct {
	ID                      int                 `bson:"_id" json:"_id"`
	JudulBuku               string              `bson:"judul_buku" json:"judul_buku"`
	Penerbit                string              `bson:"penerbit" json:"penerbit"`
	Tanggal                 string              `bson:"tanggal" json:"tanggal"`
	CreatedAt               string              `bson:"created_at" json:"created_at"`
	IsValid                 string              `bson:"isValid" json:"isValid"`
	Semester                string              `bson:"semester" json:"semester"`
	TahunAjaran             string              `bson:"tahun_ajaran" json:"tahun_ajaran"`
	TahunData               string              `bson:"tahun_data" json:"tahun_data"`
	Scope                   string              `bson:"scope" json:"scope"`
	TbName                  string              `bson:"tbName" json:"tbName"`
	PrimaryKey              string              `bson:"primaryKey" json:"primaryKey"`
	WaktuPelaksanaan        string              `bson:"waktu_pelaksanaan" json:"waktu_pelaksanaan"`
	NamaDosen               string              `bson:"nama_dosen" json:"nama_dosen"`
	KodeProdi               string              `bson:"kode_prodi" json:"kode_prodi"`
	NamaJurusan             string              `bson:"nama_jurusan" json:"nama_jurusan"`
	NamaFakultas            string              `bson:"nama_fakultas" json:"nama_fakultas"`
	Keterangan              string              `bson:"keterangan" json:"keterangan"`
	KategoriBuku            string              `bson:"kategori_buku" json:"kategori_buku"`
	ISBN                    string              `bson:"ISBN" json:"ISBN"`
	UrlDokumen              string              `bson:"url_dokumen" json:"url_dokumen"`
	UrlPerReview            string              `bson:"url_per_review" json:"url_per_review"`
	Satuan                  string              `bson:"satuan" json:"satuan"`
	JumlahHalaman           string              `bson:"jumlah_halaman" json:"jumlah_halaman"`
	VolumeKegiatan          string              `bson:"volume_kegiatan" json:"volume_kegiatan"`
	FileUpload              string              `bson:"file_upload" json:"file_upload"`
	Posisi                  string              `bson:"posisi" json:"posisi"`
	JmlNegaraPengedaran     string              `bson:"jml_negara_pengedaran" json:"jml_negara_pengedaran"`
	FilePendahuluan         string              `bson:"file_pendahuluan" json:"file_pendahuluan"`
	FileIsiBuku             string              `bson:"file_isi_buku" json:"file_isi_buku"`
	FilePenutupReferensi    string              `bson:"file_penutup_dan_referensi" json:"file_penutup_dan_referensi"`
	FilePersetujuanPenerbit string              `bson:"file_persetujuan_penerbit" json:"file_persetujuan_penerbit"`
	FileSelesaiDicetak      string              `bson:"file_selesai_dicetak" json:"file_selesai_dicetak"`
	JmlPenulis              string              `bson:"jml_penulis" json:"jml_penulis"`
	FileHasilUjiPlagiarim   string              `bson:"file_hasil_uji_plagiarim" json:"file_hasil_uji_plagiarim"`
	FilePenilaianReviewer   string              `bson:"file_penilaian_reviewer" json:"file_penilaian_reviewer"`
	UpdatedAt               string              `bson:"updated_at" json:"updated_at"`
	DeletedAt               string              `bson:"deleted_at" json:"deleted_at"`
	IsProduk                string              `bson:"is_produk" json:"is_produk"`
	ProdukPenelitianJudul   string              `bson:"produk_penelitian_judul" json:"produk_penelitian_judul"`
	ProdukPenelitianID      string              `bson:"produk_penelitian_id" json:"produk_penelitian_id"`
	ProdukPengabdianJudul   string              `bson:"produk_pengabdian_judul" json:"produk_pengabdian_judul"`
	ProdukPengabdianID      string              `bson:"produk_pengabdian_id" json:"produk_pengabdian_id"`
	Komentar                string              `bson:"komentar" json:"komentar"`
	ValidIpk                string              `bson:"valid_ipk" json:"valid_ipk"`
	ValidIpkKomentar        string              `bson:"valid_ipk_komentar" json:"valid_ipk_komentar"`
	CreateDosenID           string              `bson:"create_dosen_id" json:"create_dosen_id"`
	LevelCapaian            LevelCapaian        `bson:"level_capaian" json:"level_capaian"`
	SumberProduk            string              `bson:"sumber_produk" json:"sumber_produk"`
	ProdukPenelitian        ProdukInfo          `bson:"produk_penelitian" json:"produk_penelitian"`
	ProdukPengabdian        ProdukInfo          `bson:"produk_pengabdian" json:"produk_pengabdian"`
	MahasiswaPenelitian     []map[string]any    `bson:"mahasiswa_penelitian" json:"mahasiswa_penelitian"`
	AnggotaPenelitian       []AnggotaPenelitian `bson:"anggota_penelitian" json:"anggota_penelitian"`
	CronTahun               string              `bson:"cron_tahun" json:"cron_tahun"`
	CronSemester            string              `bson:"cron_semester" json:"cron_semester"`
	KodeKategoriBuku        string              `bson:"kode_kategori_buku" json:"kode_kategori_buku"`
	KodeScope               string              `bson:"kode_scope" json:"kode_scope"`
	Periode                 string              `bson:"periode" json:"periode"`
	SemesterType            string              `bson:"semester_type" json:"semester_type"`
	Unit                    Unit                `bson:"unit" json:"unit"`
}

type AnggotaPenelitian struct {
	Identitas Identitas `bson:"identitas" json:"identitas"`
	Peran     Peran     `bson:"peran" json:"peran"`
	UnitKerja UK        `bson:"unit_kerja" json:"unit_kerja"`
}

type Identitas struct {
	ID          int    `bson:"id" json:"id"`
	Nidn        string `bson:"nidn" json:"nidn"`
	Nip         string `bson:"nip" json:"nip"`
	NamaLengkap string `bson:"nama_lengkap" json:"nama_lengkap"`
	Email       string `bson:"email" json:"email"`
}

type Peran struct {
	PenulisKe        int    `bson:"penulis_ke" json:"penulis_ke"`
	IsKetua          bool   `bson:"is_ketua" json:"is_ketua"`
	StatusKonfirmasi string `bson:"status_konfirmasi" json:"status_konfirmasi"`
}

type UK struct {
	Institusi    string `bson:"institusi" json:"institusi"`
	Fakultas     string `bson:"fakultas" json:"fakultas"`
	KodeFakultas string `bson:"kode_fakultas" json:"kode_fakultas"`
	Jurusan      string `bson:"jurusan" json:"jurusan"`
	KodeJurusan  string `bson:"kode_jurusan" json:"kode_jurusan"`
}

type LevelCapaian struct {
	FilePendahuluan string `bson:"file_pendahuluan" json:"file_pendahuluan"`
	FileIsiBuku     string `bson:"file_isi_buku" json:"file_isi_buku"`
	FilePenutup     string `bson:"file_penutup" json:"file_penutup"`
	FilePersetujuan string `bson:"file_persetujuan" json:"file_persetujuan"`
	FileCetak       string `bson:"file_cetak" json:"file_cetak"`
}

type ProdukInfo struct {
	ID    any    `bson:"id" json:"id"`
	Judul string `bson:"judul" json:"judul"`
}
