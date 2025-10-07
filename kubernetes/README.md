# Bootstrap

## Cilium

```
kubectl kustomize bootstrap/cilium --enable-helm | kubectl apply -f -
rm -rf bootstrap/cilium/charts
```

## Argo CD

```
kubectl kustomize bootstrap/argocd --enable-helm | kubectl apply -f -
rm -rf bootstrap/argocd/charts
```

## App of Apps

```
kubectl apply -f bootstrap/application.yaml
```
