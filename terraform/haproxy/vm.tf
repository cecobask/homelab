resource "proxmox_virtual_environment_vm" "this" {
  name          = "haproxy"
  node_name     = var.proxmox_node_name
  vm_id         = var.proxmox_vm_id
  machine       = "q35"
  scsi_hardware = "virtio-scsi-single"
  tags          = ["haproxy"]
  agent {
    enabled = true
  }
  cpu {
    cores = 3
    type  = "x86-64-v2-AES"
  }
  memory {
    dedicated = 1024 * 4
  }
  network_device {
    bridge      = "vmbr0"
    mac_address = "00:00:00:00:02:00"
    vlan_id     = 40
  }
  disk {
    datastore_id = "local-lvm"
    import_from  = proxmox_virtual_environment_download_file.this.id
    interface    = "scsi0"
    iothread     = true
    discard      = "on"
    size         = 10
  }
  operating_system {
    type = "l26"
  }
  initialization {
    ip_config {
      ipv4 {
        address = "192.168.40.200/24"
        gateway = "192.168.40.1"
      }
    }
    user_data_file_id = proxmox_virtual_environment_file.this.id
  }
}

resource "proxmox_virtual_environment_download_file" "this" {
  node_name    = var.proxmox_node_name
  content_type = "import"
  datastore_id = "local"
  url          = var.debian_download_url
  file_name    = "debian-haproxy.qcow2"
}

resource "proxmox_virtual_environment_file" "this" {
  node_name    = var.proxmox_node_name
  content_type = "snippets"
  datastore_id = "local"
  source_file {
    path      = "cloudinit.yaml"
    file_name = "cloudinit-haproxy.yaml"
  }
}
