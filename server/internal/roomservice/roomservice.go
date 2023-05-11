package roomservice

import (
	"gilwong00/connect-chat/gen/proto/room/v1/roomv1connect"
	db "gilwong00/connect-chat/pkg/db/sqlc"
)

type roomService struct {
	queries *db.Queries
}

func NewRoomService(queries *db.Queries) (roomv1connect.RoomServiceHandler, error) {
	return &roomService{queries}, nil
}
