// Package models contains the models used for the database
package models

type ReviewData struct {
	Username   string `bson:"username"`
	Subject    string `bson:"subject"`
	Topic      string `bson:"topic"`
	SrsLevel   int32  `bson:"srs_level"`
	ReviewTime int64  `bson:"review_time"`
}
