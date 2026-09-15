terraform {
  required_version = ">= 1.13.3, < 2.0.0"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "home-assistant"
    }
  }
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.113.1"
    }
  }
}

provider "proxmox" {
  insecure = true
  ssh {
    agent = true
  }
}
