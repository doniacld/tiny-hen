package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/doniacld/tiny-hen/cmd/prommetric"
)

type PostMeasureResponse struct {
	Temp int `json:"temp"`
	Hum  int `json:"hum"`
}

func PostMeasure(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /measure")
	var measureResponse PostMeasureResponse

	err := json.NewDecoder(r.Body).Decode(&measureResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	promMeasure := prommetric.PromMeasure{
		Temp: measureResponse.Temp,
		Hum:  measureResponse.Hum,
	}
	promMeasure.SetTempAndHum()
}
