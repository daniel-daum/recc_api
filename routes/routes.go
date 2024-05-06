package routes

import (
	"encoding/json"
	"net/http"
)

type healthResponse struct {
	Status string
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {

	healthResponse := healthResponse{Status: "Healthy"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(healthResponse)
}
