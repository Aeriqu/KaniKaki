package resolver

import (
	"context"

	"github.com/Aeriqu/kanikaki/services/api-gateway/clients"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
)

func (*MutationResolver) ChangePassword(ctx context.Context, mutation model.ChangePasswordRequest) (*model.IdentifierResponse, error) {
	response := &model.IdentifierResponse{}
	authResponse, err := clients.GetAuthClient().SendChangePassword(ctx, *mutation.Identifier, *mutation.OldCredential, *mutation.NewCredential)
	if err != nil {
		return response, err
	}
	response.Identifier = authResponse.Identifier
	return response, nil
}

func (*MutationResolver) Signup(ctx context.Context, mutation model.SignupRequest) (*model.TokenResponse, error) {
	response := &model.TokenResponse{}
	authResponse, err := clients.GetAuthClient().SendSignupRequest(ctx, *mutation.Identifier, *mutation.Credential)
	if err != nil {
		return response, err
	}
	response.Token = authResponse.Token
	return response, nil
}

func (*MutationResolver) Login(ctx context.Context, mutation model.LoginRequest) (*model.TokenResponse, error) {
	response := &model.TokenResponse{}
	authResponse, err := clients.GetAuthClient().SendLoginRequest(ctx, *mutation.Identifier, *mutation.Credential)
	if err != nil {
		return response, err
	}
	response.Token = authResponse.Token
	return response, nil
}

func (*MutationResolver) Logout(ctx context.Context, mutation model.LogoutRequest) (*model.TokenResponse, error) {
	response := &model.TokenResponse{}
	authResponse, err := clients.GetAuthClient().SendLogoutRequest(ctx, *mutation.Identifier)
	if err != nil {
		return response, err
	}
	response.Token = authResponse.Token
	return response, nil
}

func (*MutationResolver) RefreshToken(ctx context.Context, mutation model.RefreshTokenRequest) (*model.TokenResponse, error) {
	response := &model.TokenResponse{}
	authResponse, err := clients.GetAuthClient().SendRefreshTokenRequest(ctx, *mutation.Identifier)
	if err != nil {
		return response, err
	}
	response.Token = authResponse.Token
	return response, nil
}