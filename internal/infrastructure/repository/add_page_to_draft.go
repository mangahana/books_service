package repository

import "context"

func (r *repo) AddPageToDraft(c context.Context, chapterID, page string) error {
	sql := "UPDATE chapters SET pages = array_append(pages, $2) WHERE id = $1;"
	_, err := r.db.Exec(c, sql, chapterID, page)
	return err
}
