terraform {
  required_version = ">= 1.13.3, < 2.0.0"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "talos"
    }
  }
  required_providers {
    kubernetes = {
      source  = "hashicorp/kubernetes"
      version = ">=2.38.0"
    }
    kustomization = {
      source  = "kbst/kustomization"
      version = "0.9.6"
    }
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.84.1"
    }
    talos = {
      source  = "siderolabs/talos"
      version = "0.9.0"
    }
  }
}

provider "kubernetes" {
  host                   = talos_cluster_kubeconfig.this.kubernetes_client_configuration.host
  client_certificate     = base64decode(talos_cluster_kubeconfig.this.kubernetes_client_configuration.client_certificate)
  client_key             = base64decode(talos_cluster_kubeconfig.this.kubernetes_client_configuration.client_key)
  cluster_ca_certificate = base64decode(talos_cluster_kubeconfig.this.kubernetes_client_configuration.ca_certificate)
}

provider "kustomization" {
  kubeconfig_raw = talos_cluster_kubeconfig.this.kubeconfig_raw
}

provider "proxmox" {
  insecure = true
  ssh {
    agent = true
  }
}

provider "talos" {}
