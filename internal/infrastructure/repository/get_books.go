package repository

import (
	"books_service/internal/core/models"
	"context"
	"log"
)

func (r *repo) GetBooks(c context.Context) ([]models.Book, error) {
	log.Println("here")
	output := []models.Book{}
	sql := `SELECT id, link, name, poster,
					(SELECT name FROM types WHERE id = books.type_id) as type
					FROM books;`

	rows, err := r.db.Query(c, sql)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Link, &book.Name, &book.Poster, &book.Type); err != nil {
			return output, err
		}
		output = append(output, book)
	}

	return output, nil
}
