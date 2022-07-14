package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetHi is the handler for GET /hi endpoint
func GetHi(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /hi")

	// implicit OK status
	// w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["greeting"] = "Hello tiny hen!"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Error happened while writing response. Err: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
