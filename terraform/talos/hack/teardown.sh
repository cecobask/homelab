#!/bin/bash

set -euo pipefail

kubens argocd
argocd login argocd.cecobask.com --core
argocd proj windows add default --kind=deny --schedule="* * * * *" --duration=30m --applications="*" --manual-sync
argocd appset delete media --yes
argocd appset delete platform --yes
