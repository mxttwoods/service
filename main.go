package main

import (
	"log"
	"net/http"
	"os"
	"service/index"
	"service/server"
)

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	logger := log.New(os.Stdout, " ", log.LstdFlags|log.Lshortfile)
	h := index.NewHandlers(logger)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)
	srv := server.New(mux, ServiceAddr)
	logger.Println("Server starting")
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		logger.Fatalf("Server failed to start: %v", err)
	}
}
