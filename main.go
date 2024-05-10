package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func main() {
	// Global logging settings
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
		}),
	))

	settings := GetSettings()

	StartServer(settings)
}
