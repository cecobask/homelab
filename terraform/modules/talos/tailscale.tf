resource "tailscale_acl" "this" {
  acl = jsonencode({
    acls = [{
      action = "accept"
      src    = ["*"]
      dst    = ["*:*"]
    }]
    tagOwners = {
      "tag:talos" = [
        "group:admin",
      ],
      "tag:controlplane" = [
        "group:admin",
      ],
      "tag:worker" = [
        "group:admin",
      ]
    }
    groups = {
      "group:admin" = [
        "baskski@gmail.com",
      ],
    }
  })
  overwrite_existing_content = true
}

resource "tailscale_tailnet_key" "controlplane" {
  depends_on = [
    tailscale_acl.this,
  ]
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
