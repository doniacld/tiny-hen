# Tiny-hen app

Tiny-hen app is an HTTP webserver, written in Go, that serves the following endpoints:

## Hi

Hi is an endpoint such as ping command to make sure the webserver is running properly, and 
you can access it from where you call it.

[GET /hi](http://localhost/hi)

The answer is a basic greeting 
```json
{"greet": "Hello little hen"}
```

## Measure

Measure is a HTTP endpoint where you can post your measure of temperature and humidity.

[POST /measure](http://localhost/measure) with the following payload:

```json
{
  "temp": 23.5,
  "hum": 31.5
}
```

## Metrics

Metrics is a HTTP endpoint where you can get the latest metrics scrapped by Prometheus.

[GET /metrics](http://localhost/metrics)

