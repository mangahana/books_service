package repository

import (
	"books_service/internal/core/dto"
	"context"
)

func (r *repo) AddPerson(c context.Context, person *dto.AddPerson) error {
	sql := "INSERT INTO persons (name) VALUES($1);"
	_, err := r.db.Exec(c, sql, person.Name)
	return err
}
