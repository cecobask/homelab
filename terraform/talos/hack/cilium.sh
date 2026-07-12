#!/bin/bash
set -e
set -o pipefail
set -u

REPO_ROOT="$(git rev-parse --show-toplevel)"
HELM_RELEASE_FILE="$REPO_ROOT/kubernetes/infra/apps/cilium/helmrelease.yaml"
HELM_REPO_FILE="$REPO_ROOT/kubernetes/infra/apps/cilium/source.yaml"
HELM_CHART_URL="$(yq -r '.spec.url' "$HELM_REPO_FILE")"
HELM_CHART_NAME="$(yq -r '.spec.chart.spec.chart' "$HELM_RELEASE_FILE")"
HELM_CHART_VERSION="$(yq -r '.spec.chart.spec.version' "$HELM_RELEASE_FILE")"
NAMESPACE="$(yq -r '.metadata.namespace' "$HELM_RELEASE_FILE")"
NAMESPACE_FILE="$REPO_ROOT/kubernetes/infra/apps/cilium/namespace.yaml"
OUT="$REPO_ROOT/terraform/talos/config/cilium.yaml"

VALUES="$(mktemp)"
trap 'rm -f "$VALUES"' EXIT
yq '.spec.values' "$HELM_RELEASE_FILE" >"$VALUES"

helm repo add "$HELM_CHART_NAME" "$HELM_CHART_URL"
helm repo update "$HELM_CHART_NAME"
{
	cat "$NAMESPACE_FILE"
	echo
	helm template "$HELM_CHART_NAME" "$HELM_CHART_NAME/$HELM_CHART_NAME" \
		--version="$HELM_CHART_VERSION" \
		--namespace="$NAMESPACE" \
		--values "$VALUES"
} >"$OUT"
