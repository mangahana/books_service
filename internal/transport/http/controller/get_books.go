package controller

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetBooks(c echo.Context) error {
	books, err := h.useCase.GetBooks(c.Request().Context())
	if err != nil {
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, "unknow error"))
	}

	return c.JSON(200, books)
}
