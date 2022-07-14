# TinyGo: Getting the upper hen

This project holds all the needed code to build a monitoring system for a henhouse.

## Quickstart

Prerequisites:
- Install go, tinygo, kubectl, helm, kind

First deploy the cluster, it will take a few minutes, go make boil the water during this time.
```bash
make deploy
```

You can test even you do not have a arduino board with sending a measure:
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

If you have an Arduino board. You can flash your card using this piece of code. 

```bash
tinygo flash  -target=arduino-nano33 wifinina/httppost/main.go
```

Then read from the serial port, to find it you can use the following command and using tab to autocomplete:
```bash
```bash
ll /dev/cu.usbmodem1421
```

```bash
 go run utils/read_serial.go -port /dev/cu.usbmodem1101 
```

[//]: # ( TODO add a section about the SSID parameters in the tinygo file)

Now it should send data.

