package main

import (
	"books_service/internal/application"
	"books_service/internal/core/configuration"
	"books_service/internal/infrastructure/authorization"
	"books_service/internal/infrastructure/repository"
	"books_service/internal/infrastructure/teams"
	"books_service/internal/transport/http"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := configuration.Load()
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repository.New(&cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	authService, err := authorization.New(cfg.Services.AuthServiceSocket)
	if err != nil {
		log.Fatal(err)
	}

	teamService, err := teams.New(cfg.Services.TeamsServiceSocket)
	if err != nil {
		log.Fatal(err)
	}

	useCase := application.New(repo, teamService)

	httpServer := http.New(useCase, authService)
	httpServer.Register()
	go httpServer.ListenAndServe(cfg.Server.HttpSocket)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), time.Second*5)
	defer shutdown()

	httpServer.Shutdown(ctx)
	repo.Close()
}
