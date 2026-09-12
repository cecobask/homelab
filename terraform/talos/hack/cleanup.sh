#!/bin/bash
set -e
set -o pipefail
set -u

flux suspend kustomization flux-system
flux delete kustomization services --silent
flux delete kustomization media-apps --silent
flux delete kustomization media-core --silent
flux delete kustomization infra-apps --silent
kubectl delete pvc --all-namespaces --ignore-not-found --selector=cluster-destroy-delete --wait
flux delete kustomization infra-core --silent
