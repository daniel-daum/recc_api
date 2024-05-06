package recc

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/daniel-daum/recc_api/logging"
	"github.com/daniel-daum/recc_api/routes"
	"github.com/joho/godotenv"
)

type settings struct {
	env  string
	port string
}

func Settings() *settings {
	log := logging.GetLogger()

	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file.", err)
	}

	env := os.Getenv("RECC_ENV")
	port := os.Getenv("RECC_PORT")

	if env == "" {
		env = "DEVELOPMENT"
	}
	if port == "" {
		port = "8000"
	}

	config := settings{
		env:  env,
		port: port,
	}

	return &config
}

func StartServer() {
	settings := Settings()
	log := logging.GetLogger()

	router := http.NewServeMux()
	router.HandleFunc("GET /", routes.HealthCheck)

	addr := ":" + settings.port
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGINT, syscall.SIGTERM)

	startPort := fmt.Sprintf("starting listener on port -> %v", settings.port)
	startEnv := fmt.Sprintf("server environment mode -> %v", settings.env)

	log.Info(startPort)
	log.Info(startEnv)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("error starting server.", err)
		}
	}()

	<-gracefulShutdown
	log.Info("recieved shutdown signal.")
	log.Info("shutdown complete.")
}
