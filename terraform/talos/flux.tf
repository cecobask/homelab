resource "kubernetes_namespace_v1" "flux" {
  metadata {
    name = "flux-system"
  }
  lifecycle {
    ignore_changes = [metadata]
  }
}

resource "tls_private_key" "flux" {
  algorithm = "ED25519"
}

resource "kubernetes_secret_v1" "flux" {
  metadata {
    name      = kubernetes_namespace_v1.flux.metadata[0].name
    namespace = kubernetes_namespace_v1.flux.metadata[0].name
  }
  data = {
    "identity.pub" = tls_private_key.flux.public_key_openssh
    "identity"     = tls_private_key.flux.private_key_pem
    "known_hosts"  = "github.com ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBEmKSENjQEezOmxkZMy7opKgwFB9nkt5YRrYMjNuG5N87uRgg6CLrbo5wAdT/y6v0mKV0U2w0WZ2YB/++Tpockg="
  }
}

data "github_repository" "homelab" {
  full_name = "cecobask/homelab"
}

resource "github_repository_deploy_key" "flux" {
  repository = data.github_repository.homelab.name
  key        = tls_private_key.flux.public_key_openssh
  title      = "flux"
  read_only  = false
}

resource "flux_bootstrap_git" "homelab" {
  depends_on = [
    kubernetes_secret_v1.flux,
    github_repository_deploy_key.flux
  ]
  components = [
    "source-controller",
    "kustomize-controller",
    "helm-controller"
  ]
  disable_secret_creation = true
  embedded_manifests      = true
  path                    = "kubernetes/clusters/homelab"
}
