// Package process handles the processing required to
// handle the auth requests
package process

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/Aeriqu/kanikaki/common/logger"
	tokenValidator "github.com/Aeriqu/kanikaki/common/token"
	"github.com/Aeriqu/kanikaki/services/auth/database"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func hashCredential(credential string) string {
	salt := os.Getenv("CREDENTIAL_SALT")
	hasher := sha512.New()
	hasher.Write([]byte(credential + salt))

	return hex.EncodeToString(hasher.Sum(nil))
}

func ChangePassword(db *database.Database, identifier string, oldCredential string, newCredential string, token string) (string, error) {
	logger.Info(fmt.Sprintf("attempting to change password (%s)", identifier))

	oldHashedCredential := hashCredential(oldCredential)
	newHashedCredential := hashCredential(newCredential)

	user, err := db.GetUserByIdentifier(identifier)
	if err != nil {
		return "", err
	}

	if user.Password != oldHashedCredential {
		logger.Info(fmt.Sprintf("attempted password change for user (%s) but was provided invalid credential", identifier))
		return "", status.Error(codes.Unauthenticated, "old credential is not valid")
	}

	logger.Info(fmt.Sprintf("changed password (%s)", identifier))

	if _, err := db.UpdateUserCredential(identifier, newHashedCredential); err != nil {
		return "", err
	}

	return identifier, nil
}

func Login(db *database.Database, identifier string, credential string) (string, error) {
	logger.Info(fmt.Sprintf("attemping to login (%s)", identifier))

	hashedCredential := hashCredential(credential)
	user, err := db.GetUserByIdentifier(identifier)
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "identifier not found")
	}

	if user.Password != hashedCredential {
		return "", status.Error(codes.Unauthenticated, "credentials do not match")
	}

	token, tokenExpirationTime, err := generateToken(identifier)
	if err != nil {
		return "", err
	}

	// Remove expired tokens from the user (storage purposes)
	db.RemoveUserExpiredTokens(identifier)

	// Generate new token
	// TODO: Handle update success / fail logic
	db.AddTokenToUser(identifier, token, tokenExpirationTime)

	logger.Info(fmt.Sprintf("user logged in successfully (%s)", identifier))

	return token, nil
}

func Logout(db *database.Database, identifier string, token string) (string, error) {
	logger.Info(fmt.Sprintf("attempting to log out (%s)", identifier))

	db.RemoveSpecificUserToken(identifier, token)

	return token, nil
}

func RefreshToken(db *database.Database, identifier string, token string) (string, error) {
	logger.Info(fmt.Sprintf("attempting to refresh token (%s)", identifier))

	// Check token validity
	claims, err := tokenValidator.GetClaims(token)
	if err != nil {
		return "", status.Error(codes.Unauthenticated, "invalid token provided")
	}

	// Check if the token is even intended for the user
	subject, _ := claims.GetSubject()
	if identifier != subject {
		return "", status.Error(codes.Unauthenticated, "identifier and token do not match")
	}

	logger.Info(fmt.Sprintf("refreshing token for user (%s)", identifier))

	// Generate the token and replace it
	newToken, newTokenExpirationTime, err := generateToken(identifier)
	if err != nil {
		return "", err
	}

	updatedUser, err := db.UpdateSpecificUserToken(identifier, token, newToken, newTokenExpirationTime)
	if err != nil {
		return "", err
	}

	// double check the tokens and ensure it's in there
	for _, userToken := range updatedUser.AuthTokens {
		if userToken.Jwt == newToken {
			logger.Info(fmt.Sprintf("refreshed token for user (%s)", identifier))
			return newToken, nil
		}
	}

	logger.Info(fmt.Sprintf("unable to find refreshed token for user (%s)", identifier))

	return "", status.Error(codes.NotFound, "could not refresh the given token")
}

func Signup(db *database.Database, identifier string, credential string) (string, error) {
	logger.Info(fmt.Sprintf("attempting to create account (%s)", identifier))

	hashedCredential := hashCredential(credential)
	if _, err := db.AddUser(identifier, hashedCredential); err != nil {
		return "", err
	}

	logger.Info(fmt.Sprintf("created user (%s)", identifier))

	token, tokenExpirationTime, err := generateToken(identifier)
	if err != nil {
		return "", err
	}

	// TODO: Handle update success / fail logic
	db.AddTokenToUser(identifier, token, tokenExpirationTime)

	return token, nil
}

func ValidateToken(db *database.Database, token string) (string, error) {
	logger.Info("attempting to validate token")

	user, err := db.GetUserByToken(token)
	if err != nil {
		return "", err
	}

	logger.Info(fmt.Sprintf("found user by token (%s)", user.Username))

	return user.Username, nil
}