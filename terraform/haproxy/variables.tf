variable "debian_download_url" {
  type    = string
  default = "https://cloud.debian.org/images/cloud/trixie/latest/debian-13-genericcloud-amd64.qcow2"
}

variable "proxmox_node_name" {
  type    = string
  default = "pve1"
}

variable "proxmox_vm_id" {
  type    = number
  default = 200
}
