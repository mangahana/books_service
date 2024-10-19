package application

import (
	"books_service/internal/core/models"
	"context"
)

func (u *useCase) GetPersons(c context.Context) ([]models.Person, error) {
	return u.repo.GetPersons(c)
}
