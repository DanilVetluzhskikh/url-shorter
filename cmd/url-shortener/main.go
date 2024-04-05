package main

import (
	"learn/internal/config"
	sl "learn/internal/lib/logger"
	"learn/internal/logger"
	"learn/internal/storage/sqlite"
	"log/slog"
	"os"
)

func main() {
	cfg := config.MustLoad()
	logger := logger.SetupLogger(cfg.Env)

	logger.Info("url-shortener started", slog.String("env", cfg.Env))
	logger.Info("debug has started")

	storage, err := sqlite.New(cfg.StoragePath)

	if err != nil {
		logger.Error("failed init storage", sl.Err(err))
		os.Exit(1)
	}

	if err != nil {
		logger.Error("failed init storage", sl.Err(err))
		os.Exit(1)
	}

	_ = storage
}
