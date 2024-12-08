module "talos" {
  source  = "./modules/talos"
  release = "v1.8.3"
  extensions = [
    "qemu-guest-agent",
    "tailscale",
  ]
  vms = {
    talos-ctrl-01 = {
      node_name   = "pve1"
      vm_id       = "101"
      cpu_cores   = 3
      ram_mb      = 1024 * 4
      disk_gb     = 100
      mac_address = "BC:24:11:2E:C8:01"
    }
    talos-ctrl-02 = {
      node_name   = "pve2"
      vm_id       = "102"
      cpu_cores   = 3
      ram_mb      = 1024 * 4
      disk_gb     = 100
      mac_address = "BC:24:11:2E:C8:02"
    }
    talos-ctrl-03 = {
      node_name   = "pve3"
      vm_id       = "103"
      cpu_cores   = 3
      ram_mb      = 1024 * 4
      disk_gb     = 100
      mac_address = "BC:24:11:2E:C8:03"
    }
    talos-work-01 = {
      node_name   = "pve1"
      vm_id       = "104"
      cpu_cores   = 2
      ram_mb      = 1024 * 4
      disk_gb     = 100
      mac_address = "BC:24:11:2E:C8:04"
    }
  }
}
