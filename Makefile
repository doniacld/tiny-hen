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

# Docker
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
prom:
	prometheus --config.file=deploy/prometheus.yaml

# Deployment
deploy: docker_build
	echo "---------------- Create cluster tinyhen ----------------"
	# Create cluster with the right configuration
	kind create cluster --name tinyhen --config deploy/cluster-config.yaml


	echo "---------------- Deploying configmap dashboard customization ----------------"
	kubectl create ns monitoring
	# Customize the cluster with the configmap
	kubectl apply -k deploy/monitoring

	# Install prometheus stack (Grafana, Prometheus operator, etc)
	echo "---------------- Deploying kube prometheus stack with dashboard provider ----------------"
	helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --values deploy/monitoring/values.yaml --namespace monitoring

	# Deploy the prometheus service monitoring for the app
	echo "---------------- Deploying service monitor for tinyhen app ----------------"
	kubectl apply -f deploy/monitoring/service_monitor.yaml

	# Load the docker image from the tinyhen app to the cluster
	echo "---------------- Loading docker image from tinyhen app to the cluster ----------------"
	kind load docker-image tinyhen-server:latest --name tinyhen

	# Deploy tinyhen app
	echo "---------------- Deploying tinyhen app ----------------"
	kubectl apply -f deploy/app.yaml

	# Deploy the ingress to expose the app
	echo "---------------- Deploying ingress ----------------"
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

	# Wait for the ingress to be ready
	echo "---------------- Waiting for the ingress to be ready ----------------"
	kubectl wait --namespace ingress-nginx \
	  --for=condition=ready pod \
	  --selector=app.kubernetes.io/component=controller \
	  --timeout=800s

docker_load:
	kind load docker-image tinyhen-server:latest --name tinyhen

grafana:
	# Create a port-forward to the grafana server (http://localhost:3000)
	kubectl port-forward service/kube-prometheus-stack-grafana 3000:80 -n monitoring

destroy:
	kind delete cluster --name=tinyhen