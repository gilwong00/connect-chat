package token

import (
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

func GetAuthTokenFromHeader(req connect.AnyRequest) (string, error) {
	authHeader := req.Header().Get("authorization")
	if !strings.Contains(authHeader, "Bearer") {
		return "", errors.New("missing auth header")
	}
	tokenParts := strings.Split(authHeader, " ")
	if len(tokenParts) == 1 {
		return "", errors.New("no auth token provided")
	}
	authToken := tokenParts[1]
	if authToken == "" {
		return "", errors.New("no auth token provided")
	}
	return authToken, nil
}

func ParseJwtToken(authToken string) (*jwt.Token, error) {
	claims := jwt.MapClaims{}
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}
	return jwt.ParseWithClaims(authToken, claims, keyFunc)
}
