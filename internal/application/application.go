package application

import (
	"books_service/internal/core/models"
	"books_service/internal/infrastructure/repository"
	"context"
)

type UseCase interface {
	GetTypes(c context.Context) ([]models.BookType, error)
}

type useCase struct {
	repo repository.Repository
	s3   IStorage
}

type IStorage interface {
	Put(c context.Context, file []byte) (string, error)
	Remove(c context.Context) error
}

func New(repo repository.Repository) UseCase {
	return &useCase{repo: repo}
}
