// Package database contains the Database object which abstrats database
// implementations
package database

import (
	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/auth/models"
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

// AddUser creates a new user in the database with a given identifier
// (username, email, etc) and credential.
func (db *Database) AddUser(identifier string, credential string) (result string, err error) {
	return db.provider.AddUser(identifier, credential)
}

// AddTokenToUser adds a valid token to a given user.
func (db *Database) AddTokenToUser(identifier string, newToken string, newTokenExpiration int64) (result string, err error) {
	return db.provider.AddTokenToUser(identifier, newToken, newTokenExpiration)
}

// -- DATABASE GET -- //

// GetUserByIdentifier gets all users with a given identifier
// (username/email/etc). (should be 1 user)
func (db *Database) GetUserByIdentifier(identifier string) (*models.User, error) {
	return db.provider.GetUserByIdentifier(identifier)
}

// GetUserByToken grabs all users that can be authenticated with a given
// token. (should be 1 user)
func (db *Database) GetUserByToken(token string) (*models.User, error) {
	return db.provider.GetUserByToken(token)
}

// -- DATABASE UPDATE -- //

// UpdateUserIdentifier updates a given user's identifier
// to a new provided identifier.
func (db *Database) UpdateUserIdentifier(identifier string, newIdentifier string) (result string, err error) {
	return db.provider.UpdateUserIdentifier(identifier, newIdentifier)
}

// UpdateUserCredential updates a given user's credential
func (db *Database) UpdateUserCredential(identifier string, newCredential string) (result string, err error) {
	return db.provider.UpdateUserCredential(identifier, newCredential)
}

// UpdateSpecificUserToken replaces an oldToken with a newToken on a given username
func (db *Database) UpdateSpecificUserToken(identifier string, oldToken string, newToken string, newTokenExpiration int64) (result *models.User, err error) {
	return db.provider.UpdateSpecificUserToken(identifier, oldToken, newToken, newTokenExpiration)
}

// -- DATABASE REMOVE FUNCTIONS -- //

// Removes expired tokens from a given user
func (db *Database) RemoveUserExpiredTokens(identifier string) (result *models.User, err error) {
	return db.provider.RemoveUserExpiredTokens(identifier)
}

// Removes a token from a user
func (db *Database) RemoveSpecificUserToken(identifier string, token string) (result *models.User, err error) {
	return db.provider.RemoveSpecificUserToken(identifier, token)
}