package repository

import (
	"books_service/internal/core/cerror"
	"context"
)

func (r *repo) IsGenresExists(c context.Context, arr []int) error {
	var count int
	sql := "SELECT count(id) FROM genres WHERE id = any($1);"

	err := r.db.QueryRow(c, sql, arr).Scan(&count)
	if err != nil {
		return err
	}

	if count == len(arr) {
		return nil
	}

	return cerror.New(cerror.GENRES_NOT_FOUND, "genres not found")
}
