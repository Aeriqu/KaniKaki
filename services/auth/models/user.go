// Package models contains the models used for the database
package models

type User struct {
	Username   string  `bson:"username"`
	Password   string  `bson:"password"`
	AuthTokens []Token `bson:"auth_tokens"`
}
