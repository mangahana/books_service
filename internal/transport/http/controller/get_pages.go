package controller

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *controller) GetPages(c echo.Context) error {
	chapterID := c.QueryParam("chapter_id")
	if chapterID == "" {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "empty chapter_id"))
	}

	pages, err := h.useCase.GetPages(c.Request().Context(), chapterID)
	if err != nil {
		return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, err.Error()))
	}

	return c.JSON(200, pages)
}
