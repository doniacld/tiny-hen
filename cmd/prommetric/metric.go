package prommetric

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type PromMeasure struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

var (
	TempGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celsius",
		Help: "The temperature of the hen house in Celsius.",
	})

	HumGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "humidity_percent",
		Help: "The humidity of the hen house in percent.",
	})
)

func RegisterGauges() {
	prometheus.MustRegister(TempGauge)
	prometheus.MustRegister(HumGauge)
}

func (m PromMeasure) SetTempAndHum() {
	TempGauge.Set(m.Temperature)
	HumGauge.Set(m.Humidity)

	log.Printf("Set gauges: Temperature: %v°C, Humidity: %v%%\n", m.Temperature, m.Humidity)
}
