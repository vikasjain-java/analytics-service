package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/capture", captureVisitorHandler)
	http.HandleFunc("/query", queryVisitorsHandler)

	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Could not start server: %s\n", err)
	}
}
