#!/bin/bash

set -euo pipefail

echo "Updating Argo CD credentials..."
kubens argocd
TEMP_ARGOCD_ADMIN_PASSWORD=$(kubectl get secret argocd-initial-admin-secret --output=json | jq -r '.data.password' | base64 -d)
argocd login argocd.cecobask.com --grpc-web --insecure --password="$TEMP_ARGOCD_ADMIN_PASSWORD" --username=admin
argocd account update-password --account=admin --current-password="$TEMP_ARGOCD_ADMIN_PASSWORD" --grpc-web --new-password="$ARGOCD_ADMIN_PASSWORD"
kubectl delete secret argocd-initial-admin-secret --ignore-not-found
READONLY_TOKEN=$(argocd account generate-token --account=readonly)
kubectl create secret generic argocd-readonly --from-literal=token="$READONLY_TOKEN" --namespace=media