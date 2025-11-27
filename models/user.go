package models

import (
	"time"
)

type UnitInfo struct {
	UkID    *int64  `json:"uk_id,omitempty" bson:"uk_id,omitempty"`
	UkKode  *string `json:"uk_kode,omitempty" bson:"uk_kode,omitempty"`
	UkNama  *string `json:"uk_nama,omitempty" bson:"uk_nama,omitempty"`
	UkGroup *int    `json:"uk_group,omitempty" bson:"uk_group,omitempty"`

	FktKode *string `json:"fkt_kode,omitempty" bson:"fkt_kode,omitempty"`
	JrsKode *string `json:"jrs_kode,omitempty" bson:"jrs_kode,omitempty"`
	PrdKode *string `json:"prd_kode,omitempty" bson:"prd_kode,omitempty"`

	Fakultas *string `json:"fakultas,omitempty" bson:"fakultas,omitempty"`
	Jurusan  *string `json:"jurusan,omitempty" bson:"jurusan,omitempty"`
	Prodi    *string `json:"prodi,omitempty" bson:"prodi,omitempty"`
}

type AuthInfo struct {
	IDLevel    int64     `json:"id_level" bson:"id_level"`
	EmailSSO   string    `json:"email_sso" bson:"email_sso"`
	Niu        string    `json:"niu" bson:"niu"`
	Jenis      string    `json:"jenis" bson:"jenis"`
	LevelAkun  int       `json:"level_akun" bson:"level_akun"`
	LevelKode  string    `json:"level_kode" bson:"level_kode"`
	CustomPass *string   `json:"custom_pass" bson:"custom_pass"`
	Edited     string    `json:"edited" bson:"edited"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}

type ProfileInfo struct {
	UserType string `json:"user_type" bson:"user_type"`
	Email    string `json:"email" bson:"email"`
	Niu      string `json:"niu" bson:"niu"`

	FullName *string `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Key      *string `json:"key,omitempty" bson:"key,omitempty"`
	Nidn     *string `json:"nidn,omitempty" bson:"nidn,omitempty"`
	Nip      *string `json:"nip,omitempty" bson:"nip,omitempty"`
}

type UserAuth struct {
	ID       string      `json:"_id" bson:"_id"`
	AuthInfo AuthInfo    `json:"auth_info" bson:"auth_info"`
	Profile  ProfileInfo `json:"profile" bson:"profile"`
	Unit     UnitInfo    `json:"unit" bson:"unit"`
}
