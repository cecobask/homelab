variable "proxmox_endpoint" {
  type        = string
  description = "The endpoint for the Proxmox Virtual Environment API"
}

variable "proxmox_username" {
  type        = string
  sensitive   = true
  description = "The username and realm for the Proxmox Virtual Environment API"
}

variable "proxmox_password" {
  type        = string
  sensitive   = true
  description = "The password for the Proxmox Virtual Environment API"
}
