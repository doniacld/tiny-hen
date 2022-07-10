package handlers

import (
	"fmt"
	"net/http"
)

func GetHi(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET /hi")

	// TODO return a json
	_, _ = w.Write([]byte("Hello my little hen"))
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
}
