#!/bin/bash

set -euo pipefail

echo "Updating Kubernetes manifests..."
REPO_ROOT="$(git rev-parse --show-toplevel)"
APP_DIRS=(
  "kubernetes/platform/argocd"
  "kubernetes/platform/cert-manager"
  "kubernetes/platform/cilium"
  "kubernetes/platform/csi-proxmox"
  "kubernetes/platform/envoy-gateway"
  "kubernetes/platform/external-secrets"
  "kubernetes/platform/prometheus"
  "kubernetes/platform/snapshot-controller"
)

for DIR in "${APP_DIRS[@]}"; do
  APP_DIR="$REPO_ROOT/$DIR"
  MANIFESTS_FILE="$APP_DIR/manifests.yaml"
  kubectl kustomize "$APP_DIR" --enable-helm > "$MANIFESTS_FILE"
  echo "Updated $MANIFESTS_FILE"
done
