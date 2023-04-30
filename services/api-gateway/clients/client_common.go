package clients

import (
	"context"
	"net/http"
	"strings"

	tokenValidator "github.com/Aeriqu/kanikaki/common/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func getHeaders(ctx context.Context) (http.Header) {
	return ctx.Value("headers").(http.Header)
}

func getTokenFromHeaders(headers http.Header) (string) {
	authorization := headers.Get("Authorization")
	token := strings.Split(authorization, " ")
	if len(token) < 2 { return "" }
	return token[len(token)-1]
}

func validateToken(ctx context.Context) (string, string, error) {
	token := getTokenFromHeaders(getHeaders(ctx))
	tokenClaims, err := tokenValidator.GetClaims(token)
	if err != nil {
		return "", "", err
	}

	authResponse, err := GetAuthClient().SendValidateTokenRequest(ctx, token)
	if err != nil {
		return "", "", err
	}

	identifier, err := tokenClaims.GetSubject()
	if err != nil {
		return "", "", err
	} else if identifier != authResponse.Identifier {
		return "", "", status.Error(codes.Unauthenticated, "token claim identifier does not match server result identifier")
	}

	return authResponse.Identifier, token, nil
}