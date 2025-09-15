data "tailscale_device" "this" {
  depends_on = [
    proxmox_virtual_environment_vm.this,
  ]
  hostname = "homeassistant"
  wait_for = "15m" # wait for manual installation of the tailscale addon
}

data "cloudflare_zone" "this" {
  name = local.cloudflare_zone
}

resource "cloudflare_record" "this" {
  zone_id = data.cloudflare_zone.this.zone_id
  type    = "A"
  name    = local.proxmox_vm_name
  content = data.tailscale_device.this.addresses[0]
  ttl     = 300
}
