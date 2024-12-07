terraform {
  required_providers {
    proxmox = {
      source  = "bpg/proxmox"
      version = "0.68.1"
    }
  }
}

provider "proxmox" {
  endpoint = var.endpoint
  insecure = var.insecure
  username = var.username
  password = var.password
  ssh {
    agent = true
  }
}
