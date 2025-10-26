data "kustomization_build" "cilium" {
  depends_on = [data.talos_cluster_health.this]
  path       = "../../kubernetes/platform/cilium"
  kustomize_options {
    enable_helm = true
    helm_path   = "helm"
  }
}

resource "kustomization_resource" "cilium_p0" {
  for_each = data.kustomization_build.cilium.ids_prio[0]
  manifest = data.kustomization_build.cilium.manifests[each.value]
}

resource "kustomization_resource" "cilium_p1" {
  depends_on = [kustomization_resource.cilium_p0]
  for_each   = data.kustomization_build.cilium.ids_prio[1]
  manifest   = data.kustomization_build.cilium.manifests[each.value]
  lifecycle {
    ignore_changes = [manifest]
  }
}

data "kustomization_build" "argocd" {
  depends_on = [data.talos_cluster_health.this]
  path       = "../../kubernetes/platform/argocd"
  kustomize_options {
    enable_helm = true
    helm_path   = "helm"
  }
}

resource "kustomization_resource" "argocd_p0" {
  for_each = data.kustomization_build.argocd.ids_prio[0]
  manifest = data.kustomization_build.argocd.manifests[each.value]
}

resource "kustomization_resource" "argocd_p1" {
  depends_on = [kustomization_resource.argocd_p0]
  for_each   = data.kustomization_build.argocd.ids_prio[1]
  manifest   = data.kustomization_build.argocd.manifests[each.value]
}

data "kustomization_build" "sealedsecrets" {
  depends_on = [data.talos_cluster_health.this]
  path       = "../../kubernetes/platform/sealed-secrets"
  kustomize_options {
    enable_helm = true
    helm_path   = "helm"
  }
}

resource "kustomization_resource" "sealedsecrets_p0" {
  for_each = data.kustomization_build.sealedsecrets.ids_prio[0]
  manifest = data.kustomization_build.sealedsecrets.manifests[each.value]
}

resource "kubernetes_secret" "sealedsecrets" {
  depends_on = [kustomization_resource.sealedsecrets_p0]
  type       = "kubernetes.io/tls"
  metadata {
    name      = "sealed-secrets-key"
    namespace = "kube-system"
  }
  data = {
    "tls.crt" = file("../../kubernetes/platform/sealed-secrets/tls.crt")
    "tls.key" = file("../../kubernetes/platform/sealed-secrets/tls.key")
  }
}

resource "kustomization_resource" "sealedsecrets_p1" {
  depends_on = [kubernetes_secret.sealedsecrets]
  for_each   = data.kustomization_build.sealedsecrets.ids_prio[1]
  manifest   = data.kustomization_build.sealedsecrets.manifests[each.value]
}
