package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", GetRoot)
	mux.HandleFunc("/hello", GetHello)

	fmt.Println("Starting server on port: 8000.")

	error := http.ListenAndServe(":8000", mux)

	if errors.Is(error, http.ErrServerClosed) {

		fmt.Printf("Server closed.\n")

	} else if error != nil {

		fmt.Printf("Error starting server: %s\n", error)
		os.Exit(1)

	}
}
