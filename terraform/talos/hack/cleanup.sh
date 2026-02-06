#!/bin/bash

set -euo pipefail

echo "Cleaning up applications..."
argocd app delete --yes --selector='argocd.argoproj.io/instance, argocd.argoproj.io/instance notin (argocd,cilium,csi-proxmox,snapshot-controller)'
argocd app delete --yes media
kubectl wait pv --all --for=delete --timeout=3m
