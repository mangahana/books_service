package controller

import (
	"books_service/internal/core/cerror"
	"books_service/internal/core/dto"
	"log"

	"github.com/labstack/echo/v4"
)

func (h *controller) UploadPage(c echo.Context) error {
	var dto dto.UploadPage
	if err := c.Bind(&dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "BAD_REQUEST"))
	}

	if err := h.validator.Struct(dto); err != nil {
		return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "invalid data"))
	}

	user, err := h.getUser(c)
	if err != nil {
		return c.JSON(401, cerror.New(cerror.UNAUTHORIZED, "UNAUTHORIZED"))
	}

	filename, err := h.useCase.UploadPage(c.Request().Context(), &user, &dto)
	if err != nil {
		log.Println(err)
		return c.JSON(400, err.Error())
	}

	return c.String(200, filename)
}
