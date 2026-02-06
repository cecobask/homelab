#!/bin/bash

set -euo pipefail

kubens argocd
argocd login --core
argocd proj windows add default --kind=deny --schedule="* * * * *" --applications="*" --manual-sync --duration=3m
