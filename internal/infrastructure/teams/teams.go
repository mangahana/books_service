package teams

import (
	pb "books_service/proto/teams"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type service struct {
	client pb.TeamsClient
}

func New(socket string) (*service, error) {
	conn, err := grpc.NewClient(socket, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return &service{}, err
	}

	return &service{client: pb.NewTeamsClient(conn)}, nil
}
