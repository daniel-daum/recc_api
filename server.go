package main

import (
	"fmt"
	"net/http"
)

func StartServer() {

	router := http.NewServeMux()

	router.HandleFunc("GET /", HealthCheck)

	server := &http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}
