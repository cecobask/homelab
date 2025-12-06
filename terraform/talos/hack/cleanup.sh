#!/bin/bash

set -euo pipefail

echo "Cleaning up applications..."
kubens argocd
argocd login --core
argocd proj windows add default --kind=deny --schedule="* * * * *" --duration=30m --applications="*" --manual-sync
argocd app delete --yes --selector='argocd.argoproj.io/instance, argocd.argoproj.io/instance notin (argocd,cilium,csi-proxmox,csi-smb)'
argocd app delete media --yes
kubectl wait pv --all --for=delete --timeout=3m
