<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.9.8 |
| <a name="requirement_proxmox"></a> [proxmox](#requirement\_proxmox) | 0.68.1 |
| <a name="requirement_talos"></a> [talos](#requirement\_talos) | 0.6.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_proxmox"></a> [proxmox](#provider\_proxmox) | 0.68.1 |
| <a name="provider_talos"></a> [talos](#provider\_talos) | 0.6.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [proxmox_virtual_environment_download_file.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_download_file) | resource |
| [proxmox_virtual_environment_vm.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_vm) | resource |
| [talos_image_factory_schematic.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/image_factory_schematic) | resource |
| [talos_image_factory_extensions_versions.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/image_factory_extensions_versions) | data source |
| [talos_image_factory_urls.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/image_factory_urls) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_architecture"></a> [architecture](#input\_architecture) | The architecture to be used for the Talos image | `string` | `"amd64"` | no |
| <a name="input_extensions"></a> [extensions](#input\_extensions) | The list of extensions to be used for the Talos image | `list(string)` | n/a | yes |
| <a name="input_platform"></a> [platform](#input\_platform) | The platform to be used for the Talos image | `string` | `"metal"` | no |
| <a name="input_release"></a> [release](#input\_release) | The release version of the Talos image | `string` | n/a | yes |
| <a name="input_vms"></a> [vms](#input\_vms) | Virtual machine configurations | <pre>map(object({<br/>    node_name   = string<br/>    vm_id       = number<br/>    cpu_cores   = number<br/>    ram_mb      = number<br/>    disk_gb     = number<br/>    mac_address = string<br/>  }))</pre> | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->