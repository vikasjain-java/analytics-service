package main

import (
	"fmt"
	"net/http"
)

var visitorTracker = NewVisitorTracker()

func captureVisitorHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	visitorName := r.URL.Query().Get("visitorName")

	if url == "" || visitorName == "" {
		http.Error(w, "Missing url or visitorName", http.StatusBadRequest)
		return
	}

	visitorTracker.AddVisit(url, visitorName)
	w.WriteHeader(http.StatusNoContent)
}

func queryVisitorsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	url := r.URL.Query().Get("url")

	if url != "" {
		visitors := visitorTracker.GetUniqueVisitors(url)
		fmt.Fprintf(w, "<html><body><h1>Unique Visitors for URL</h1><ul>")

		fmt.Fprintf(w, "<li>%s: %d visitor(s)</li>", url, visitors)

		fmt.Fprintf(w, "</ul></body></html>")
	} else {
		allVisitors := visitorTracker.GetAllVisitors()
		fmt.Fprintf(w, "<html><body><h1>Unique Visitors per URL</h1><ul>")
		for url, count := range allVisitors {
			fmt.Fprintf(w, "<li>%s: %d visitor(s)</li>", url, count)
		}
		fmt.Fprintf(w, "</ul></body></html>")
	}
}
