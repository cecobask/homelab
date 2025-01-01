<!-- BEGIN_TF_DOCS -->
## Requirements

| Name | Version |
|------|---------|
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | ~> 1.9.8 |
| <a name="requirement_proxmox"></a> [proxmox](#requirement\_proxmox) | 0.68.1 |
| <a name="requirement_tailscale"></a> [tailscale](#requirement\_tailscale) | 0.17.2 |
| <a name="requirement_talos"></a> [talos](#requirement\_talos) | 0.6.1 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_proxmox"></a> [proxmox](#provider\_proxmox) | 0.68.1 |
| <a name="provider_tailscale"></a> [tailscale](#provider\_tailscale) | 0.17.2 |
| <a name="provider_talos"></a> [talos](#provider\_talos) | 0.6.1 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [proxmox_virtual_environment_download_file.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_download_file) | resource |
| [proxmox_virtual_environment_vm.this](https://registry.terraform.io/providers/bpg/proxmox/0.68.1/docs/resources/virtual_environment_vm) | resource |
| [tailscale_acl.this](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/resources/acl) | resource |
| [tailscale_tailnet_key.controlplane](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/resources/tailnet_key) | resource |
| [tailscale_tailnet_key.worker](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/resources/tailnet_key) | resource |
| [talos_cluster_kubeconfig.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/cluster_kubeconfig) | resource |
| [talos_image_factory_schematic.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/image_factory_schematic) | resource |
| [talos_machine_bootstrap.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/machine_bootstrap) | resource |
| [talos_machine_configuration_apply.final](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/machine_configuration_apply) | resource |
| [talos_machine_configuration_apply.init](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/machine_configuration_apply) | resource |
| [talos_machine_secrets.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/resources/machine_secrets) | resource |
| [tailscale_device.this](https://registry.terraform.io/providers/tailscale/tailscale/0.17.2/docs/data-sources/device) | data source |
| [talos_client_configuration.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/client_configuration) | data source |
| [talos_cluster_health.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/cluster_health) | data source |
| [talos_image_factory_extensions_versions.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/image_factory_extensions_versions) | data source |
| [talos_image_factory_urls.this](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/image_factory_urls) | data source |
| [talos_machine_configuration.final](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/machine_configuration) | data source |
| [talos_machine_configuration.init](https://registry.terraform.io/providers/siderolabs/talos/0.6.1/docs/data-sources/machine_configuration) | data source |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_cluster"></a> [cluster](#input\_cluster) | The Talos cluster configuration | <pre>object({<br/>    name           = string<br/>    gateway        = string<br/>    bootstrap_node = string<br/>  })</pre> | n/a | yes |
| <a name="input_image"></a> [image](#input\_image) | The Talos image factory configuration | <pre>object({<br/>    version      = string<br/>    platform     = optional(string, "nocloud")<br/>    architecture = optional(string, "amd64")<br/>    extensions   = list(string)<br/>  })</pre> | n/a | yes |
| <a name="input_vms"></a> [vms](#input\_vms) | The Talos virtual machines configuration | <pre>map(object({<br/>    node_name    = string<br/>    vm_id        = number<br/>    cpu_cores    = number<br/>    ram_mb       = number<br/>    disk_gb      = number<br/>    ipv4         = string<br/>    mac_address  = string<br/>    machine_type = string<br/>  }))</pre> | n/a | yes |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_kubeconfig"></a> [kubeconfig](#output\_kubeconfig) | n/a |
| <a name="output_secrets"></a> [secrets](#output\_secrets) | n/a |
| <a name="output_talosconfig"></a> [talosconfig](#output\_talosconfig) | n/a |
<!-- END_TF_DOCS -->