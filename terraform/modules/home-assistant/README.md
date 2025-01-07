<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.9.8 |
| <a name="requirement_cloudflare"></a> [cloudflare](#requirement\_cloudflare) | 4.49.1 |
| <a name="requirement_proxmox"></a> [proxmox](#requirement\_proxmox) | 0.68.1 |
| <a name="requirement_tailscale"></a> [tailscale](#requirement\_tailscale) | 0.17.2 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_cloudflare"></a> [cloudflare](#provider\_cloudflare) | 4.49.1 |
| <a name="provider_proxmox"></a> [proxmox](#provider\_proxmox) | 0.68.1 |
| <a name="provider_tailscale"></a> [tailscale](#provider\_tailscale) | 0.17.2 |
| <a name="provider_terraform"></a> [terraform](#provider\_terraform) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [cloudflare_record.this](https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/resources/record) | resource |
| [proxmox_virtual_environment_file.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_file) | resource |
| [proxmox_virtual_environment_vm.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_vm) | resource |
| [terraform_data.this](https://registry.terraform.io/providers/hashicorp/terraform/latest/docs/resources/data) | resource |
| [cloudflare_zone.this](https://registry.terraform.io/providers/cloudflare/cloudflare/4.49.1/docs/data-sources/zone) | data source |
| [tailscale_device.this](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/data-sources/device) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cloudflare_zone"></a> [cloudflare\_zone](#input\_cloudflare\_zone) | Cloudflare zone name | `string` | n/a | yes |
| <a name="input_haos_version"></a> [haos\_version](#input\_haos\_version) | Home Assistant OS version | `string` | n/a | yes |
| <a name="input_vm"></a> [vm](#input\_vm) | The Home Assistant virtual machine configuration | <pre>object({<br/>    node_name   = string<br/>    vm_id       = number<br/>    cpu_cores   = number<br/>    ram_mb      = number<br/>    disk_gb     = number<br/>    mac_address = string<br/>  })</pre> | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->