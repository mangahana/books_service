package controller

import (
	"books_service/internal/core/cerror"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
)

func (h *controller) GetBooks(c echo.Context) error {
	{
		// by link
		link := c.QueryParam("link")
		if link != "" {
			books, err := h.useCase.GetOneByLink(c.Request().Context(), link)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					return c.JSON(404, cerror.New(cerror.NOT_FOUND, "the book not found"))
				}
				return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, err.Error()))
			}
			return c.JSON(200, books)
		}
	}

	{
		// get all
		books, err := h.useCase.GetBooks(c.Request().Context())
		if err != nil {
			return c.JSON(500, cerror.New(cerror.INTERNAL_SERVER_ERROR, err.Error()))
		}
		return c.JSON(200, books)
	}
}
