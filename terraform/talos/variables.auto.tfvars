image = {
  version      = "v1.12.0"
  platform     = "nocloud"
  architecture = "amd64"
  extensions = [
    "intel-ucode",
    "i915",
    "qemu-guest-agent"
  ]
}

vms = {
  talos-ctrl-01 = {
    node_name    = "pve1"
    vm_id        = 201
    cpu_cores    = 2
    ram_mb       = 1024 * 3
    disk_gb      = 100
    ipv4         = "192.168.40.201"
    mac_address  = "00:00:00:00:02:01"
    machine_type = "controlplane"
  }
  talos-ctrl-02 = {
    node_name    = "pve2"
    vm_id        = 202
    cpu_cores    = 2
    ram_mb       = 1024 * 3
    disk_gb      = 100
    ipv4         = "192.168.40.202"
    mac_address  = "00:00:00:00:02:02"
    machine_type = "controlplane"
  }
  talos-ctrl-03 = {
    node_name    = "pve2"
    vm_id        = 203
    cpu_cores    = 2
    ram_mb       = 1024 * 3
    disk_gb      = 100
    ipv4         = "192.168.40.203"
    mac_address  = "00:00:00:00:02:03"
    machine_type = "controlplane"
  }
  talos-work-01 = {
    node_name    = "pve1"
    vm_id        = 211
    cpu_cores    = 2
    ram_mb       = 1024 * 5
    disk_gb      = 200
    ipv4         = "192.168.40.211"
    mac_address  = "00:00:00:00:02:11"
    machine_type = "worker"
    gpu          = true
  }
  talos-work-02 = {
    node_name    = "pve2"
    vm_id        = 212
    cpu_cores    = 4
    ram_mb       = 1024 * 5
    disk_gb      = 200
    ipv4         = "192.168.40.212"
    mac_address  = "00:00:00:00:02:12"
    machine_type = "worker"
    gpu          = true
  }
}

cluster = {
  name               = "homelab"
  bootstrap_node     = "talos-ctrl-01"
  kubernetes_version = "v1.35.0"
  gateway            = "192.168.40.1"
  vip                = "192.168.40.2"
}
