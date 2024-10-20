package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetStatuses(c context.Context) ([]models.Status, error) {
	output := []models.Status{}
	sql := "SELECT id, name FROM statuses;"

	rows, err := r.db.Query(c, sql)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var status models.Status
		if err := rows.Scan(&status.ID, &status.Name); err != nil {
			return output, err
		}
		output = append(output, status)
	}

	return output, nil
}
