#!/bin/bash

set -euo pipefail

echo "Bootstrapping applications..."
MANIFESTS_DIR="$(dirname "$0")/manifests"
kubectl apply -k "$MANIFESTS_DIR"
