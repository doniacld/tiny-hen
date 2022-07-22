package prommetric

import (
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type PromMeasure struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

var (
	TempGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celsius",
		Help: "The temperature of the hen house in celsius.",
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

	fmt.Printf("%s: SetTempAndHum set values: Temperature: %f °C, Humidity: %f %%\n",
		time.Now(), m.Temperature, m.Humidity)
}
