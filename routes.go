package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
)

type HealthResponse struct {
	Status string `json:"status" example:"healthy"`
}

// HealthCheck godoc
// @Summary      Health
// @Description  Returns "healthy" API status
// @Produce      json
// @Success  	 200  object  HealthResponse  "HealthResponse JSON"
// @Router       /health [get]
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

func Reference(w http.ResponseWriter, r *http.Request) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "./docs/swagger.json",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "Recc API",
		},
		DarkMode: true,
	})

	if err != nil {
		slog.Error("error", err)
	}

	fmt.Fprintln(w, htmlContent)
}
