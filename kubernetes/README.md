# Bootstrap

## Cilium

```
kubectl kustomize bootstrap/cilium --enable-helm | kubectl apply -f -
```

## Proxmox Container Storage Interface Plugin

```
kubectl kustomize bootstrap/csi-proxmox --enable-helm | kubectl apply -f -
kubectl apply -f bootstrap/csi-proxmox/secret.yaml
```

## Argo CD

```
kubectl kustomize bootstrap/argocd --enable-helm | kubectl apply -f -
```

## Applications

```
kubectl apply -f bootstrap/platform.yaml
```
