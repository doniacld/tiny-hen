# TinyGo: Getting the upper hen

This project holds all the needed code to build a monitoring system for a henhouse.

## Quickstart

Prerequisites:
- Install go, tinygo, kubectl, helm, kind

First deploy the cluster, it will take a few minutes, go get the water boiling during this time.
```bash
make deploy
```

You can test (even you do not have an Arduino board) by sending a measure:
```bash
 make curl_measure
```

You should expect a 200 HTTP answer:
```bash
< HTTP/1.1 200 OK
< Date: Thu, 14 Jul 2022 17:55:02 GMT
< Content-Length: 0
< Connection: keep-alive
```

## Hardware setup

The needed material are 
- an Arduino Nano 33 IoT or an equivalent board with a WiFi antenna
- a DHT sensor to measure temperature and humidity

You can then flash the [tinygo program](https://github.com/doniacld/tinygo-discovery/blob/main/tiny-hen/main.go) on your board using this command. 

```bash
tinygo flash -target=arduino-nano33 <yourpath>/tiny-hen/
```

Then read from the serial port. To find it, you can use the following command, using tab to autocomplete the usb modem number:

```bash
ll /dev/cu.usbmodem1421
```

```bash
 go run utils/read_serial.go -port /dev/cu.usbmodem1101 
```

[//]: # ( TODO add a section about the SSID parameters in the tinygo file)

Now it should send data.

