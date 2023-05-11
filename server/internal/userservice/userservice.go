package userservice

import (
	"gilwong00/connect-chat/gen/proto/user/v1/userv1connect"
	db "gilwong00/connect-chat/pkg/db/sqlc"
)

type userService struct {
	queries *db.Queries
}

func NewUserService(queries *db.Queries) (userv1connect.UserServiceHandler, error) {
	return &userService{
		queries,
	}, nil
}
