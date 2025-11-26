#!/bin/bash

set -euo pipefail

kubens argocd
argocd login --core
argocd proj windows add default --kind=deny --schedule="* * * * *" --duration=30m --applications="*" --manual-sync
argocd app delete media --yes --wait
argocd app delete platform --yes --wait
