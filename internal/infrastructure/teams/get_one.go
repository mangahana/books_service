package teams

import (
	"books_service/internal/core/models"
	pb "books_service/proto/teams"
	"context"
)

func (s *service) GetOne(c context.Context, teamId int) (models.Team, error) {
	res, err := s.client.GetOne(c, &pb.GetOneRequest{Id: int32(teamId)})
	if err != nil {
		return models.Team{}, err
	}

	return models.Team{
		ID:          int(res.Id),
		Name:        res.Name,
		Photo:       res.Photo,
		IsModerated: res.IsModerated,
		OwnerId:     int(res.OwnerId),
	}, nil
}
