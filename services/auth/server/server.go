// Package server implements all of the gRPC related server request handling.
package server

import (
	"context"

	"github.com/Aeriqu/kanikaki/services/auth/database"
	"github.com/Aeriqu/kanikaki/services/auth/process"
	authpb "github.com/Aeriqu/kanikaki/services/auth/proto"
)

type AuthServer struct {
	authpb.UnimplementedAuthServer
	*database.Database
}

func (server *AuthServer) ChangePassword(ctx context.Context, req *authpb.ChangePasswordRequest) (*authpb.IdentifierResponse, error) {
	identifier, err := process.ChangePassword(server.Database, req.Identifier, req.OldCredential, req.NewCredential, req.Token)
	if err != nil {
		return nil, err
	}
	response := &authpb.IdentifierResponse{
		Identifier: identifier,
	}
	return response, nil
}

func (server *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.TokenResponse, error) {
	token, err := process.Login(server.Database, req.Identifier, req.Credential)
	if err != nil {
		return nil, err
	}
	response := &authpb.TokenResponse{
		Token: token,
	}
	return response, nil
}

func (server *AuthServer) Logout(ctx context.Context, req *authpb.LogoutRequest) (*authpb.TokenResponse, error) {
	token, err := process.Logout(server.Database, req.Identifier, req.Token)
	if err != nil {
		return nil, err
	}
	response := &authpb.TokenResponse{
		Token: token,
	}
	return response, nil
}

func (server *AuthServer) RefreshToken(ctx context.Context, req *authpb.RefreshTokenRequest) (*authpb.TokenResponse, error) {
	token, err := process.RefreshToken(server.Database, req.Identifier, req.Token)
	if err != nil {
		return nil, err
	}
	response := &authpb.TokenResponse{
		Token: token,
	}
	return response, nil
}

func (server *AuthServer) Signup(ctx context.Context, req *authpb.SignupRequest) (*authpb.TokenResponse, error) {
	token, err := process.Signup(server.Database, req.Identifier, req.Credential)
	if err != nil {
		return nil, err
	}
	response := &authpb.TokenResponse{
		Token: token,
	}
	return response, nil
}

func (server *AuthServer) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.IdentifierResponse, error) {
	identifier, err := process.ValidateToken(server.Database, req.Token)
	if err != nil {
		return nil, err
	}
	response := &authpb.IdentifierResponse{
		Identifier: identifier,
	}
	return response, nil
}
