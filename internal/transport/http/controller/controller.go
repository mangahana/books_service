package controller

import (
	"books_service/internal/application"
	"books_service/internal/core/cerror"
	"books_service/internal/core/models"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type controller struct {
	validator *validator.Validate
	useCase   application.UseCase
}

func New(useCase application.UseCase) *controller {
	return &controller{
		validator: validator.New(),
		useCase:   useCase,
	}
}

func (h *controller) getUser(c echo.Context) (models.User, error) {
	output := models.User{}

	user, ok := c.Get("user").(models.User)
	if !ok {
		return output, cerror.New(cerror.UNAUTHORIZED, "you don't authorized")
	}

	return user, nil
}
