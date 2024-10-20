package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetChapters(c context.Context, bookId int, endSorting bool) ([]models.Chapter, error) {
	return u.repo.GetChapters(c, bookId, endSorting)
}
