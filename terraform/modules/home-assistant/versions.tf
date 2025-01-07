terraform {
  required_version = "~> 1.9.8"
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "4.49.1"
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
