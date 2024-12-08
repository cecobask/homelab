resource "proxmox_virtual_environment_vm" "this" {
  for_each      = var.vms
  name          = each.key
  node_name     = each.value.node_name
  vm_id         = each.value.vm_id
  machine       = "q35"
  scsi_hardware = "virtio-scsi-single"
  boot_order    = ["scsi0"]
  agent {
    enabled = true
  }
  cpu {
    cores = each.value.cpu_cores
    type  = "host"
  }
  memory {
    dedicated = each.value.ram_mb
  }
  network_device {
    bridge      = "vmbr0"
    mac_address = each.value.mac_address
  }
  disk {
    datastore_id = "local-lvm"
    interface    = "scsi0"
    iothread     = true
    discard      = "on"
    ssd          = true
    file_format  = "raw"
    size         = each.value.disk_gb
    file_id      = proxmox_virtual_environment_download_file.this[each.value.node_name].id
  }
  operating_system {
    type = "l26"
  }
  initialization {
    datastore_id = "local-lvm"
    ip_config {
      ipv4 {
        address = "dhcp"
      }
    }
  }
}

resource "proxmox_virtual_environment_download_file" "this" {
  for_each            = toset(values(var.vms)[*].node_name)
  node_name           = each.value
  content_type        = "iso"
  datastore_id        = "local"
  url                 = data.talos_image_factory_urls.this.urls.iso
  file_name           = format("talos_%s_%s_%s.iso", var.release, var.platform, var.architecture)
  overwrite           = true
  overwrite_unmanaged = true
}
