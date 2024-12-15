variable "cluster" {
  type = object({
    name        = string
    gateway     = string
    tailnet_key = string
  })
  description = "The Talos cluster configuration"
}

variable "image" {
  type = object({
    version      = string
    platform     = optional(string, "nocloud")
    architecture = optional(string, "amd64")
    extensions   = list(string)
  })
  description = "The Talos image factory configuration"
}

variable "vms" {
  type = map(object({
    node_name    = string
    vm_id        = number
    cpu_cores    = number
    ram_mb       = number
    disk_gb      = number
    ip_address   = string
    mac_address  = string
    machine_type = string
  }))
  description = "The Talos virtual machines configuration"
}
