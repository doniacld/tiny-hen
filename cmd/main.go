package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/doniacld/tiny-hen/cmd/handlers"
	"github.com/doniacld/tiny-hen/cmd/prommetric"
)

// TODO Move to README.md
//  curl -v -X POST -H "Content-Type: application/json" http://localhost:10010/measure  -d '{"temp": 270, "hum": 300}'
//  curl -v -X GET -H "Content-Type: application/json" http://localhost:10010/hi
//  curl -v -X GET -H "Content-Type: application/json" http://localhost:10010/metrics

func init() {
	prometheus.MustRegister(prommetric.Temp)
	prometheus.MustRegister(prommetric.Hum)
}

func main() {
	fmt.Println("Listening on port 10010")
	mux := http.NewServeMux()

	// POST /measure
	mux.HandleFunc("/measure", handlers.PostMeasure)

	// GET /hi
	mux.HandleFunc("/hi", handlers.GetHi)

	// GET /metrics
	mux.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":10010", mux)
	if err != nil {
		panic(err)
	}
}
