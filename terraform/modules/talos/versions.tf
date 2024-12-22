terraform {
  required_version = "~> 1.9.8"
  required_providers {
    talos = {
      source  = "siderolabs/talos"
      version = "0.6.1"
    }
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.68.1"
    }
    tailscale = {
      source  = "tailscale/tailscale"
      version = "0.17.2"
    }
  }
}
