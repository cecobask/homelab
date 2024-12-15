<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.9.8 |
| <a name="requirement_proxmox"></a> [proxmox](#requirement\_proxmox) | 0.68.1 |
| <a name="requirement_tailscale"></a> [tailscale](#requirement\_tailscale) | 0.17.2 |
| <a name="requirement_talos"></a> [talos](#requirement\_talos) | 0.6.1 |

## Providers

No providers.

## Modules

| Name | Source | Version |
|------|--------|---------|
| <a name="module_tailscale"></a> [tailscale](#module\_tailscale) | ./modules/tailscale | n/a |
| <a name="module_talos"></a> [talos](#module\_talos) | ./modules/talos | n/a |

## Resources

No resources.

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_proxmox_endpoint"></a> [proxmox\_endpoint](#input\_proxmox\_endpoint) | The endpoint for the Proxmox Virtual Environment API | `string` | n/a | yes |
| <a name="input_proxmox_password"></a> [proxmox\_password](#input\_proxmox\_password) | The password for the Proxmox Virtual Environment API | `string` | n/a | yes |
| <a name="input_proxmox_username"></a> [proxmox\_username](#input\_proxmox\_username) | The username and realm for the Proxmox Virtual Environment API | `string` | n/a | yes |
| <a name="input_tailscale_oauth_client_id"></a> [tailscale\_oauth\_client\_id](#input\_tailscale\_oauth\_client\_id) | The Tailscale OAuth client ID | `string` | n/a | yes |
| <a name="input_tailscale_oauth_client_secret"></a> [tailscale\_oauth\_client\_secret](#input\_tailscale\_oauth\_client\_secret) | The Tailscale OAuth client secret | `string` | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->