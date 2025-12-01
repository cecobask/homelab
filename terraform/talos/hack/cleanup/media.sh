#!/bin/bash

set -euo pipefail

echo "Cleaning up media..."
kubens argocd
argocd login --core
argocd app delete media --yes --wait