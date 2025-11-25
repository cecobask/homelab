# Cilium

## Bootstrap

The manifests are generated manually and applied automatically during Talos bootstrap.
Argo CD takes over the Cilium resources after the Talos cluster has been bootstrapped.

```
kubectl kustomize --enable-helm > manifests.yaml
```
