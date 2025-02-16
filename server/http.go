package server

import (
	"fmt"
	"log"
	"net/http"
	"symetrical-fishstick-go/main.go/routes"
)

func Http_Server() {
	fmt.Println("HTTP server on port 8080 is a Go!")
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(routes.Test))
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Fatal(err)
	}
}
