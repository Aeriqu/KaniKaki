// Package token handles the validation of tokens
package token

import (
	"fmt"
	"os"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateTokenMethod(token *jwt.Token) (interface{}, error) {
	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		logger.Info(fmt.Sprintf("token with invalid signing method (%s)", token.Raw))
		return nil, status.Error(codes.Aborted, "token with invalid signing method")
	}

	return []byte(signingKey), nil
}

// GetClaims validates the token and ensures the algorithm,
// HMAC, and expiration are proper.
func GetClaims(token string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims

	tokenData, err := jwt.Parse(token, validateTokenMethod)
	if err != nil {
		logger.Error("error parsing token", err)
		return claims, status.Error(codes.Aborted, "error parsing token")
	}

	if claims, ok := tokenData.Claims.(jwt.MapClaims); ok && tokenData.Valid {
		expirationTime, _ := claims.GetExpirationTime()
		if expirationTime.After(time.Now()) {
			return claims, nil
		}
	}

	logger.Info("invalid token: " + token)
	return claims, status.Error(codes.Aborted, "provided token is invalid")
}
