package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/doniacld/tiny-hen/cmd/prommetric"
)

// PostMeasureResponse is the response from PostMeasure endpoint
type PostMeasureResponse struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

// PostMeasure is the handler for POST /measure endpoint
func PostMeasure(w http.ResponseWriter, r *http.Request) {
	log.Println("POST /measure")
	var measureResponse PostMeasureResponse

	err := json.NewDecoder(r.Body).Decode(&measureResponse)
	if err != nil {
		log.Printf("Error happened in JSON unmarshal. Err: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// send the data to prometheus
	promMeasure := prommetric.PromMeasure{
		Temperature: measureResponse.Temperature,
		Humidity:    measureResponse.Humidity,
	}
	promMeasure.SetTempAndHum()

	w.WriteHeader(http.StatusCreated)
}
