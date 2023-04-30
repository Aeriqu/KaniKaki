package clients

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Aeriqu/kanikaki/common/logger"
	"github.com/Aeriqu/kanikaki/services/api-gateway/model"
	authpb "github.com/Aeriqu/kanikaki/services/auth/proto"
)

var authClientInstance *AuthClient
var authClientOnce sync.Once

type AuthClient struct {
	grpcClient authpb.AuthClient
}

// SendChangePassword sends a request to the auth service to change a given user's password.
// This requires the Authorization header to be set.
func (client *AuthClient) SendChangePassword(ctx context.Context, identifier string, oldCredential string, newCredential string) (model.IdentifierResponse, error) {
	_, token, err := validateToken(ctx)
	if err != nil {
		return model.IdentifierResponse{}, err
	}

	request := &authpb.ChangePasswordRequest{
		Identifier: identifier,
		OldCredential: oldCredential,
		NewCredential: newCredential,
		Token: token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	authResponse, err := client.grpcClient.ChangePassword(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with identifier (%s)", identifier), err)
		return model.IdentifierResponse{}, err
	}

	return model.IdentifierResponse{
		Identifier: authResponse.Identifier,
	}, nil
}

// SendLoginRequest sends a request to the auth service to obtain a jwt token with
// the given credentials.
func (client *AuthClient) SendLoginRequest(ctx context.Context, identifier string, credential string) (model.TokenResponse, error) {
	request := &authpb.LoginRequest{
		Identifier: identifier,
		Credential: credential,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	authResponse, err := client.grpcClient.Login(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with identifier (%s)", identifier), err)
		return model.TokenResponse{}, err
	}

	return model.TokenResponse{
		Token: authResponse.Token,
	}, nil
}

// SendLogoutRequest sends a request to the auth service to remove the jwt token
// from the list of valid jwt tokens for the user.
// This requires the Authorization header to be set.
func (client *AuthClient) SendLogoutRequest(ctx context.Context, identifier string) (model.TokenResponse, error) {
	_, token, err := validateToken(ctx)
	if err != nil {
		return model.TokenResponse{}, err
	}

	request := &authpb.LogoutRequest{
		Identifier: identifier,
		Token: token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	authResponse, err := client.grpcClient.Logout(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with token (%s)", token), err)
		return model.TokenResponse{}, err
	}

	return model.TokenResponse{
		Token: authResponse.Token,
	}, nil
}

// SendRefreshTokenRequest sends a request to the auth service to invalidate the
// current token and provide a new one that has a refreshed expiration.
// This requires the Authorization header to be set.
func (client *AuthClient) SendRefreshTokenRequest(ctx context.Context, identifier string) (model.TokenResponse, error) {
	_, token, err := validateToken(ctx)
	if err != nil {
		return model.TokenResponse{}, err
	}

	request := &authpb.RefreshTokenRequest{
		Identifier: identifier,
		Token: token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	authResponse, err := client.grpcClient.RefreshToken(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with token (%s)", token), err)
		return model.TokenResponse{}, err
	}

	return model.TokenResponse{
		Token: authResponse.Token,
	}, nil
}

// SendSignupRequest sends a request to create a new user with the provided credentials.
func (client *AuthClient) SendSignupRequest(ctx context.Context, identifier string, credential string) (model.TokenResponse, error) {
	request := &authpb.SignupRequest{
		Identifier: identifier,
		Credential: credential,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	authResponse, err := client.grpcClient.Signup(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with identifier (%s)", identifier), err)
		return model.TokenResponse{}, err
	}

	return model.TokenResponse{
		Token: authResponse.Token,
	}, nil
}

// SendValidateTokenRequest sends a request to ensure the token is valid. Further
// processing is required to ensure that the token is meant for the current
// requesting identity. This should be used whenever a request uses a token.
func (client *AuthClient) SendValidateTokenRequest(ctx context.Context, token string) (model.IdentifierResponse, error) {
	request := &authpb.ValidateTokenRequest{
		Token: token,
	}
	requestContext, requestCancel := context.WithTimeout(ctx, time.Second)
	defer requestCancel()
	authResponse, err := client.grpcClient.ValidateToken(requestContext, request)
	if err != nil {
		logger.Error(fmt.Sprintf("error sending request with token (%s)", token), err)
		return model.IdentifierResponse{}, err
	}

	return model.IdentifierResponse{
		Identifier: authResponse.Identifier,
	}, nil
}

func GetAuthClient() *AuthClient {
	authClientOnce.Do(func() {
		authClientInstance = &AuthClient{
			grpcClient: authpb.NewAuthClient(
				getConnection("auth-service.kanikaki.svc.cluster.local:80"),
			),
		}
	})
	return authClientInstance
}
