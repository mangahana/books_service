package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetOneByLink(c context.Context, bookLink string) (models.OneBook, error) {
	return u.repo.GetOneByLink(c, bookLink)
}
