package roomservice

import (
	"context"
	"errors"
	roomv1 "gilwong00/connect-chat/gen/proto/room/v1"

	"github.com/bufbuild/connect-go"
)

func (s *roomService) CreateRoom(
	ctx context.Context,
	req *connect.Request[roomv1.CreateRoomRequest],
) (*connect.Response[roomv1.CreateRoomResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("unimplemented"))
}
