variable "cloudflare_zone" {
  type        = string
  description = "Cloudflare zone name"
}

variable "haos_version" {
  type        = string
  description = "Home Assistant OS version"
}

variable "vm" {
  type = object({
    node_name   = string
    vm_id       = number
    cpu_cores   = number
    ram_mb      = number
    disk_gb     = number
    mac_address = string
  })
  description = "The Home Assistant virtual machine configuration"
}
