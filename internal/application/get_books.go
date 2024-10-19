package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetBooks(c context.Context) ([]models.Book, error) {
	return u.repo.GetBooks(c)
}
