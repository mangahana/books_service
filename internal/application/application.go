package application

import (
	"books_service/internal/core/models"
	"books_service/internal/infrastructure"
	"context"
)

type UseCase interface {
	GetBooks(c context.Context) ([]models.Book, error)
	GetOneByLink(c context.Context, bookLink string) (models.OneBook, error)
	GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error)
	GetPages(c context.Context, chapterID string) ([]string, error)

	GetTypes(c context.Context) ([]models.BookType, error)
	GetPersons(c context.Context) ([]models.Person, error)
	GetGenres(c context.Context) ([]models.Genre, error)
	GetStatuses(c context.Context) ([]models.Status, error)
}

type useCase struct {
	repo infrastructure.Repository
	s3   infrastructure.Storage
}

func New(repo infrastructure.Repository) UseCase {
	return &useCase{repo: repo}
}
