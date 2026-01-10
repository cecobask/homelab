resource "proxmox_virtual_environment_download_file" "this" {
  for_each     = toset(values(var.vms)[*].node_name)
  node_name    = each.value
  content_type = "iso"
  datastore_id = "local"
  url          = data.talos_image_factory_urls.this.urls.iso
  file_name    = "talos.iso"
}

resource "proxmox_virtual_environment_vm" "this" {
  for_each      = var.vms
  name          = each.key
  node_name     = each.value.node_name
  vm_id         = each.value.vm_id
  machine       = "q35"
  scsi_hardware = "virtio-scsi-single"
  tags          = ["talos", each.value.machine_type]
  agent {
    enabled = true
  }
  cpu {
    cores = each.value.cpu_cores
    type  = "x86-64-v2-AES"
  }
  memory {
    dedicated = each.value.ram_mb
  }
  network_device {
    bridge      = "vmbr0"
    mac_address = each.value.mac_address
    vlan_id     = 40
  }
  disk {
    datastore_id = "local-lvm"
    file_id      = proxmox_virtual_environment_download_file.this[each.value.node_name].id
    interface    = "scsi0"
    iothread     = true
    discard      = "on"
    size         = each.value.disk_gb
  }
  operating_system {
    type = "l26"
  }
  initialization {
    datastore_id = "local-lvm"
    ip_config {
      ipv4 {
        address = format("%s/24", each.value.ipv4)
        gateway = var.cluster.gateway
      }
    }
  }
  dynamic "hostpci" {
    for_each = each.value.gpu ? [1] : []
    content {
      device = "hostpci0"
      id     = "0000:00:02.0"
      pcie   = true
      rombar = true
    }
  }
}
