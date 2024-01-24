package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type configuration struct {
	environment string
	port        string
}

func SetupConfig() *configuration {

	error := godotenv.Load()

	if error != nil {
		log.Error().Err(error).Msg("Error loading env file.")
		os.Exit(1)
	}

	env := os.Getenv("MALICIOUS_PICKLE_ENV")
	port := os.Getenv("PORT")

	if env == "" {
		env = "development"
	}

	config := configuration{environment: env, port: port}

	log.Info().Msg("Config loaded successfully.")

	return &config
}
