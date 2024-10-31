package application

import (
	"books_service/internal/core/dto"
	"books_service/internal/core/models"
	"books_service/internal/infrastructure"
	"context"
	"log"
)

type UseCase interface {
	Add(c context.Context, user models.User, dto dto.AddBook) error
	CreateDraft(c context.Context, user *models.User, dto *dto.CreateDraft) (string, error)
	UploadPage(c context.Context, user *models.User, dto *dto.UploadPage) (string, error)
	AddPerson(c context.Context, person *dto.AddPerson) error

	GetBooks(c context.Context) ([]models.Book, error)
	GetOneByLink(c context.Context, bookLink string) (models.OneBook, error)
	GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error)
	GetPages(c context.Context, chapterID string) ([]string, error)

	GetPersons(c context.Context) ([]models.Person, error)

	GetParams(c context.Context) (map[string]any, error)
}

type useCase struct {
	repo infrastructure.Repository
	s3   infrastructure.Storage

	teamsService infrastructure.TeamsService
}

func New(repo infrastructure.Repository, teamsService infrastructure.TeamsService, s3 infrastructure.Storage) UseCase {
	log.Println(s3)
	return &useCase{
		repo:         repo,
		teamsService: teamsService,
		s3:           s3,
	}
}
