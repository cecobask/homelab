terraform {
  required_version = ">= 1.13.3, < 2.0.0"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "talos"
    }
  }
  required_providers {
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

provider "proxmox" {
  insecure = true
  ssh {
    agent = true
  }
}

provider "talos" {}
