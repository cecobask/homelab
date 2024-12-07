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
    file_id      = proxmox_virtual_environment_download_file.this[format("%s_%s", each.value.node_name, sha256(each.value.iso_download_url))].id
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

locals {
  deduped_download_files = distinct([
    for vm in var.vms : {
      node_name = vm.node_name
      url       = vm.iso_download_url
      file_name = vm.iso_file_name
    }
  ])
  download_files_map = {
    for ddf in local.deduped_download_files : format("%s_%s", ddf.node_name, sha256(ddf.url)) => ddf
  }
}

resource "proxmox_virtual_environment_download_file" "this" {
  for_each            = local.download_files_map
  node_name           = each.value.node_name
  content_type        = "iso"
  datastore_id        = "local"
  url                 = each.value.url
  file_name           = each.value.file_name
  overwrite           = true
  overwrite_unmanaged = true
}
