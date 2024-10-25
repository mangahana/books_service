package repository

import "context"

func (r *repo) IsDraftExists(c context.Context, chapterID string, userID int) error {
	var id string
	sql := "SELECT id FROM chapters WHERE id = $1 AND member_id = $2 AND is_draft = true;"
	return r.db.QueryRow(c, sql, chapterID, userID).Scan(&id)
}
