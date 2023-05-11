package userservice

import (
	"context"
	"database/sql"
	"errors"
	userv1 "gilwong00/connect-chat/gen/proto/user/v1"
	db "gilwong00/connect-chat/pkg/db/sqlc"
	"strconv"
	"time"

	"github.com/bufbuild/connect-go"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	secretKey = "secret"
)

type jWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *userService) Login(
	ctx context.Context,
	req *connect.Request[userv1.LoginRequest],
) (*connect.Response[userv1.LoginResponse], error) {
	// validate
	if req.Msg.Username == "" && req.Msg.Email == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("username or email is required"))
	}
	if req.Msg.Password == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("password is required"))
	}
	var user db.User
	var err error
	if req.Msg.Email != "" {
		user, err = s.queries.GetUserByEmail(ctx, req.Msg.Email)
	} else {
		user, err = s.queries.GetUserByUsername(ctx, req.Msg.Username)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, connect.NewError(connect.CodeNotFound, errors.New("unable to find user"))
		}
		return nil, err
	}
	// check password
	err = checkPassword(req.Msg.Password, user.Password)
	if err != nil {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("invalid password"))
	}
	// create access token
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})
	accessToken, err := jwt.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	res := connect.NewResponse(&userv1.LoginResponse{
		AccessToken: accessToken,
		User: &userv1.User{
			Id:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
	// maybe we dont need to set it as a header?
	res.Header().Add("Set-Cookie", accessToken)
	return res, nil
}

func checkPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
