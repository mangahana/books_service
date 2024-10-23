package repository

import (
	"books_service/internal/core/cerror"
	"context"
)

func (r *repo) IsFormatExists(c context.Context, arr []string) error {
	var count int
	sql := "SELECT count(id) FROM formats WHERE id = any($1);"

	err := r.db.QueryRow(c, sql, arr).Scan(&count)
	if err != nil {
		return err
	}

	if count == len(arr) {
		return nil
	}

	return cerror.New(cerror.FORMAT_NOT_FOUND, "format not found")
}
