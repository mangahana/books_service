package controller

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetStatuses(c echo.Context) error {
	statuses, err := h.useCase.GetStatuses(c.Request().Context())
	if err != nil {
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, "unknow error"))
	}

	return c.JSON(200, statuses)
}
