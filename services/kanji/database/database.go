// Package database contains the Database object which abstrats database
// implementations
package database

import (
	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/common/models"
)

type Database struct {
	provider ProviderDatabase
}

// Init instantializes the database with a given database provider
// (ProviderDatabase).
func Init(provider ProviderDatabase) *Database {
	db := &Database{}
	db.provider = provider
	db.provider.Connect()
	logger.Info("database connection established")
	return db
}

// -- DATABASE ADD -- //

	// AddMultipleKanji adds a list of kanji to the database.
func (db *Database)	AddMultipleKanji(kanjiList []models.Kanji) (string, error) {
	return db.provider.AddMultipleKanji(kanjiList)
}

// -- DATABASE GET -- //

// GetKanjiByCharacter grabs a kanji with the a given identifier (character).
func (db *Database)	GetKanjiByCharacter(character string) (*models.Kanji, error) {
	return db.provider.GetKanjiByCharacter(character)
}

// GetKanjiByRange grabs a list a kanji by a given range
func (db *Database) GetKanjiByRange(lowerLimit int, upperLimit int) (*[]models.Kanji, error) {
	return db.provider.GetKanjiByRange(lowerLimit, upperLimit)
}
