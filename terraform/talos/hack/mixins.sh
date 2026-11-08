#!/bin/bash
set -e
set -o pipefail
set -u

MIXINS_TAG=version-1.5.6
MIXINS_OUTPUT_DIR=mixins-output
DASHBOARDS_DIR="$MIXINS_OUTPUT_DIR/dashboards"
REPO_ROOT="$(git rev-parse --show-toplevel)"
GRAFANA_DIR="$REPO_ROOT/kubernetes/infra/apps/grafana"
PROMETHEUS_DIR="$REPO_ROOT/kubernetes/infra/apps/prometheus"

jb install "github.com/kubernetes-monitoring/kubernetes-mixin@$MIXINS_TAG"
mkdir -p "$DASHBOARDS_DIR"
jsonnet -J vendor -S -e 'std.manifestYamlDoc((import "mixin.libsonnet").prometheusAlerts)' > "$MIXINS_OUTPUT_DIR/alerts.yaml"
jsonnet -J vendor -S -e 'std.manifestYamlDoc((import "mixin.libsonnet").prometheusRules)' > "$MIXINS_OUTPUT_DIR/rules.yaml"
jsonnet -J vendor -m "$DASHBOARDS_DIR" -e '(import "mixin.libsonnet").grafanaDashboards'

yq -i 'del(.groups[] | select(.name == "kubernetes-system-kube-proxy"))' "$MIXINS_OUTPUT_DIR/alerts.yaml"
yq -i '.spec = load("'"$MIXINS_OUTPUT_DIR"'/alerts.yaml")' "$PROMETHEUS_DIR/alerts.yaml"
yq -i '.spec = load("'"$MIXINS_OUTPUT_DIR"'/rules.yaml")' "$PROMETHEUS_DIR/rules.yaml"
find "$DASHBOARDS_DIR" -type f \( -iname '*windows*' -o -iname '*proxy*' \) -delete
cp -r "$DASHBOARDS_DIR" "$GRAFANA_DIR"
