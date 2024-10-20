package repository

import (
	"books_service/internal/core/models"
	"context"
	"fmt"
)

func (r *repo) GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error) {
	output := []models.Chapter{}

	sql := "SELECT id, name, volume, chapter_number, created_at FROM chapters WHERE book_id = $1 ORDER BY chapter_number %s;"

	if endSorting {
		sql = fmt.Sprintf(sql, "DESC")
	} else {
		sql = fmt.Sprintf(sql, "ASC")
	}

	rows, err := r.db.Query(c, sql, bookId)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var c models.Chapter
		if err := rows.Scan(&c.ID, &c.Name, &c.Volume, &c.Number, &c.CreatedAt); err != nil {
			return output, err
		}
		output = append(output, c)
	}

	return output, nil
}
