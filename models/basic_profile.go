package models

import "time"

type BasicProfile struct {
	UserType  string    `bson:"user_type" json:"user_type"`
	Email     string    `bson:"email" json:"email"`
	NIU       string    `bson:"niu" json:"niu"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
