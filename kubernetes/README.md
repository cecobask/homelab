# Bootstrap

## Cilium

```
kubectl kustomize bootstrap/cilium --enable-helm | kubectl apply -f -
```

## Proxmox CSI Plugin

```
kubectl kustomize bootstrap/csi-proxmox --enable-helm | kubectl apply -f -
```

## Argo CD

```
kubectl kustomize bootstrap/argocd --enable-helm | kubectl apply -f -
```

## Applications

```
kubectl apply -f bootstrap/platform.yaml
```
