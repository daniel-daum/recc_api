package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	healthResponse := HealthResponse{Status: "healthy"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(healthResponse)

	if err != nil {
		slog.Error(
			"Error marshalling json response",
			slog.Any("error", err),
			slog.String("method", r.Method),
			slog.String("url", r.URL.RequestURI()),
			slog.String("agent", r.UserAgent()),
		)
	}
}
