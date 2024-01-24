package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

func GetRoot(response http.ResponseWriter, request *http.Request) {

	log.Info().Str("Method", request.Method).Str("Path", request.URL.Path).Msg("Request")

	io.WriteString(response, "This is my api!\n")
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
