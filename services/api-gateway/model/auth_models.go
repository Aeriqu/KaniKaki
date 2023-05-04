// Package model contains all of the packages for graphql related requests
// this ends being similar, but different from the protobuf models because
// graphql requires a pointers under certain conditions.
package model

type ChangePasswordRequest struct {
	Identifier *string
	OldCredential *string
	NewCredential *string
}

type LoginRequest struct {
	Identifier *string
	Credential *string
}

type LogoutRequest struct {
	Identifier *string
}

type RefreshTokenRequest struct {
	Identifier *string
}

type SignupRequest struct {
	Identifier *string
	Credential *string
}

type TokenResponse struct {
	Token string
}

type IdentifierResponse struct {
	Identifier string
}
