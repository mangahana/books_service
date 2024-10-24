package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetOneByID(c context.Context, id int) (models.OneBook, error) {
	var b models.OneBook
	sql := `SELECT id, link, name, original_title, description, poster, release_date,
					(SELECT name FROM types WHERE id = books.type_id) as type,
					(SELECT name FROM statuses WHERE id = books.status_id) as status,
					(SELECT array_agg(name) FROM genres WHERE id = any(books.genres)) as genres,
					(SELECT array_agg(name) FROM persons WHERE id = any(books.authors_ids)) as authors,
					(SELECT array_agg(name) FROM persons WHERE id = any(books.artists_ids)) as artists
				  FROM books WHERE id = $1;`

	err := r.db.QueryRow(c, sql, id).Scan(
		&b.ID, &b.Link, &b.Name, &b.OriginalTitle,
		&b.Description, &b.Poster, &b.ReleaseDate,
		&b.Type, &b.Status, &b.Genres, &b.Authors, &b.Artists,
	)

	return b, err
}
