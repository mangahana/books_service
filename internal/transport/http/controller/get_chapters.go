package controller

import (
	"books_service/internal/core/cerror"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetChapters(c echo.Context) error {
	bookIdParam := c.QueryParam("bookId")
	bookId, err := strconv.Atoi(bookIdParam)
	if err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, err.Error()))
	}

	endSorting := false
	if c.QueryParam("endSorting") == "1" {
		endSorting = true
	}

	chapters, err := h.useCase.GetChapters(c.Request().Context(), bookId, endSorting)
	if err != nil {
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, err.Error()))
	}

	return c.JSON(200, chapters)
}
