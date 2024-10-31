package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetPersonsByName(c context.Context, name string) ([]models.Person, error) {
	output := []models.Person{}

	sql := "SELECT id, name FROM persons WHERE name LIKE '%$1%'"
	rows, err := r.db.Query(c, sql, name)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var p models.Person
		if err := rows.Scan(&p.ID, &p.Name); err != nil {
			return output, err
		}
		output = append(output, p)
	}

	return output, nil
}
