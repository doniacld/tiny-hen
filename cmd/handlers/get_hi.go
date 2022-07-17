package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// GetHi is the handler for GET /hi endpoint
func GetHi(w http.ResponseWriter, r *http.Request) {
	log.Println("GET /hi")

	// default status code is 200
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["greeting"] = "Cluck!"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("Error happened in JSON marshal. Err: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(jsonResp)
	if err != nil {
		log.Printf("Error happened while writing response. Err: %s", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
