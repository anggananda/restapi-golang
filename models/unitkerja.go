package models

import "time"

type UnitKerjaMapping struct {
	ID          string    `bson:"_id" json:"_id"`
	Data        []Data    `bson:"data" json:"data"`
	LastUpdated time.Time `bson:"last_updated" json:"last_updated"`
}

type Data struct {
	UKID     int    `bson:"uk_id" json:"uk_id"`
	UKKode   string `bson:"uk_kode" json:"uk_kode"`
	UKNama   string `bson:"uk_nama" json:"uk_nama"`
	Children []Data `bson:"children" json:"children"`
}
