package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetGenres(c context.Context) ([]models.Genre, error) {
	var output []models.Genre
	sql := "SELECT id, name FROM genres;"

	rows, err := r.db.Query(c, sql)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var genre models.Genre
		if err := rows.Scan(&genre.ID, &genre.Name); err != nil {
			return output, err
		}
		output = append(output, genre)
	}

	return output, nil
}
