package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	log.Println("Server started on localhost:2112")
	if err != nil {
		log.Fatal("Server failed to start")
		return
	}
}
