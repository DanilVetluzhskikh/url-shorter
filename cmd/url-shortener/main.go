package main

import (
	"learn/internal/config"
	redirectUrl "learn/internal/http-server/handlers/url/redirect"
	saveUrl "learn/internal/http-server/handlers/url/save"
	mwLogger "learn/internal/http-server/middleware"
	sl "learn/internal/lib/logger"
	"learn/internal/logger"
	"learn/internal/storage/sqlite"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(logger))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", saveUrl.New(logger, storage))
	router.Get("/{alias}", redirectUrl.New(logger, storage))

	server := &http.Server{
		Addr:         cfg.HTTPServer.Adress,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server ebnulsya")
	}
}
