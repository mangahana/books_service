package application

import (
	"books_service/internal/core/models"
	"books_service/internal/infrastructure/repository"
	"context"
)

type UseCase interface {
	GetOneByLink(c context.Context, bookLink string) (models.OneBook, error)

	GetBooks(c context.Context) ([]models.Book, error)
	GetTypes(c context.Context) ([]models.BookType, error)
	GetPersons(c context.Context) ([]models.Person, error)
	GetGenres(c context.Context) ([]models.Genre, error)
	GetStatuses(c context.Context) ([]models.Status, error)
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
