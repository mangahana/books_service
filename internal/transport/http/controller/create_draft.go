package controller

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"

	"github.com/labstack/echo/v4"
)

func (h *controller) CreateDraft(c echo.Context) error {
	var dto dto.CreateDraft
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, cerror.New(cerror.UNAUTHORIZED, "UNAUTHORIZED"))
	}

	chapterID, err := h.useCase.CreateDraft(c.Request().Context(), &user, &dto)
	if err != nil {
		return c.JSON(400, err.Error())
	}

	return c.JSON(200, map[string]string{"chapter_id": chapterID})
}
