resource "terraform_data" "this" {
  input = {
    url = format("https://github.com/home-assistant/operating-system/releases/download/%s/haos_ova-%s.qcow2.xz",
      var.haos_version,
      var.haos_version,
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
  node_name    = var.vm.node_name
  datastore_id = "local"
  content_type = "iso"
  source_file {
    path      = terraform_data.this.output.filename
    file_name = format("haos_ova-%s.img", var.haos_version)
  }
}
