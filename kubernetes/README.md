# Bootstrap

## Cilium

```
kubectl kustomize platform/cilium --enable-helm | kubectl apply -f -
```

## Argo CD

```
kubectl kustomize bootstrap/argocd --enable-helm | kubectl apply -f -
```

## Applications

```
kubectl apply -f bootstrap/platform.yaml
```
