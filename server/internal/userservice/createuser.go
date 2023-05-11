package userservice

import (
	"context"
	"errors"
	userv1 "gilwong00/connect-chat/gen/proto/user/v1"
	db "gilwong00/connect-chat/pkg/db/sqlc"

	"github.com/bufbuild/connect-go"
	"golang.org/x/crypto/bcrypt"
)

func (s *userService) CreateUser(
	ctx context.Context,
	req *connect.Request[userv1.CreateUserRequest],
) (*connect.Response[userv1.CreateUserResponse], error) {
	// validate payload
	if req.Msg.Username == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("username is required"))
	}
	if req.Msg.Email == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("email is required"))
	}
	if req.Msg.Password == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("password is required"))
	}
	hashedPassword, err := hashPassword(req.Msg.Password)
	if err != nil {
		return nil, err
	}
	// TODO implement transaction
	newUser, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Username: req.Msg.Username,
		Email:    req.Msg.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&userv1.CreateUserResponse{
		User: &userv1.User{
			Id:       newUser.ID,
			Username: newUser.Username,
			Email:    newUser.Email,
		},
	}), nil
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("failed to hash password")
	}
	return string(hash), nil
}
