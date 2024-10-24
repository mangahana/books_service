package teams

import (
	"books_service/internal/core/models"
	pb "books_service/proto/teams"
	"context"
)

func (s *service) GetMember(c context.Context, teamId, memberId int) (models.TeamMember, error) {
	res, err := s.client.GetMember(c, &pb.GetMemberRequest{TeamId: int32(teamId), MemberId: int32(memberId)})
	if err != nil {
		return models.TeamMember{}, err
	}

	return models.TeamMember{
		UserID:      int(res.UserId),
		Username:    res.Username,
		Permissions: res.Permissions,
	}, nil
}
