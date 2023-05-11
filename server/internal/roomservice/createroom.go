package roomservice

import (
	"context"
	"errors"
	roomv1 "gilwong00/connect-chat/gen/proto/room/v1"

	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"
)

func (s *roomService) CreateRoom(
	ctx context.Context,
	req *connect.Request[roomv1.CreateRoomRequest],
) (*connect.Response[roomv1.CreateRoomResponse], error) {
	// validate
	if req.Msg.Name == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, errors.New("room name is required"))
	}
	roomID := uuid.New().String()
	s.hubClient.AppendNewRoom(roomID, req.Msg.Name)
	return connect.NewResponse(&roomv1.CreateRoomResponse{
		RoomId: roomID,
		Name:   req.Msg.Name,
	}), nil
}
