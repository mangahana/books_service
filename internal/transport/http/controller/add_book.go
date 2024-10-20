package controller

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) AddBook(c echo.Context) error {
	var dto dto.AddBook

	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "Invalid data"))
	}

	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, err.Error())
	}

	if err := h.useCase.Add(c.Request().Context(), user, dto); err != nil {
		return c.JSON(400, err.Error())
	}

	return c.String(200, "OK")
}
