#!/bin/bash

set -euo pipefail

REPO_ROOT=$(git rev-parse --show-toplevel)
MANIFESTS=$(kubectl kustomize "$REPO_ROOT/kubernetes/bootstrap" --enable-helm)
echo "$MANIFESTS" | yq 'select(.kind == "CustomResourceDefinition")' | kubectl apply -f -
kubectl wait --for=condition=Established --all customresourcedefinitions
echo "$MANIFESTS" | yq 'select(.kind != "CustomResourceDefinition")' | kubectl apply -f -
