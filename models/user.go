package models

type User struct {
	IDUser   int    `bson:"_id" json:"_id"`
	Username string `bson:"username" json:"username"`
}
