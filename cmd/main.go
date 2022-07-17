package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/doniacld/tiny-hen/cmd/handlers"
	"github.com/doniacld/tiny-hen/cmd/prommetric"
)

func main() {
	// register the prometheus metrics
	prommetric.RegisterGauges()

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
