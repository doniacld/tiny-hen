# Deploy

## Create cluster
```bash
kind create cluster --name tinyhen --config deploy/cluster-config.yaml
```

Expected output:
```bash
Creating cluster "tinyhen-test" ...
 ✓ Ensuring node image (kindest/node:v1.24.0) 🖼 
 ✓ Preparing nodes 📦  
 ✓ Writing configuration 📜 
 ✓ Starting control-plane 🕹️ 
 ✓ Installing CNI 🔌 
 ✓ Installing StorageClass 💾 
Set kubectl context to "kind-tinyhen-test"
You can now use your cluster with:

kubectl cluster-info --context kind-tinyhen-test

Not sure what to do next? 😅  Check out https://kind.sigs.k8s.io/docs/user/quick-start/
```

## Create namespace monitoring
```bash
kind create cluster --name tinyhen --config deploy/cluster-config.yaml
```

EExpected output:
```bash
```

## Deploy dashboard as configmao
```bash
kubectl apply -k deploy/monitoring
```

## Deploy Kube Prometheus Stack

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install kube-prometheus-stack prometheus-community/kube-prometheus-stack --values deploy/monitoring/values.yaml --namespace monitoring
```

## Deploy service monitoring

```bash
kubectl apply -f deploy/monitoring/service_monitor.yaml
```

## Deploy app

```bash
kubectl apply -f deploy/app.yaml
```

## Deploy ingress Nginx

[Ingress Nginx documentation](https://kind.sigs.k8s.io/docs/user/ingress#ingress-nginx)

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
```

Expected output:
```bash
namespace/ingress-nginx created
serviceaccount/ingress-nginx created
serviceaccount/ingress-nginx-admission created
role.rbac.authorization.k8s.io/ingress-nginx created
role.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrole.rbac.authorization.k8s.io/ingress-nginx created
clusterrole.rbac.authorization.k8s.io/ingress-nginx-admission created
rolebinding.rbac.authorization.k8s.io/ingress-nginx created
rolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx created
clusterrolebinding.rbac.authorization.k8s.io/ingress-nginx-admission created
configmap/ingress-nginx-controller created
service/ingress-nginx-controller created
service/ingress-nginx-controller-admission created
deployment.apps/ingress-nginx-controller created
job.batch/ingress-nginx-admission-create created
job.batch/ingress-nginx-admission-patch created
ingressclass.networking.k8s.io/nginx created
validatingwebhookconfiguration.admissionregistration.k8s.io/ingress-nginx-admission created
```

## Deploy nginx-ingress-controller
```bash
kubectl wait --namespace ingress-nginx \
  --for=condition=ready pod \
  --selector=app.kubernetes.io/component=controller \
  --timeout=90s
```

> Wait until the command is finished

Expected output:
```bash
pod/ingress-nginx-controller-5458c46d7d-kc7rl condition met
```

Verify deployment:
```bash
kubectl get pod -n ingress-nginx                                                                                                                                                                                 21:07:49
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-82sdh        0/1     Completed   0          4m5s
ingress-nginx-admission-patch-jj7xq         0/1     Completed   2          4m5s
ingress-nginx-controller-5458c46d7d-xk6q8   1/1     Running     0          4m5s
```
