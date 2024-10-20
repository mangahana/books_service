package infrastructure

import (
	"books_service/internal/core/dto"
	"books_service/internal/core/models"
	"context"
)

type AuthService interface {
	Authenticate(c context.Context, token string) (models.User, error)
}

type TeamsService interface {
	GetTeam(c context.Context, teamId int) (models.Team, error)
	GetMember(c context.Context, teamId, memberId int) (models.TeamMember, error)
}

type Storage interface {
	Put(c context.Context, file []byte) (string, error)
	Remove(c context.Context) error
}

type Repository interface {
	CreateBook(c context.Context, dto *dto.AddBook) error

	GetBooks(c context.Context) ([]models.Book, error)
	GetOneByLink(c context.Context, bookLink string) (models.OneBook, error)
	GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error)
	GetPages(c context.Context, chapterID string) ([]string, error)

	GetTypes(c context.Context) ([]models.BookType, error)
	GetPersons(c context.Context) ([]models.Person, error)
	GetGenres(c context.Context) ([]models.Genre, error)
	GetStatuses(c context.Context) ([]models.Status, error)

	IsTypeExists(c context.Context, typeId int) error
	IsStatusExists(c context.Context, statusId int) error
	IsPersonsExists(c context.Context, arr []int) error
	IsGenresExists(c context.Context, arr []int) error

	// db connection close
	Close()
}
