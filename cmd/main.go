package main

import (
	"log/slog"
	"os"

	"github.com/moxicom/vk-internship-2024-spring/internal/handlers"
	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/service"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	// init logger: slog
	log := setupLogger(envLocal)
	log.Info("starting server...")
	log.Debug("debug mode enabled")

	// init storage: postgres
	db, err := postgres.New(models.PostgresConfig{
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		PostgresUser:     "postgres",
		PostgresPassword: "postgres",
		PostgresName:     "mydatabase",
		PostgresSSLMode:  "disable",
	})
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	storage := storage.New(db)

	// init service
	service := service.New(storage)

	// init router and run server
	log.Info("initializing routes...")
	if err := handlers.Run(log, service); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
