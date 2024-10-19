package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetStatuses(c context.Context) ([]models.Status, error) {
	return u.repo.GetStatuses(c)
}
