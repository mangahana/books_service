package controller

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetParams(c echo.Context) error {
	params, err := h.useCase.GetParams(c.Request().Context())
	if err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, err.Error()))
	}

	return c.JSON(200, params)
}
