variable "vms" {
  type = map(object({
    node_name        = string
    vm_id            = number
    cpu_cores        = number
    ram_mb           = number
    disk_gb          = number
    mac_address      = string
    iso_download_url = string
    iso_file_name    = string
  }))
  description = "Virtual machine configuration"
}

variable "endpoint" {
  type        = string
  description = "The endpoint for the PVE API"
}

variable "insecure" {
  type        = bool
  default     = true
  description = "Whether to skip the TLS verification step"
}

variable "username" {
  type        = string
  sensitive   = true
  description = "The username and realm for the PVE API"
}

variable "password" {
  type        = string
  sensitive   = true
  description = "The password for the PVE API"
}
