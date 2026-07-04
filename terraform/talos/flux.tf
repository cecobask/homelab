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
  depends_on = [github_repository_deploy_key.flux]
  components = [
    "source-controller",
    "kustomize-controller",
    "helm-controller"
  ]
  embedded_manifests = true
  path               = "kubernetes/clusters/homelab"
}
