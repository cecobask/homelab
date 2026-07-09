data "github_repository" "homelab" {
  full_name = "cecobask/homelab"
}

resource "tls_private_key" "flux" {
  algorithm = "ED25519"
}

resource "github_repository_deploy_key" "flux" {
  repository = data.github_repository.homelab.name
  key        = tls_private_key.flux.public_key_openssh
  title      = "flux"
  read_only  = false
}

resource "flux_bootstrap_git" "homelab" {
  depends_on = [
    data.talos_cluster_health.this,
    github_repository_deploy_key.flux
  ]
  components = [
    "source-controller",
    "kustomize-controller",
    "helm-controller"
  ]
  delete_git_manifests   = false
  embedded_manifests     = true
  kustomization_override = file("${path.root}/config/flux-patches.yaml")
  path                   = "kubernetes/clusters/homelab"
}

resource "kubernetes_namespace_v1" "eso" {
  depends_on = [data.talos_cluster_health.this]
  metadata {
    name = "external-secrets"
  }
  lifecycle {
    ignore_changes = [
      metadata[0].labels,
      metadata[0].annotations,
    ]
  }
}

resource "kubernetes_secret_v1" "bitwarden" {
  metadata {
    name      = "bitwarden"
    namespace = kubernetes_namespace_v1.eso.metadata[0].name
  }
  data = {
    token = var.bitwarden_token
  }
}
