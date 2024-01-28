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
	config := SetupConfig()
	mux := http.NewServeMux()

	log.Info().Str("ENV", config.environment).Str("PORT", config.port).Msg("Starting listener.")

	// Register handlers
	mux.HandleFunc("/", GetRoot)

	error := http.ListenAndServe(":"+config.port, mux)

	if error != nil {

		log.Error().Err(error).Msg("Error starting server.")
		os.Exit(1)

	}

}
