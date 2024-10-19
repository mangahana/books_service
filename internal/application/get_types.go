package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetTypes(c context.Context) ([]models.BookType, error) {
	return u.repo.GetTypes(c)
}
