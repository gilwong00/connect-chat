package roomservice

import "gilwong00/connect-chat/gen/proto/room/v1/roomv1connect"

type roomService struct{}

func NewRoomService() (roomv1connect.RoomServiceHandler, error) {
	return &roomService{}, nil
}
