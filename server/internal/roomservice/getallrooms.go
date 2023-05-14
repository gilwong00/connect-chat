package roomservice

import (
	"context"
	"errors"
	roomv1 "gilwong00/connect-chat/gen/proto/room/v1"

	"github.com/bufbuild/connect-go"
)

func (s *roomService) GetAllRooms(
	ctx context.Context,
	req *connect.Request[roomv1.GetAllRoomsRequest],
) (*connect.Response[roomv1.GetAllRoomsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not"))
}
