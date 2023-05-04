package database

import "github.com/Aeriqu/kanikaki/common/models"

// ProviderDatabase is an interface to enforce similar basic actions amongst
// all database providers.
type ProviderDatabase interface {
	// -- DATABASE META -- //

	// Connect starts the connection to the database and exits of the connection
	// is invalid.
	Connect()

	// -- DATABASE ADD -- //

	// AddMultipleKanji adds a list of kanji to the database.
	AddMultipleKanji(kanjiList []models.Kanji) (string, error)

	// -- DATABASE GET -- //

	// GetKanjiByCharacter grabs a kanji with the a given identifier (character).
	GetKanjiByCharacter(character string) (*models.Kanji, error)

	// GetKanjiByRange grabs a list a kanji by a given range
	GetKanjiByRange(lowerLimit int, upperLimit int) (*[]models.Kanji, error)
}
