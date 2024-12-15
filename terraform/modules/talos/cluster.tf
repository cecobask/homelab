locals {
  node_ips = [
    for vm in var.vms : vm.ip_address
  ]
  controlplane_ips = [
    for vm in var.vms : vm.ip_address if vm.machine_type == "controlplane"
  ]
  # worker_ips = [
  #   for vm in var.vms : vm.ip_address if vm.machine_type == "worker"
  # ]
}

resource "talos_machine_secrets" "this" {
  talos_version = var.image.version
}

data "talos_client_configuration" "this" {
  cluster_name         = var.cluster.name
  client_configuration = talos_machine_secrets.this.client_configuration
  nodes                = local.node_ips
  endpoints            = local.controlplane_ips
}

data "talos_machine_configuration" "this" {
  for_each         = var.vms
  cluster_name     = var.cluster.name
  cluster_endpoint = format("https://%s:6443", local.controlplane_ips[0])
  talos_version    = var.image.version
  machine_type     = each.value.machine_type
  machine_secrets  = talos_machine_secrets.this.machine_secrets
  config_patches = [
    templatefile(format("%s/templates/node.tftpl", path.module), {
      installer_url = data.talos_image_factory_urls.this.urls.installer,
    }),
    templatefile(format("%s/templates/tailscale.tftpl", path.module), {
      hostname    = each.key
      tailnet_key = var.cluster.tailnet_key,
    })
  ]
}

resource "talos_machine_configuration_apply" "this" {
  for_each                    = var.vms
  node                        = each.value.ip_address
  client_configuration        = talos_machine_secrets.this.client_configuration
  machine_configuration_input = data.talos_machine_configuration.this[each.key].machine_configuration
  lifecycle {
    replace_triggered_by = [
      proxmox_virtual_environment_vm.this[each.key]
    ]
  }
}

resource "talos_machine_bootstrap" "this" {
  depends_on = [
    talos_machine_configuration_apply.this
  ]
  node                 = local.controlplane_ips[0]
  endpoint             = local.controlplane_ips[0]
  client_configuration = talos_machine_secrets.this.client_configuration
}

# data "talos_cluster_health" "this" {
#   depends_on = [
#     talos_machine_bootstrap.this
#   ]
#   client_configuration = data.talos_client_configuration.this.client_configuration
#   control_plane_nodes  = local.controlplane_ips
#   worker_nodes         = local.worker_ips
#   endpoints            = data.talos_client_configuration.this.endpoints
#   timeouts = {
#     read = "10m"
#   }
# }

resource "talos_cluster_kubeconfig" "this" {
  depends_on = [
    talos_machine_bootstrap.this
  ]
  node                 = local.controlplane_ips[0]
  endpoint             = local.controlplane_ips[0]
  client_configuration = data.talos_client_configuration.this.client_configuration
  timeouts = {
    read = "1m"
  }
}
