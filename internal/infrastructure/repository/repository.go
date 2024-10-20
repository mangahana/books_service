package repository

import (
	"books_service/internal/core/configuration"
	"books_service/internal/core/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

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

type repo struct {
	db *pgxpool.Pool
}

func New(config *configuration.DBConfig) (Repository, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s", config.User, config.Pass, config.Host, config.Name)
	conn, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &repo{db: conn}, nil
}

func (r *repo) Close() {
	r.db.Close()
}
