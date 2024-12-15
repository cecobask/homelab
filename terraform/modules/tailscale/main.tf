resource "tailscale_acl" "this" {
  acl = jsonencode({
    acls = [{
      action = "accept"
      src    = ["*"]
      dst    = ["*:*"]
    }]
    tagOwners = {
      "tag:terraform" = [
        "autogroup:member",
      ]
    }
  })
  overwrite_existing_content = true
}

resource "tailscale_tailnet_key" "this" {
  depends_on = [
    tailscale_acl.this,
  ]
  preauthorized = true
  reusable      = true
  tags = [
    "tag:terraform",
  ]
}
