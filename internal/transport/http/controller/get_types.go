package controller

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetTypes(c echo.Context) error {
	types, err := h.useCase.GetTypes(c.Request().Context())
	if err != nil {
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, err.Error()))
	}

	return c.JSON(200, types)
}
