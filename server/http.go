package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func Http_Server() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("HTTP_PORT")

	fmt.Println("HTTP server on port 8080 is a Go!")
	mux := http.NewServeMux()
	// mux.Handle("/", http.HandlerFunc(routes.TestHandler))
	srvErr := http.ListenAndServe(port, mux)

	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
