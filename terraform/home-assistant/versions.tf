terraform {
  required_version = "~> 1.9.8"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "home-assistant"
    }
  }
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.68.1"
    }
  }
}

provider "proxmox" {
  insecure = true
  ssh {
    agent = true
  }
}
