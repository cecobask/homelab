# Cilium

## Bootstrap

The base manifests are generated manually, then automatically applied during Talos bootstrap.
ArgoCD takes over the Cilium installation after the Talos cluster has been bootstrapped.

```
kubectl kustomize --enable-helm cilium/base > cilium/base/manifests.yaml
```
