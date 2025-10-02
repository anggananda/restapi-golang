package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type UserAuth struct {
	ID        string      `bson:"_id" json:"_id"`
	AuthInfo  AuthInfo    `bson:"auth_info" json:"auth_info"`
	UnitKerja UnitKerja   `bson:"unit_kerja" json:"unit_kerja"`
	Profile   UserProfile `bson:"profile" json:"profile"`
}

type AuthInfo struct {
	IDLevel    int64     `bson:"id_level" json:"id_level"`
	EmailSSO   string    `bson:"email_sso" json:"email_sso"`
	Niu        string    `bson:"niu" json:"niu"`
	Jenis      string    `bson:"jenis" json:"jenis"`
	LevelAkun  int64     `bson:"level_akun" json:"level_akun"`
	LevelKode  string    `bson:"level_kode" json:"level_kode"`
	CustomPass string    `bson:"custom_pass" json:"custom_pass"`
	Edited     string    `bson:"edited" json:"edited"`
	CreatedAt  time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt  time.Time `bson:"updated_at" json:"updated_at"`
}

type UserProfile struct {
	Profile interface{} `json:"profile"`
}

func (up *UserProfile) UnmarshalBSON(data []byte) error {
	var d Dosen
	if err := bson.Unmarshal(data, &d); err == nil && d.DosenID != 0 {
		up.Profile = d
		return nil
	}

	var p Pegawai
	if err := bson.Unmarshal(data, &p); err == nil && p.PegawaiID != 0 {
		up.Profile = p
		return nil
	}

	var b BasicProfile
	if err := bson.Unmarshal(data, &b); err == nil {
		up.Profile = b
		return nil
	}

	return fmt.Errorf("unknown profile format")
}
