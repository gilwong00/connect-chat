package userservice

import "gilwong00/connect-chat/gen/proto/user/v1/userv1connect"

type userService struct{}

func NewUserService() (userv1connect.UserServiceHandler, error) {
	return &userService{}, nil
}
