package roomservice

import (
	"context"
	"errors"
	roomv1 "gilwong00/connect-chat/gen/proto/room/v1"

	"github.com/bufbuild/connect-go"
)

func (s *roomService) GetRoomMembers(
	ctx context.Context,
	req *connect.Request[roomv1.GetRoomMembersRequest],
) (*connect.Response[roomv1.GetRoomsMemberResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("not"))
}
