terraform {
  required_version = "~> 1.9.8"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "homelab"
    }
  }
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.68.1"
    }
    talos = {
      source  = "siderolabs/talos"
      version = "0.6.1"
    }
  }
}

provider "proxmox" {
  endpoint = var.proxmox_endpoint
  username = var.proxmox_username
  password = var.proxmox_password
  insecure = true
  ssh {
    agent = true
  }
}
