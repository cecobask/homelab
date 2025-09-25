image = {
  version      = "v1.11.0"
  platform     = "nocloud"
  architecture = "amd64"
  extensions   = ["intel-ucode", "qemu-guest-agent"]
}

vms = {
  talos-ctrl-01 = {
    node_name    = "pve1"
    vm_id        = 201
    cpu_cores    = 3
    ram_mb       = 1024 * 4
    disk_gb      = 100
    ipv4         = "192.168.10.201"
    mac_address  = "00:00:00:00:02:01"
    machine_type = "controlplane"
  }
  talos-ctrl-02 = {
    node_name    = "pve2"
    vm_id        = 202
    cpu_cores    = 3
    ram_mb       = 1024 * 4
    disk_gb      = 100
    ipv4         = "192.168.10.202"
    mac_address  = "00:00:00:00:02:02"
    machine_type = "controlplane"
  }
  talos-ctrl-03 = {
    node_name    = "pve2"
    vm_id        = 203
    cpu_cores    = 3
    ram_mb       = 1024 * 4
    disk_gb      = 100
    ipv4         = "192.168.10.203"
    mac_address  = "00:00:00:00:02:03"
    machine_type = "controlplane"
  }
  talos-work-01 = {
    node_name    = "pve1"
    vm_id        = 211
    cpu_cores    = 2
    ram_mb       = 1024 * 4
    disk_gb      = 100
    ipv4         = "192.168.10.211"
    mac_address  = "00:00:00:00:02:11"
    machine_type = "worker"
  }
}

cluster = {
  name               = "homelab"
  gateway            = "192.168.10.1"
  bootstrap_node     = "talos-ctrl-01"
  kubernetes_version = "v1.34.1"
  cilium_cli_version = "v0.18.7"
  cilium_version     = "v1.18.2"
}
