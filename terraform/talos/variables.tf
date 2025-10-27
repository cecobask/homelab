variable "image" {
  type = object({
    version      = string
    platform     = string
    architecture = string
    extensions   = list(string)
  })
}

variable "vms" {
  type = map(object({
    node_name    = string
    vm_id        = number
    cpu_cores    = number
    ram_mb       = number
    disk_gb      = number
    ipv4         = string
    mac_address  = string
    machine_type = string
  }))
}

variable "cluster" {
  type = object({
    name               = string
    bootstrap_node     = string
    gateway            = string
    kubernetes_version = string
    vip                = string
  })
}
