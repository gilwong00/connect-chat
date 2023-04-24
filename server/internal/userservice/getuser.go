package userservice

import (
	"context"
	"errors"
	userv1 "gilwong00/connect-chat/gen/proto/user/v1"

	"github.com/bufbuild/connect-go"
)

func (s *userService) GetUser(
	ctx context.Context,
	req *connect.Request[userv1.GetUserRequest],
) (*connect.Response[userv1.GetUserResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
