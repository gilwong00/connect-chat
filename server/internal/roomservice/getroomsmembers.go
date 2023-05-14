package roomservice

import (
	"context"
	roomv1 "gilwong00/connect-chat/gen/proto/room/v1"

	"github.com/bufbuild/connect-go"
)

func (s *roomService) GetRoomMembers(
	ctx context.Context,
	req *connect.Request[roomv1.GetRoomMembersRequest],
) (*connect.Response[roomv1.GetRoomsMemberResponse], error) {
	members := make([]*roomv1.Member, 0)
	for _, member := range s.hubClient.GetClientsFromRoom(req.Msg.RoomId) {
		members = append(members, &roomv1.Member{
			MemberId:   member.ID,
			MemberName: member.Username,
		})
	}
	return connect.NewResponse(&roomv1.GetRoomsMemberResponse{
		Members: members,
	}), nil
}
