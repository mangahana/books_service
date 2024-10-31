package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetPersonsByName(c context.Context, name string) ([]models.Person, error) {
	return u.repo.GetPersonsByName(c, name)
}
