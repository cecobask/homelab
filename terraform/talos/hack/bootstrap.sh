#!/bin/bash

set -euo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
MANIFESTS=$(kubectl kustomize "$REPO_ROOT/kubernetes/bootstrap" --enable-helm)
echo "$MANIFESTS" | yq 'select(.kind == "CustomResourceDefinition")' | kubectl apply -f -
kubectl wait --for=condition=Established --all customresourcedefinitions
kubectl kustomize "$REPO_ROOT/kubernetes/platform/cilium" --enable-helm | kubectl apply -f -
until kubectl get customresourcedefinition ciliumloadbalancerippools.cilium.io &>/dev/null; do echo "Waiting for Cilium CRDs..."; sleep 3; done
echo "$MANIFESTS" | yq 'select(.kind != "CustomResourceDefinition")' | kubectl apply -f -
