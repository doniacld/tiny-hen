# Tiny-hen app

Tiny-hen app is an HTTP webserver, written in Go, that serves the following endpoints:

## Hi

Hi is an endpoint such as ping command to make sure the webserver is running properly, and 
you can access it from where you call it.

[GET /hi](http://localhost/hi)

The answer is a basic greeting 
```json
{"greeting": "Hello tiny hen!"}
```

> Calling the endpoint on cluster, remove the port.
```bash
curl -v -X GET http://localhost:10010/hi
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

```bash
 curl -v -X POST -H "Content-Type: application/json" http://localhost:10010/measure  -d '{"temp": 360, "hum": 480}' 
```

## Metrics

Metrics is an HTTP endpoint where you can get the latest metrics scrapped by Prometheus.

[GET /metrics](http://localhost/metrics)

```bash
 curl -v -X GET http://localhost:10010/metrics