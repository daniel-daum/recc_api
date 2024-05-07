package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {

	t.Run("returns health check", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		HealthCheck(response, request)

		expected := HealthResponse{Status: "Healthy"}
		var result HealthResponse

		err := json.NewDecoder(response.Body).Decode(&result)

		assert.Nil(t, err)
		assert.Equal(t, expected, result)
	})
}
