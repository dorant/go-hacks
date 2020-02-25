# zipkin-client

## Build
make container
make push

## Install in K8s
kubectl create -f job.yaml

## Get logs
kubectl logs $(kubectl get pods --selector=job-name=zipkin-client-job -o jsonpath='{.items[*].metadata.name}')

## Remove finished job
kubectl delete -f job.yaml
