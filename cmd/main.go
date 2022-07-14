package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/doniacld/tiny-hen/cmd/handlers"
	"github.com/doniacld/tiny-hen/cmd/prommetric"
)

// init is called before main to initialize the prometheus registries
func init() {
	prometheus.MustRegister(prommetric.Temp)
	prometheus.MustRegister(prommetric.Hum)
}

func main() {
	fmt.Println("Listening on port 10010")
	mux := http.NewServeMux()

	// GET /hi
	mux.HandleFunc("/hi", handlers.GetHi)

	// POST /measure
	mux.HandleFunc("/measure", handlers.PostMeasure)

	// GET /metrics
	mux.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":10010", mux)
	if err != nil {
		panic(err)
	}
}
