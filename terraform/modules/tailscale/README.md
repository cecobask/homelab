<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.9.8 |
| <a name="requirement_tailscale"></a> [tailscale](#requirement\_tailscale) | 0.17.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_tailscale"></a> [tailscale](#provider\_tailscale) | 0.17.2 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [tailscale_acl.this](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/resources/acl) | resource |
| [tailscale_tailnet_key.this](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/resources/tailnet_key) | resource |

## Inputs

No inputs.

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_tailnet_key"></a> [tailnet\_key](#output\_tailnet\_key) | The tailnet key allows you to register new nodes without needing to sign in via a web browser |
<!-- END_TF_DOCS -->