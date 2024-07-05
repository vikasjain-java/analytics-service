package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/capture", captureVisitorHandler)
	http.HandleFunc("/query", queryVisitorsHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
