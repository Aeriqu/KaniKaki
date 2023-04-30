package database

import "github.com/Aeriqu/kanikaki/services/auth/models"

// ProviderDatabase is an interface to enforce similar basic actions amongst
// all database providers.
type ProviderDatabase interface {
	// -- DATABASE META -- //

	// Connect starts the connection to the database and exits of the connection
	// is invalid.
	Connect()

	// -- DATABASE ADD -- //

	// AddUser creates a new user in the database with a given identifier
	// (username, email, etc) and credential.
	AddUser(identifier string, credential string) (result string, err error)

	// AddTokenToUser adds a valid token to a given user.
	AddTokenToUser(identifier string, newToken string, newTokenExpiration int64) (result string, err error)

	// -- DATABASE GET -- //

	// GetUserByIdentifier gets all users with a given identifier
	// (username/email/etc). (should be 1 user)
	GetUserByIdentifier(identifier string) (*models.User, error)

	// GetUserByToken grabs all users that can be authenticated with a given
	// token. (should be 1 user)
	GetUserByToken(token string) (*models.User, error)

	// -- DATABASE UPDATE -- //

	// UpdateUserIdentifier updates a given user's identifier
	// to a new provided identifier.
	UpdateUserIdentifier(identifier string, newIdentifier string) (result string, err error)

	// UpdateUserCredential updates a given user's credential
	UpdateUserCredential(identifier string, newCredential string) (result string, err error)

	// UpdateSpecificUserToken replaces an oldToken with a newToken on a given username
	UpdateSpecificUserToken(identifier string, oldToken string, newToken string, newTokenExpiration int64) (result *models.User, err error)

	// -- DATABASE REMOVE FUNCTIONS -- //

	// Removes expired tokens from a given user
	RemoveUserExpiredTokens(username string) (result *models.User, err error)

	// Removes a tokens from a given user
	RemoveSpecificUserToken(username string, token string) (result *models.User, err error)
}
