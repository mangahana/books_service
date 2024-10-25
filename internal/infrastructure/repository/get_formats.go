package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetFormats(c context.Context) ([]models.Format, error) {
	output := []models.Format{}
	sql := "SELECT id, name FROM formats;"
	rows, err := r.db.Query(c, sql)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var f models.Format
		if err := rows.Scan(&f.ID, &f.Name); err != nil {
			return output, err
		}
		output = append(output, f)
	}

	return output, nil
}
