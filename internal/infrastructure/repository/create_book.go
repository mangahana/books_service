package repository

import (
	"books_service/internal/core/dto"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (r *repo) CreateBook(c context.Context, dto *dto.AddBook) error {
	sql := `INSERT INTO books (link, name, original_title, description, poster, genres, authors_ids, artists_ids, status_id, type_id, release_date, owner_team_id)
					VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12);`
	err := r.db.QueryRow(
		c, sql,
		dto.Link, dto.Name, dto.OriginalTitle, dto.Description,
		dto.Poster, dto.Genres, dto.Authors, dto.Artists,
		dto.StatusId, dto.TypeId, dto.ReleaseDate, dto.OwnerTeamID,
	).Scan()

	if err == nil || errors.Is(err, pgx.ErrNoRows) {
		return nil
	}

	return err
}
