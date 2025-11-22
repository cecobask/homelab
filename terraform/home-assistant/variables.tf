variable "haos_download_url" {
  type    = string
  default = "https://github.com/home-assistant/operating-system/releases/download/16.3/haos_ova-16.3.qcow2.xz"
}

variable "proxmox_node_name" {
  type    = string
  default = "pve1"
}

variable "proxmox_vm_id" {
  type    = number
  default = 100
}
