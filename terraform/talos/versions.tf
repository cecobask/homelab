terraform {
  required_version = "~> 1.9.8"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "talos"
    }
  }
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.68.1"
    }
    talos = {
      source  = "siderolabs/talos"
      version = "0.9.0"
    }
  }
}

provider "proxmox" {
  insecure = true
  ssh {
    agent = true
  }
}

provider "talos" {}
