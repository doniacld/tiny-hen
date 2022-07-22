package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/doniacld/tiny-hen/cmd/handlers"
	"github.com/doniacld/tiny-hen/cmd/prommetric"
)

const serverPort = ":10010"

func main() {
	// register the prometheus metrics
	prommetric.RegisterGauges()

	fmt.Printf("Listening on port%s\n", serverPort)
	mux := http.NewServeMux()

	// GET /hi
	mux.HandleFunc("/hi", handlers.GetHi)

	// POST /measure
	mux.HandleFunc("/measure", handlers.PostMeasure)

	// GET /metrics
	mux.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(serverPort, mux)
	if err != nil {
		panic(err)
	}
}
