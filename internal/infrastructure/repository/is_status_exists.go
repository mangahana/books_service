package repository

import "context"

func (r *repo) IsStatusExists(c context.Context, statusId int) error {
	var id int
	sql := "SELECT id FROM statuses WHERE id = $1;"
	return r.db.QueryRow(c, sql, statusId).Scan(&id)
}
