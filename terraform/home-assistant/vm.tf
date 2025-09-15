resource "terraform_data" "this" {
  input = {
    url = format("https://github.com/home-assistant/operating-system/releases/download/%s/haos_ova-%s.qcow2.xz",
      local.haos_version,
      local.haos_version,
    )
    filename = "haos_ova.qcow2"
  }
  provisioner "local-exec" {
    when    = create
    command = format("curl -sL %s | xz -d > %s", self.input.url, self.input.filename)
  }
  provisioner "local-exec" {
    when    = destroy
    command = format("rm -f %q", self.input.filename)
  }
}

resource "proxmox_virtual_environment_file" "this" {
  node_name    = local.proxmox_node
  datastore_id = "local"
  content_type = "iso"
  source_file {
    path      = terraform_data.this.output.filename
    file_name = format("haos_ova-%s.img", local.haos_version)
  }
}

resource "proxmox_virtual_environment_vm" "this" {
  name          = local.proxmox_vm_name
  node_name     = local.proxmox_node
  vm_id         = local.proxmox_vm_id
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
    cores = 3
    type  = "x86-64-v2-AES"
  }
  memory {
    dedicated = 1024 * 4
  }
  network_device {
    bridge      = "vmbr0"
    mac_address = "00:00:00:00:01:00"
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
    size         = 128
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
