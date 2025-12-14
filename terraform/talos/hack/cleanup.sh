#!/bin/bash

set -euo pipefail

echo "Cleaning up applications..."
kubens argocd
argocd login --core
argocd proj windows add default --kind=deny --schedule="* * * * *" --applications="*" --duration=3m
argocd app delete --yes --selector='argocd.argoproj.io/instance, argocd.argoproj.io/instance notin (argocd,cilium,longhorn,sealed-secrets)'
argocd app delete media --yes
kubectl wait pv --all --for=delete --timeout=3m
kubectl patch setting deleting-confirmation-flag --namespace=longhorn-system --patch='{"value": "true"}' --type=merge
kubectl apply --namespace=longhorn-system -f https://raw.githubusercontent.com/longhorn/longhorn/v1.10.1/uninstall/uninstall.yaml
kubectl wait job longhorn-uninstall --namespace=longhorn-system --for=condition=complete --timeout=3m
