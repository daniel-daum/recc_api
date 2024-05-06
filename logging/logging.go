package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func GetLogger() *Logger {

	var log zerolog.Logger

	if os.Getenv("RECC_ENV") == "PRODUCTION" {
		runLogFile, _ := os.OpenFile("recc.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
		output := zerolog.MultiLevelWriter(os.Stdout, runLogFile)

		log = zerolog.New(output).With().Timestamp().Logger()

	} else {
		output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
		log = zerolog.New(output).With().Timestamp().Logger()
	}

	logger := Logger{logger: log}

	return &logger
}

func (l Logger) Info(message string) {
	l.logger.Info().Str("message", message).Send()
}

func (l Logger) Debug(message string) {
	l.logger.Debug().Str("message", message).Send()
}

func (l Logger) Warn(message string) {
	l.logger.Warn().Str("message", message).Send()
}

func (l Logger) Error(message string, err error) {
	l.logger.Err(err).Str("message", message).Send()
}
