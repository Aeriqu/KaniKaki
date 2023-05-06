// Package models contains the models used for the database
package models

type User struct {
	Username           string  `bson:"username"`
	Password           string  `bson:"password"`
	Type               int     `bson:"type"`
	WanikaniLevelLimit int     `bson:"wanikani_level_limit"`
	AuthTokens         []Token `bson:"auth_tokens"`
}
