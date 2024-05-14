package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

// @title			Recc API
// @version		0.1.0
// @description	The Recc API documentation.
// @contact.name	Daniel Daum
// @host		localhost:8000
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
