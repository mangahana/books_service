package repository

import (
	"books_service/internal/core/cerror"
	"context"
)

func (r *repo) IsPersonsExists(c context.Context, arr []int) error {
	var count int
	sql := "SELECT count(id) FROM persons WHERE id = any($1);"

	err := r.db.QueryRow(c, sql, arr).Scan(&count)
	if err != nil {
		return err
	}

	if count == len(arr) {
		return nil
	}

	return cerror.New(cerror.PERSONS_NOT_FOUND, "persons not found")
}
