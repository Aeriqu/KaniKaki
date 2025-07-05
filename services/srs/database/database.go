// Package database contains the Database object which abstrats database
// implementations
package database

import (
	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/srs/models"
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

// -- DATABASE GET -- //

// GetUserReviewData retrieves all existing review data for a user based on their identifier and subject.
func (db *Database) GetUserReviewData(identifier string, subject string) (*[]models.ReviewData, error)  {
	return db.provider.GetUserReviewData(identifier, subject)
}

// GetDueReviewData retrieves all review data for a user that is due based on their identifier and subject.
func (db *Database) GetDueReviewData(identifier string, subject string) (*[]models.ReviewData, error) {
	return db.provider.GetDueReviewData(identifier, subject)
}

// -- DATABASE UPDATE -- //

// UpdateSrs updates the SRS data for a user based on their identifier, subject, topic,
// level, and review time.
func (db *Database) UpdateSrsItem(identifier string, subject string, topic string, level int32, reviewTime int64) (*models.ReviewData, error) {
	return db.provider.UpdateSrsItem(identifier, subject, topic, level, reviewTime)
}