package http

import (
	"books_service/internal/application"
	"books_service/internal/infrastructure"
	"books_service/internal/transport/http/controller"
	"context"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type httpServer struct {
	server  *echo.Echo
	useCase application.UseCase

	authService infrastructure.AuthService
}

func New(useCase application.UseCase, authService infrastructure.AuthService) *httpServer {
	return &httpServer{
		server:      echo.New(),
		useCase:     useCase,
		authService: authService,
	}
}

func (h *httpServer) ListenAndServe(socket string) {
	h.server.Start(socket)
}

func (h *httpServer) Shutdown(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

func (h *httpServer) Register() {
	h.server.Use(middleware.CORS())
	h.server.Use(middleware.BodyLimit("10M"))

	controller := controller.New(h.useCase)

	//
	api := h.server.Group("/api/v1/books")

	api.GET("", controller.GetBooks)
	api.GET("/chapters", controller.GetChapters)
	api.GET("/pages", controller.GetPages)
	api.GET("/types", controller.GetTypes)
	api.GET("/persons", controller.GetPersons)
	api.GET("/genres", controller.GetGenres)
	api.GET("/statuses", controller.GetStatuses)

	//
	private := api.Group("", h.authenticateMiddleware)

	private.POST("", controller.AddBook)
	private.POST("/add_draft", controller.CreateDraft)
}
