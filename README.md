# simserver

## How to use

### Deploy

Notes: We assume that you test this product w/ Docker Desktop and Kubernetes
You might need to comment in `imagePullPolicy: Never` if you deploy this image in the EKS or whatever.

```
$ kubectl apply -f application.yaml
```

### Expose

```
$ kubectl apply -f myservice.yaml
```
