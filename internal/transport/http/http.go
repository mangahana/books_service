package http

import (
	"books_service/internal/application"
	"books_service/internal/transport/http/controller"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HttpServer struct {
	server  *echo.Echo
	useCase application.UseCase
}

func New(useCase application.UseCase) *HttpServer {
	return &HttpServer{
		server:  echo.New(),
		useCase: useCase,
	}
}

func (h *HttpServer) ListenAndServe(socket string) {
	h.server.Start(socket)
}

func (h *HttpServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *HttpServer) Register() {
	h.server.Use(middleware.CORS())
	h.server.Use(middleware.BodyLimit("10M"))

	controller := controller.New(h.useCase)

	api := h.server.Group("/api/v1/books")

	api.GET("", controller.GetBooks)
	api.GET("/types", controller.GetTypes)
	api.GET("/persons", controller.GetPersons)
	api.GET("/genres", controller.GetGenres)
	api.GET("/statuses", controller.GetStatuses)
}
