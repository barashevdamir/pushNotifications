package main

import (
	"log/slog"
	"os"
	"pushNotifications/internal/config"
	"pushNotifications/internal/storage/postgresql"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	db, err := postgresql.Connect(cfg.HTTPServer.User, cfg.HTTPServer.Password, cfg.HTTPServer.DbName, cfg.HTTPServer.Host, cfg.HTTPServer.Port)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	log.Info("Connected to PostgreSQL database")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
