#!/bin/bash

set -e

echo "
       ___  _              _ _
      |_ _|<_>._ _  _ _   | | | ___ ._ _
       | | | || ' || | |  | - |/ ._>| ' |
🐓     |_| |_||_|_|\_. |  |_|_|\___.|_|_|v     🐓
                   <___|                "

echo "---------------- 🏠 Create cluster tinyhen ----------------"
# Create cluster with the right configuration
if [[ "$1" == "sudo" ]] ; then
    sudo kind create cluster --name tinyhen --config deploy/cluster-config.yaml
else
    kind create cluster --name tinyhen --config deploy/cluster-config.yaml
fi

kube_config_file="/home/$(whoami)/.kube/kind-tinyhen"
sudo kind get kubeconfig --name tinyhen > "${kube_config_file}"
chmod 600 "${kube_config_file}"
export KUBECONFIG="${kube_config_file}"

echo "---------------- 📈 Deploying configmap dashboard customization ----------------"
kubectl create ns monitoring
# Customize the cluster with the configmap
kubectl apply -k deploy/monitoring

# Install prometheus stack (Grafana, Prometheus operator, etc)
echo "---------------- 🔎 Deploying kube prometheus stack with dashboard provider ----------------"
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --values deploy/monitoring/values.yaml --namespace monitoring

# Deploy the prometheus service monitoring for the app
echo "---------------- 🩺 Deploying service monitor for tinyhen app ----------------"
kubectl apply -f deploy/monitoring/service_monitor.yaml

# Deploy tinyhen app
echo "---------------- 🐓 Deploying tinyhen app ----------------"
kubectl create ns tiny-hen
kubectl apply -f deploy/app.yaml

# Deploy the ingress to expose the app
echo "---------------- 🌐 Deploying ingress ----------------"
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

# Wait for the ingress to be ready
echo "---------------- ⏳ Waiting for the ingress to be ready ----------------"
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=800s