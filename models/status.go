package models

import "time"

type Status struct {
	ID        int64     `bson:"_id" json:"_id"`
	IDStatus  int64     `bson:"id_status" json:"id_status"`
	Status    string    `bson:"status" json:"status"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
