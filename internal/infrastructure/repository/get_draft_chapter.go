package repository

import (
	"books_service/internal/core/dto"
	"context"
)

func (r *repo) GetDraftChapterID(c context.Context, memberId int, dto *dto.CreateDraft) (string, error) {
	var chapterID string
	sql := "SELECT id FROM chapters WHERE book_id = $1 AND team_id = $2 AND member_id = $3;"
	err := r.db.QueryRow(c, sql, dto.BookID, dto.TeamId, memberId).Scan(&chapterID)
	return chapterID, err
}
