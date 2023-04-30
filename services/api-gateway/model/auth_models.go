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

type ValidateTokenRequest struct {
	Token string
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
