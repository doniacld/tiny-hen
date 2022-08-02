# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Source parameters
SOURCE_ENTRYPOINT=./cmd/main.go

# Binary parameters
BINARY_NAME=tinyhen-server
BINARY_DESTINATION=./bin
BINARY_PATH=$(BINARY_DESTINATION)/$(BINARY_NAME)

# Docker parameters
DOCKERCMD=docker
DOCKERBBUILD=$(DOCKERCMD) build
DOCKERRUN=$(DOCKERCMD) run
DOCKERSTOP=$(DOCKERCMD) stop
DOCKERRM=$(DOCKERCMD) rm

# Development
tidy:
		$(GOMOD) tidy
build:
		$(GOBUILD) -o $(BINARY_PATH) -v $(SOURCE_ENTRYPOINT)
run:
		$(GORUN) $(SOURCE_ENTRYPOINT)
unit_test:
		$(GOTEST) -v ./... -coverprofile=coverage.txt -covermode=atomic
clean:
		$(GOCLEAN) $(SOURCE_ENTRYPOINT)
		rm -f $(BINARY_PATH)

# Docker for local operations
build-linux:
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_PATH)-amd64 -v $(SOURCE_ENTRYPOINT)
docker_build: build-linux
		$(DOCKERBBUILD) -t tinyhen-server .
docker_run: build-linux docker_build
		$(DOCKERRUN) -d -p 10010:10010 tinyhen-server
docker_stop: docker_run
		$(DOCKERSTOP) tinyhen-server
docker_rm: docker_run
		$(DOCKERRM) tinyhen-server
docker_load:
	kind load docker-image tinyhen-server:latest --name tinyhen

# Deployment
.PHONY: deploy
deploy:
	./deploy/deploy.sh

grafana:
	# Create a port-forward to the grafana server (http://localhost:3000)
	kubectl port-forward service/kube-prometheus-stack-grafana 3000:80 -n monitoring

prometheus:
	# Create a port-forward to the prometheus server (http://localhost:9090)
	kubectl port-forward service/kube-prometheus-stack-prometheus 9090:9090 -n monitoring

destroy:
	kind delete cluster --name=tinyhen

# Curls
curl_measure:
	curl -v -X POST -H "Content-Type: application/json" http://localhost:1919/measure  -d '{"temperature": 31.2, "humidity": 41.6}'
