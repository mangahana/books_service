package repository

import "context"

func (r *repo) IsTypeExists(c context.Context, typeId int) error {
	var id int
	sql := "SELECT id FROM types WHERE id = $1;"
	return r.db.QueryRow(c, sql, typeId).Scan(&id)
}
