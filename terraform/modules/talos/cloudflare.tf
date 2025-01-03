data "cloudflare_zone" "this" {
  name = var.cluster.cloudflare_zone
}

resource "cloudflare_record" "this" {
  for_each = local.controlplane_names
  zone_id  = data.cloudflare_zone.this.zone_id
  type     = "A"
  name     = "talos"
  content  = talos_machine_configuration_apply.final[each.key].node
  ttl      = 300
}
