// Package models contains the models shared amongst multiple go services
package models

type Kanji struct {
	Character string `bson:"character" json:"character"`
	WaniKaniId int `bson:"wanikani_id" json:"wanikani_id"`
	WaniKanilevel int `bson:"wanikani_level" json:"wanikani_level"`
	Meanings []string `bson:"meanings" json:"meanings"`
	Onyomi []string `bson:"onyomi" json:"onyomi"`
	Kunyomi []string `bson:"kunyomi" json:"kunyomi"`
	Nanori []string `bson:"nanori" json:"nanori"`
}