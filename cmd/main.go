package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/moxicom/vk-internship-2024-spring/internal/models"
	"github.com/moxicom/vk-internship-2024-spring/internal/storage/postgres"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// TODO: init config: cleanenv

	// TODO: init logger: slog
	log := setupLogger(envLocal)
	log.Info("starting server...")
	log.Debug("debug mode enabled")

	// TODO: init storage: postgres
	storage, err := postgres.New(models.PostgresConfig{
		PostgresHost:     "localhost",
		PostgresPort:     "5432",
		PostgresUser:     "postgres",
		PostgresPassword: "postgres",
		PostgresName:     "postgres",
		PostgresSSLMode:  "disable",
	})
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println(storage)

	// TODO: init router

	// TODO: run server

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
