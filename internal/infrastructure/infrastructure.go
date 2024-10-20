package infrastructure

import (
	"books_service/internal/core/models"
	"context"
)

type AuthService interface {
	Authenticate(c context.Context, token string) (models.User, error)
}

type Storage interface {
	Put(c context.Context, file []byte) (string, error)
	Remove(c context.Context) error
}

type Repository interface {
	// db connection close
	Close()

	GetBooks(c context.Context) ([]models.Book, error)
	GetOneByLink(c context.Context, bookLink string) (models.OneBook, error)
	GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error)
	GetPages(c context.Context, chapterID string) ([]string, error)

	GetTypes(c context.Context) ([]models.BookType, error)
	GetPersons(c context.Context) ([]models.Person, error)
	GetGenres(c context.Context) ([]models.Genre, error)
	GetStatuses(c context.Context) ([]models.Status, error)
}
