package auth_service

import (
	"books_service/internal/core/models"
	"context"
)

type service struct{}

func New() *service {
	return &service{}
}

func (s *service) Authenticate(c context.Context, token string) (models.User, error) {
	return models.User{}, nil
}
