#!/bin/bash
set -e
set -o pipefail
set -u

flux suspend kustomization flux-system
flux delete kustomization services --silent
flux delete kustomization media-apps --silent
flux delete kustomization media-core --silent
flux delete kustomization infra-apps --silent
flux delete kustomization infra-core --silent
