variable "cluster" {
  type = object({
    name           = string
    gateway        = string
    bootstrap_node = string
  })
  description = "The Talos cluster configuration"
  validation {
    condition     = one([for node_name, vm in var.vms : vm if node_name == var.cluster.bootstrap_node && vm.machine_type == "controlplane"]) != null
    error_message = "Must have exactly one bootstrap control plane node"
  }
}

variable "image" {
  type = object({
    version      = string
    platform     = optional(string, "nocloud")
    architecture = optional(string, "amd64")
    extensions   = list(string)
  })
  description = "The Talos image factory configuration"
  validation {
    condition     = can(regex("^v(?P<major>0|[1-9]\\d*)\\.(?P<minor>0|[1-9]\\d*)\\.(?P<patch>0|[1-9]\\d*)(?:-(?P<prerelease>(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\\.(?:0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\\+(?P<buildmetadata>[0-9a-zA-Z-]+(?:\\.[0-9a-zA-Z-]+)*))?$", var.image.version))
    error_message = "Talos image version must follow semantic versioning format"
  }
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
