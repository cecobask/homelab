#!/bin/bash

set -euo pipefail

echo "Cleaning up applications..."
kubens argocd
argocd login --core
argocd proj windows add default --kind=deny --schedule="* * * * *" --duration=30m --applications="*" --manual-sync
argocd app delete media grafana prometheus --yes --wait