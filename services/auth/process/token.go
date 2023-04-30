package process

import (
	"os"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// generateToken generates a token for a given identifier and
// returns the encoded token, its expiration time, and any potential errors
func generateToken(identifier string) (string, int64, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	expirationTime := time.Now().Add(time.Hour * 24 * 7)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"sub": identifier,
			"exp": jwt.NewNumericDate(expirationTime),
		},
	)

	signedToken, err := token.SignedString(signingKey)
	if err != nil {
		logger.Error("error signing token", err)
		return "", 0, status.Error(codes.Unknown, "error generating token")
	}

	return signedToken, expirationTime.Unix(), nil
}
