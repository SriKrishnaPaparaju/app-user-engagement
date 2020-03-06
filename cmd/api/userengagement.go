package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/app_usage_metrics/v1/applications", metrics.createTrackID())
	http.HandleFunc("/api/app_usage_metrics/v1/metrics", metrics.createTrackID())
	http.HandleFunc("/api/app_usage_metrics/v1/metrics/publish", metrics.createTrackID())
	address := ":8080"
	log.Println("Starting server on address", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		panic(err)
	}
}
