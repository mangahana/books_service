package main

import (
	"books_service/internal/application"
	"books_service/internal/core/configuration"
	"books_service/internal/infrastructure/repository"
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

	useCase := application.New(repo)

	httpServer := http.New(useCase)
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
