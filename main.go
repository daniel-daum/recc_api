package main

import (
	"net/http"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Info().Msg("Starting server.")

	config := SetupConfig()
	mux := http.NewServeMux()

	// Register handlers
	mux.HandleFunc("/", GetRoot)

	log.Info().Str("Env", config.environment).Str("Port", config.port).Msg("Starting started successfully.")
	error := http.ListenAndServe(":"+config.port, mux)

	if error != nil {

		log.Error().Err(error).Msg("Error starting server.")
		os.Exit(1)

	}

}
