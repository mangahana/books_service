package repository

import (
	"books_service/internal/core/models"
	"context"
)

func (r *repo) GetPersons(c context.Context) ([]models.Person, error) {
	var output []models.Person
	sql := "SELECT id, name FROM persons;"

	rows, err := r.db.Query(c, sql)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var person models.Person
		if err := rows.Scan(&person.ID, &person.Name); err != nil {
			return output, err
		}
		output = append(output, person)
	}

	return output, nil
}
