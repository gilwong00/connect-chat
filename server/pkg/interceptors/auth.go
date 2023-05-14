package interceptors

import (
	"context"
	"errors"
	"strings"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v4"
)

const (
	secretKey = "secret"
)

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			authHeader := req.Header().Get("authorization")
			if !strings.Contains(authHeader, "Bearer") {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("missing auth header"))
			}
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) == 1 {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no auth token provided"))
			}
			authToken := tokenParts[1]
			if authToken == "" {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("no auth token provided"))
			}
			claims := jwt.MapClaims{}
			keyFunc := func(token *jwt.Token) (interface{}, error) {
				return []byte(secretKey), nil
			}
			token, err := jwt.ParseWithClaims(authToken, claims, keyFunc)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("error parsing token"))
			}
			if !token.Valid {
				return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("invalid auth token"))
			}
			return next(ctx, req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}
