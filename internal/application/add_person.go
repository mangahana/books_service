package application

import (
	"books_service/internal/core/dto"
	"context"
)

func (u *useCase) AddPerson(c context.Context, person *dto.AddPerson) error {
	return u.repo.AddPerson(c, person)
}
