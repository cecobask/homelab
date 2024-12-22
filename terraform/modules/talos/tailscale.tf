resource "tailscale_acl" "this" {
  acl = jsonencode({
    acls = [{
      action = "accept"
      src    = ["*"]
      dst    = ["*:*"]
    }]
    tagOwners = {
      "tag:talos" = [
        "autogroup:admin",
      ],
      "tag:controlplane" = [
        "autogroup:admin",
      ],
      "tag:worker" = [
        "autogroup:admin",
      ]
    }
  })
  overwrite_existing_content = true
}

resource "tailscale_tailnet_key" "controlplane" {
  depends_on = [
    tailscale_acl.this,
  ]
  ephemeral     = true
  preauthorized = true
  reusable      = true
  tags = [
    "tag:talos",
    "tag:controlplane",
  ]
}

resource "tailscale_tailnet_key" "worker" {
  depends_on = [
    tailscale_acl.this,
  ]
  ephemeral     = true
  preauthorized = true
  reusable      = true
  tags = [
    "tag:talos",
    "tag:worker",
  ]
}

data "tailscale_device" "this" {
  for_each = talos_machine_configuration_apply.init
  hostname = each.key
  wait_for = "5m"
}
