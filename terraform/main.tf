terraform {
  required_version = "1.9.8"
  cloud {
    organization = "cecobask"
    workspaces {
      name = "homelab"
    }
  }
}

module "proxmox" {
  source   = "./modules/proxmox"
  endpoint = "https://192.168.0.69:8006"
  username = var.proxmox_username
  password = var.proxmox_password
  vms = {
    talos-ctrl-01 = {
      node_name        = "pve1"
      vm_id            = "101"
      cpu_cores        = 3
      ram_mb           = 1024 * 4
      disk_gb          = 100
      mac_address      = "BC:24:11:2E:C8:01"
      iso_download_url = "https://factory.talos.dev/image/7d4c31cbd96db9f90c874990697c523482b2bae27fb4631d5583dcd9c281b1ff/v1.8.3/metal-amd64.iso"
      iso_file_name    = "talos-metal-amd64.iso"
    }
    talos-ctrl-02 = {
      node_name        = "pve2"
      vm_id            = "102"
      cpu_cores        = 3
      ram_mb           = 1024 * 4
      disk_gb          = 100
      mac_address      = "BC:24:11:2E:C8:02"
      iso_download_url = "https://factory.talos.dev/image/7d4c31cbd96db9f90c874990697c523482b2bae27fb4631d5583dcd9c281b1ff/v1.8.3/metal-amd64.iso"
      iso_file_name    = "talos-metal-amd64.iso"
    }
    talos-ctrl-03 = {
      node_name        = "pve3"
      vm_id            = "103"
      cpu_cores        = 3
      ram_mb           = 1024 * 4
      disk_gb          = 100
      mac_address      = "BC:24:11:2E:C8:03"
      iso_download_url = "https://factory.talos.dev/image/7d4c31cbd96db9f90c874990697c523482b2bae27fb4631d5583dcd9c281b1ff/v1.8.3/metal-amd64.iso"
      iso_file_name    = "talos-metal-amd64.iso"
    }
    talos-work-01 = {
      node_name        = "pve1"
      vm_id            = "104"
      cpu_cores        = 2
      ram_mb           = 1024 * 4
      disk_gb          = 100
      mac_address      = "BC:24:11:2E:C8:04"
      iso_download_url = "https://factory.talos.dev/image/7d4c31cbd96db9f90c874990697c523482b2bae27fb4631d5583dcd9c281b1ff/v1.8.3/metal-amd64.iso"
      iso_file_name    = "talos-metal-amd64.iso"
    }
  }
}
