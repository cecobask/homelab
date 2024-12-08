variable "release" {
  type        = string
  description = "The release version of the Talos image"
}

variable "extensions" {
  type        = list(string)
  description = "The list of extensions to be used for the Talos image"
}

variable "architecture" {
  type        = string
  default     = "amd64"
  description = "The architecture to be used for the Talos image"
}

variable "platform" {
  type        = string
  default     = "metal"
  description = "The platform to be used for the Talos image"
}

variable "vms" {
  type = map(object({
    node_name   = string
    vm_id       = number
    cpu_cores   = number
    ram_mb      = number
    disk_gb     = number
    mac_address = string
  }))
  description = "Virtual machine configurations"
}