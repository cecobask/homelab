variable "haos_version" {
  type    = string
  default = "16.2"
}

variable "proxmox_node_name" {
  type    = string
  default = "pve1"
}

variable "proxmox_vm_id" {
  type    = number
  default = 100
}
