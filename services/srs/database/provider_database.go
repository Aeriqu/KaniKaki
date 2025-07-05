package database

import "github.com/Aeriqu/kanikaki/services/srs/models"

// ProviderDatabase is an interface to enforce similar basic actions amongst
// all database providers.
type ProviderDatabase interface {
	// -- DATABASE META -- //

	// Connect starts the connection to the database and exits of the connection
	// is invalid.
	Connect()

	// -- DATABASE ADD -- //

	// -- DATABASE GET -- //
	GetUserReviewData(identifier string, subject string) (*[]models.ReviewData, error)
	GetDueReviewData(identifier string, subject string) (*[]models.ReviewData, error)

	// -- DATABASE UPDATE -- //
	UpdateSrsItem(identifier string, subject string, topic string, level int32, reviewTime int64) (*models.ReviewData, error)
}
