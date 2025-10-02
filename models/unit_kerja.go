package models

type UnitKerja struct {
	UkID          int64         `bson:"uk_id" json:"uk_id"`
	UkPid         int64         `bson:"uk_pid" json:"uk_pid"`
	UkKode        string        `bson:"uk_kode" json:"uk_kode"`
	UkNama        string        `bson:"uk_nama" json:"uk_nama"`
	UkNamaOtk2016 string        `bson:"uk_nama_otk_2016" json:"uk_nama_otk_2016"`
	UkName        string        `bson:"uk_name" json:"uk_name"`
	UkFormatSurat string        `bson:"uk_format_surat" json:"uk_format_surat"`
	IsActive      int64         `bson:"is_active" json:"is_active"`
	UkType        int64         `bson:"uk_type" json:"uk_type"`
	UkProgram     string        `bson:"uk_program" json:"uk_program"`
	UkGroup       int64         `bson:"uk_group" json:"uk_group"`
	IsVokasi      int64         `bson:"is_vokasi" json:"is_vokasi"`
	Gelar         string        `bson:"gelar" json:"gelar"`
	GelarNama     string        `bson:"gelar_nama" json:"gelar_nama"`
	Title         string        `bson:"title" json:"title"`
	TitleName     string        `bson:"title_name" json:"title_name"`
	GelarPosisi   string        `bson:"gelar_posisi" json:"gelar_posisi"`
	ProgramStudi  *ProgramStudi `bson:"program_studi,omitempty" json:"program_studi,omitempty"`
	Jurusan       *Jurusan      `bson:"jurusan,omitempty" json:"jurusan,omitempty"`
	Fakultas      *Fakultas     `bson:"fakultas,omitempty" json:"fakultas,omitempty"`
}

type ProgramStudi struct {
	PrdID    int64  `bson:"prd_id" json:"prd_id"`
	PrdPid   int64  `bson:"prd_pid" json:"prd_pid"`
	PrdKode  string `bson:"prd_kode" json:"prd_kode"`
	PrdNama  string `bson:"prd_nama" json:"prd_nama"`
	PrdGroup int64  `bson:"prd_group" json:"prd_group"`
}

type Jurusan struct {
	JrsID    int64  `bson:"jrs_id" json:"jrs_id"`
	JrsPid   int64  `bson:"jrs_pid" json:"jrs_pid"`
	JrsKode  string `bson:"jrs_kode" json:"jrs_kode"`
	JrsNama  string `bson:"jrs_nama" json:"jrs_nama"`
	JrsGroup int64  `bson:"jrs_group" json:"jrs_group"`
}

type Fakultas struct {
	FktID    int64  `bson:"fkt_id" json:"fkt_id"`
	FktPid   int64  `bson:"fkt_pid" json:"fkt_pid"`
	FktKode  string `bson:"fkt_kode" json:"fkt_kode"`
	FktNama  string `bson:"fkt_nama" json:"fkt_nama"`
	FktGroup int64  `bson:"fkt_group" json:"fkt_group"`
}
