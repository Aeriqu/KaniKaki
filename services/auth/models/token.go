package models

type Token struct {
	Jwt        string `bson:"jwt"`
	Expiration int64  `bson:"expiration"`
}
