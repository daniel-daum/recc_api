package main

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	healthResponse := HealthResponse{Status: "Healthy"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(healthResponse)
}
