package controller

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) AddPerson(c echo.Context) error {
	var dto dto.AddPerson
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "invalid data"))
	}

	err := h.useCase.AddPerson(c.Request().Context(), &dto)
	if err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, err.Error()))
	}

	return c.JSON(200, "OK")
}
