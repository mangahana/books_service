package repository

import (
	"books_service/internal/core/dto"
	"context"
)

func (r *repo) CreateDraft(c context.Context, memberId int, teamName string, dto *dto.CreateDraft) (string, error) {
	var id string
	sql := "INSERT INTO chapters (book_id, team_id, team_name, member_id) VALUES ($1, $2, $3, $4) RETURNING id;"
	err := r.db.QueryRow(c, sql, dto.BookID, dto.TeamId, teamName, memberId).Scan(&id)
	return id, err
}
