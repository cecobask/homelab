module "talos" {
  source = "./modules/talos"
  cluster = {
    name           = "homelab"
    gateway        = "192.168.0.1"
    bootstrap_node = "talos-ctrl-01"
  }
  image = {
    version = "v1.9.0"
    extensions = [
      "intel-ucode",
      "qemu-guest-agent",
      "tailscale",
    ]
  }
  vms = {
    talos-ctrl-01 = {
      node_name    = "pve1"
      vm_id        = "101"
      cpu_cores    = 3
      ram_mb       = 1024 * 4
      disk_gb      = 100
      ip_address   = "192.168.0.101"
      mac_address  = "00:00:00:00:01:01"
      machine_type = "controlplane"
    }
    talos-ctrl-02 = {
      node_name    = "pve2"
      vm_id        = "102"
      cpu_cores    = 3
      ram_mb       = 1024 * 4
      disk_gb      = 100
      ip_address   = "192.168.0.102"
      mac_address  = "00:00:00:00:01:02"
      machine_type = "controlplane"
    }
    talos-ctrl-03 = {
      node_name    = "pve3"
      vm_id        = "103"
      cpu_cores    = 3
      ram_mb       = 1024 * 4
      disk_gb      = 100
      ip_address   = "192.168.0.103"
      mac_address  = "00:00:00:00:01:03"
      machine_type = "controlplane"
    }
    talos-work-01 = {
      node_name    = "pve1"
      vm_id        = "104"
      cpu_cores    = 2
      ram_mb       = 1024 * 4
      disk_gb      = 100
      ip_address   = "192.168.0.104"
      mac_address  = "00:00:00:00:01:04"
      machine_type = "worker"
    }
  }
}
