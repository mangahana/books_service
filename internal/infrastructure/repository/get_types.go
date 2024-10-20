package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetTypes(c context.Context) ([]models.BookType, error) {
	output := []models.BookType{}

	sql := "SELECT id, name FROM types;"
	rows, err := r.db.Query(c, sql)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var t models.BookType
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			return output, err
		}
		output = append(output, t)
	}

	return output, nil
}
