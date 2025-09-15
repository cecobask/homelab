terraform {
  required_version = "~> 1.9.8"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "home-assistant"
    }
  }
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

provider "cloudflare" {}

provider "proxmox" {
  insecure = true
  ssh {
    agent = true
  }
}

provider "tailscale" {}
