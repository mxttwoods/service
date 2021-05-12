package main

import (
	"log"
	"net/http"
	"os"
	"service/homepage"
	"service/server"
)

// env variables
var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)

	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.New(mux, ServiceAddr)

	logger.Println("server starting")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}
