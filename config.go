package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type configuration struct {
	environment string
	port        string
	database    string
}

func SetupConfig() *configuration {

	error := godotenv.Load()

	if error != nil {
		log.Error().Err(error).Msg("Error loading env file.")
		os.Exit(1)
	}

	env := os.Getenv("MALICIOUS_PICKLE_ENV")
	port := os.Getenv("MALICIOUS_PICKLE_PORT")
	database := os.Getenv("MALICIOUS_PICKLE_DB")

	if env == "" {
		env = "development"
	}

	config := configuration{environment: env, port: port, database: database}

	log.Info().Msg("Environment loaded successfully.")

	return &config
}
