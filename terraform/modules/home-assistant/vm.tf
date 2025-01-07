resource "proxmox_virtual_environment_vm" "this" {
  name          = "haos"
  node_name     = var.vm.node_name
  vm_id         = var.vm.vm_id
  machine       = "q35"
  scsi_hardware = "virtio-scsi-single"
  bios          = "ovmf"
  tags = [
    "haos",
  ]
  agent {
    enabled = true
  }
  cpu {
    cores = var.vm.cpu_cores
    type  = "x86-64-v2-AES"
  }
  memory {
    dedicated = var.vm.ram_mb
  }
  network_device {
    bridge      = "vmbr0"
    mac_address = var.vm.mac_address
  }
  efi_disk {
    datastore_id = "local-lvm"
    type         = "4m"
  }
  disk {
    datastore_id = "local-lvm"
    file_id      = proxmox_virtual_environment_file.this.id
    interface    = "scsi0"
    iothread     = true
    discard      = "on"
    ssd          = true
    file_format  = "raw"
    size         = var.vm.disk_gb
  }
  operating_system {
    type = "l26"
  }
  usb {
    host = "10c4:ea60"
    usb3 = true
  }
  lifecycle {
    prevent_destroy = true
  }
}
