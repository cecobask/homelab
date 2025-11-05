data "kustomization_build" "bootstrap" {
  depends_on = [data.talos_cluster_health.this]
  path       = "../../kubernetes/bootstrap"
  kustomize_options {
    enable_helm = true
    helm_path   = "helm"
  }
}

resource "kustomization_resource" "bootstrap_p0" {
  for_each = data.kustomization_build.bootstrap.ids_prio[0]
  manifest = data.kustomization_build.bootstrap.manifests[each.value]
}

resource "kubernetes_secret" "sealedsecrets" {
  depends_on = [kustomization_resource.bootstrap_p0]
  type       = "kubernetes.io/tls"
  metadata {
    name      = "sealed-secrets-key"
    namespace = "sealed-secrets"
  }
  data = {
    "tls.crt" = file("../../kubernetes/platform/sealed-secrets/tls.crt")
    "tls.key" = file("../../kubernetes/platform/sealed-secrets/tls.key")
  }
}

resource "kustomization_resource" "bootstrap_p1" {
  depends_on = [kubernetes_secret.sealedsecrets]
  for_each   = data.kustomization_build.bootstrap.ids_prio[1]
  manifest   = data.kustomization_build.bootstrap.manifests[each.value]
}
