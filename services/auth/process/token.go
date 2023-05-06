package process

import (
	"os"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/auth/models"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// generateToken generates a token for a given identifier and
// returns the encoded token, its expiration time, and any potential errors
func generateToken(user *models.User) (string, int64, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	expirationTime := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"sub": user.Username,
			"exp": jwt.NewNumericDate(expirationTime),
			"type": user.Type,
			"wanikani_level_limit": user.WanikaniLevelLimit,
		},
	)

	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		logger.Error("error signing token", err)
		return "", 0, status.Error(codes.Unknown, "error generating token")
	}

	return signedToken, expirationTime.Unix(), nil
}
