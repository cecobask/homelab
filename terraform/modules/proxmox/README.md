<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_proxmox"></a> [proxmox](#requirement\_proxmox) | 0.68.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_proxmox"></a> [proxmox](#provider\_proxmox) | 0.68.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [proxmox_virtual_environment_download_file.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_download_file) | resource |
| [proxmox_virtual_environment_vm.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_vm) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_endpoint"></a> [endpoint](#input\_endpoint) | The endpoint for the PVE API | `string` | n/a | yes |
| <a name="input_insecure"></a> [insecure](#input\_insecure) | Whether to skip the TLS verification step | `bool` | `true` | no |
| <a name="input_password"></a> [password](#input\_password) | The password for the PVE API | `string` | n/a | yes |
| <a name="input_username"></a> [username](#input\_username) | The username and realm for the PVE API | `string` | n/a | yes |
| <a name="input_vms"></a> [vms](#input\_vms) | Virtual machine configuration | <pre>map(object({<br/>    node_name        = string<br/>    vm_id            = number<br/>    cpu_cores        = number<br/>    ram_mb           = number<br/>    disk_gb          = number<br/>    mac_address      = string<br/>    iso_download_url = string<br/>    iso_file_name    = string<br/>  }))</pre> | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->