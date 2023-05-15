package interceptors

import (
	"context"
	"errors"
	"gilwong00/connect-chat/pkg/token"

	"github.com/bufbuild/connect-go"
)

func NewAuthInterceptor() connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			authToken, err := token.GetAuthTokenFromHeader(req)
			if err != nil {
				return nil, connect.NewError(connect.CodeUnauthenticated, err)
			}
			token, err := token.ParseJwtToken(authToken)
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
