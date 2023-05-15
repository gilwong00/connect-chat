package userservice

import (
	"context"
	"database/sql"
	"errors"
	userv1 "gilwong00/connect-chat/gen/proto/user/v1"
	"gilwong00/connect-chat/pkg/token"
	"strconv"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v4"
)

func (s *userService) WhoAmI(
	ctx context.Context,
	req *connect.Request[userv1.WhoAmIRequest],
) (*connect.Response[userv1.WhoAmIReponse], error) {
	authToken, err := token.GetAuthTokenFromHeader(req)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("no auth token"))
	}
	token, err := token.ParseJwtToken(authToken)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("error parsing token"))
	}
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	dbUser, err := s.queries.GetUserById(ctx, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("could not find user"))
		}
		return nil, err
	}
	return connect.NewResponse(&userv1.WhoAmIReponse{
		User: &userv1.User{
			Id:       dbUser.ID,
			Username: dbUser.Username,
			Email:    dbUser.Email,
		},
	}), nil
}
