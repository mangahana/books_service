package infrastructure

import (
	"books_service/internal/core/dto"
	"books_service/internal/core/models"
	"context"
)

type AuthService interface {
	GetUser(c context.Context, token string) (models.User, error)
}

type TeamsService interface {
	GetOne(c context.Context, teamId int) (models.Team, error)
	GetMember(c context.Context, teamId, memberId int) (models.TeamMember, error)
}

type Storage interface {
	Put(c context.Context, file []byte) (string, error)
	Remove(c context.Context, filename string) error
}

type Repository interface {
	CreateBook(c context.Context, dto *dto.AddBook) error
	CreateDraft(c context.Context, memberId int, teamName string, dto *dto.CreateDraft) (string, error)
	AddPageToDraft(c context.Context, chapterID, page string) error

	GetBooks(c context.Context) ([]models.Book, error)
	GetOneByLink(c context.Context, bookLink string) (models.OneBook, error)
	GetOneByID(c context.Context, id int) (models.OneBook, error)
	GetDraftChapterID(c context.Context, memberId int, dto *dto.CreateDraft) (string, error)
	GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error)
	GetPages(c context.Context, chapterID string) ([]string, error)

	GetTypes(c context.Context) ([]models.BookType, error)
	GetPersons(c context.Context) ([]models.Person, error)
	GetGenres(c context.Context) ([]models.Genre, error)
	GetStatuses(c context.Context) ([]models.Status, error)
	GetFormats(c context.Context) ([]models.Format, error)

	IsTypeExists(c context.Context, typeId string) error
	IsStatusExists(c context.Context, statusId string) error
	IsPersonsExists(c context.Context, arr []string) error
	IsGenresExists(c context.Context, arr []string) error
	IsFormatExists(c context.Context, arr []string) error
	IsDraftExists(c context.Context, chapterID string, userID int) error

	// db connection close
	Close()
}
