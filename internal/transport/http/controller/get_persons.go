package controller

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetPersons(c echo.Context) error {
	persons, err := h.useCase.GetPersons(c.Request().Context())
	if err != nil {
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, "unknow error"))
	}

	return c.JSON(200, persons)
}
