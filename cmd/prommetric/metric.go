package prommetric

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type PromMeasure struct {
	Temp int `json:"temp"`
	Hum  int `json:"hum"`
}

var (
	Temp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celsius",
		Help: "The temperature of the hen house in celsius.",
	})

	Hum = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "humidity_percent",
		Help: "The humidity of the hen house in percent.",
	})
)

func (m PromMeasure) SetTempAndHum() {
	Temp.Set(float64(m.Temp / 10))
	Hum.Set(float64(m.Hum / 10))

	fmt.Printf("%s: SetTempAndHum set values: Temperature: %d °C, Humidity: %d %%\n",
		time.Now().String(), m.Temp/10, m.Hum/10)
}
