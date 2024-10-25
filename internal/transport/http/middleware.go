package http

import (
	"books_service/internal/core/cerror"

	"github.com/labstack/echo/v4"
)

func (h *httpServer) authenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")
		if token == "" {
			return c.JSON(400, cerror.New(cerror.BAD_REQUEST, "token is empty"))
		}

		user, err := h.authService.GetUser(c.Request().Context(), token)
		if err != nil {
			return c.JSON(403, cerror.New(cerror.FORBIDDEN, "invalid token"))
		}

		c.Set("user", user)

		return next(c)
	}
}
