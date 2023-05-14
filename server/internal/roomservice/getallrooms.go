package roomservice

import (
	"context"
	roomv1 "gilwong00/connect-chat/gen/proto/room/v1"

	"github.com/bufbuild/connect-go"
)

func (s *roomService) GetAllRooms(
	ctx context.Context,
	req *connect.Request[roomv1.GetAllRoomsRequest],
) (*connect.Response[roomv1.GetAllRoomsResponse], error) {
	rooms := make([]*roomv1.Room, 0)
	for _, room := range s.hubClient.GetAllRooms() {
		rooms = append(rooms, &roomv1.Room{
			RoomId:   room.ID,
			RoomName: room.Name,
		})
	}
	return connect.NewResponse(&roomv1.GetAllRoomsResponse{
		Rooms: rooms,
	}), nil
}
