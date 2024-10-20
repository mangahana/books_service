package repository

import "context"

func (r *repo) GetPages(c context.Context, chapterID string) ([]string, error) {
	output := []string{}
	sql := "SELECT image FROM pages WHERE chapter_id = $1 ORDER BY page_number ASC;"

	rows, err := r.db.Query(c, sql, chapterID)
	if err != nil {
		return output, err
	}

	for rows.Next() {
		var p string
		if err := rows.Scan(&p); err != nil {
			return output, err
		}
		output = append(output, p)
	}

	return output, nil
}
