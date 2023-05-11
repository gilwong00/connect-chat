package roomservice

import (
	"gilwong00/connect-chat/gen/proto/room/v1/roomv1connect"
	db "gilwong00/connect-chat/pkg/db/sqlc"
	"gilwong00/connect-chat/pkg/ws"
)

type roomService struct {
	queries   *db.Queries
	hubClient *ws.HubHandler
}

func NewRoomService(queries *db.Queries, hubHandler *ws.HubHandler) (roomv1connect.RoomServiceHandler, error) {
	return &roomService{queries, hubHandler}, nil
}
