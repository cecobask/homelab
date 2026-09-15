resource "proxmox_download_file" "this" {
  node_name               = var.proxmox_node_name
  datastore_id            = "local"
  content_type            = "iso"
  url                     = var.haos_download_url
  file_name               = "haos.qcow2.img"
  decompression_algorithm = "zst"
}

resource "proxmox_virtual_environment_vm" "this" {
  name      = "haos"
  node_name = var.proxmox_node_name
  vm_id     = var.proxmox_vm_id
  machine   = "q35"
  bios      = "ovmf"
  tags      = ["haos"]
  agent {
    enabled = true
  }
  cpu {
    cores = 2
    type  = "x86-64-v2-AES"
  }
  memory {
    dedicated = 1024 * 2
  }
  network_device {
    bridge      = "vmbr0"
    mac_address = "00:00:00:00:01:00"
    vlan_id     = 40
  }
  efi_disk {
    datastore_id = "local-lvm"
    type         = "4m"
  }
  disk {
    datastore_id = "local-lvm"
    import_from  = proxmox_download_file.this.id
    interface    = "scsi0"
    discard      = "on"
    size         = 64
  }
  operating_system {
    type = "l26"
  }
  usb {
    host = "10c4:ea60"
    usb3 = true
  }
}
