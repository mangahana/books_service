package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetGenres(c context.Context) ([]models.Genre, error) {
	return u.repo.GetGenres(c)
}
