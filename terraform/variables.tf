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

variable "tailscale_oauth_client_id" {
  type        = string
  sensitive   = true
  description = "The Tailscale OAuth client ID"
}

variable "tailscale_oauth_client_secret" {
  type        = string
  sensitive   = true
  description = "The Tailscale OAuth client secret"
}
