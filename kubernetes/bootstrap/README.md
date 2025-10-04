# Bootstrap

## Argo CD

```
kubectl kustomize argocd --enable-helm | kubectl apply -f -
rm -rf argocd/charts
kubectl get secret argocd-initial-admin-secret -n argocd -o json | jq -r '.data.password | @base64d'
```

## Main Application

```
kubectl apply -f app.yaml
```
