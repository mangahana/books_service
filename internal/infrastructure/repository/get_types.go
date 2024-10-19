package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetTypes(c context.Context) ([]models.BookType, error) {
	var bookTypes []models.BookType

	sql := "SELECT id, name FROM types;"
	rows, err := r.db.Query(c, sql)
	if err != nil {
		return bookTypes, err
	}

	for rows.Next() {
		var bookType models.BookType
		if err := rows.Scan(&bookType.ID, &bookType.Name); err != nil {
			return bookTypes, err
		}
		bookTypes = append(bookTypes, bookType)
	}

	return bookTypes, nil
}
