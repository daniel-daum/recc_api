package main

import (
	"log/slog"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			slog.Info(
				"request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.RequestURI()),
				slog.String("agent", r.UserAgent()),
				slog.Duration("elapsed", time.Since(start)),
			)
		}()
		next.ServeHTTP(w, r)
	})
}
