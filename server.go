package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func StartServer(settings *Settings) {
	router := http.NewServeMux()

	router.HandleFunc("GET /health", HealthCheck)
	router.HandleFunc("GET /reference", Reference)

	finalRouter := LoggingMiddleware(router)

	server := &http.Server{
		Addr:    ":" + settings.port,
		Handler: finalRouter,
	}

	slog.Info("Starting server", slog.String("port", settings.port), slog.String("env", settings.env))
	err := server.ListenAndServe()

	if err != nil {
		slog.Error("Error starting server", slog.Any("error", err))
	}
}

type Settings struct {
	env  string
	port string
}

func GetSettings() *Settings {
	slog.Info("Reading environment variables")

	if os.Getenv("RECC_ENV") != "PRODUCTION" {
		err := godotenv.Load()
		if err != nil {
			slog.Error("Error loading .env file", slog.Any("error", err))
		}
	}
	env := os.Getenv("RECC_ENV")
	port := os.Getenv("RECC_PORT")

	if env == "" {
		env = "DEVELOPMENT"
		slog.Warn("RECC_ENV env var is empty, using default", slog.String("env", env))
	}
	if port == "" {
		port = "8000"
		slog.Warn("RECC_PORT env var is empty, using default", slog.String("port", port))
	}

	settings := Settings{
		env:  env,
		port: port,
	}

	return &settings
}
