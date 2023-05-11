package userservice

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	userv1 "gilwong00/connect-chat/gen/proto/user/v1"

	"github.com/bufbuild/connect-go"
)

func (s *userService) GetUserByEmail(
	ctx context.Context,
	req *connect.Request[userv1.GetUserByEmailRequest],
) (*connect.Response[userv1.GetUserByEmailResponse], error) {
	// validate
	if req.Msg.Email == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("email is required"))
	}
	user, err := s.queries.GetUserByEmail(ctx, req.Msg.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("could not find user with email: %s", req.Msg.Email))
		}
		return nil, err
	}
	return connect.NewResponse(&userv1.GetUserByEmailResponse{
		User: &userv1.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}), nil
}
