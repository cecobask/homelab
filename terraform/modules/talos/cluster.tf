locals {
  node_ips = [
    for d in data.tailscale_device.this : d.addresses[0] if contains(d.tags, "tag:talos")
  ]
  controlplane_ips = [
    for d in data.tailscale_device.this : d.addresses[0] if contains(d.tags, "tag:controlplane")
  ]
  worker_ips = [
    for d in data.tailscale_device.this : d.addresses[0] if contains(d.tags, "tag:worker")
  ]
  bootstrap_node_ip = one([
    for d in data.tailscale_device.this : d.addresses[0] if contains(d.tags, "tag:controlplane") && d.hostname == var.cluster.bootstrap_node
  ])
}

resource "talos_machine_secrets" "this" {
  talos_version = var.image.version
}

data "talos_machine_configuration" "init" {
  for_each         = var.vms
  cluster_name     = var.cluster.name
  cluster_endpoint = format("https://%s:6443", var.vms[var.cluster.bootstrap_node].ip_address)
  talos_version    = var.image.version
  machine_type     = each.value.machine_type
  machine_secrets  = talos_machine_secrets.this.machine_secrets
  config_patches = flatten([
    templatefile(format("%s/templates/node.tftpl", path.module), {
      hostname      = each.key
      installer_url = data.talos_image_factory_urls.this.urls.installer
    }),
    templatefile(format("%s/templates/tailscale.tftpl", path.module), {
      hostname    = each.key
      tailnet_key = each.value.machine_type == "controlplane" ? tailscale_tailnet_key.controlplane.key : tailscale_tailnet_key.worker.key
    }),
    each.value.machine_type == "controlplane" ? [file(format("%s/templates/controlplane.tftpl", path.module))] : []
  ])
}

resource "talos_machine_configuration_apply" "init" {
  for_each                    = data.talos_machine_configuration.init
  node                        = var.vms[each.key].ip_address
  client_configuration        = talos_machine_secrets.this.client_configuration
  machine_configuration_input = each.value.machine_configuration
  lifecycle {
    replace_triggered_by = [
      proxmox_virtual_environment_vm.this[each.key]
    ]
    ignore_changes = [
      machine_configuration_input,
    ]
  }
}

resource "talos_machine_bootstrap" "this" {
  depends_on = [
    talos_machine_configuration_apply.init,
  ]
  node                 = var.vms[var.cluster.bootstrap_node].ip_address
  client_configuration = talos_machine_secrets.this.client_configuration
}

data "talos_machine_configuration" "final" {
  depends_on = [
    talos_machine_bootstrap.this,
  ]
  for_each         = data.talos_machine_configuration.init
  cluster_name     = each.value.cluster_name
  cluster_endpoint = format("https://%s:6443", local.bootstrap_node_ip)
  talos_version    = each.value.talos_version
  machine_type     = each.value.machine_type
  machine_secrets  = each.value.machine_secrets
  config_patches   = each.value.config_patches
}

resource "talos_machine_configuration_apply" "final" {
  for_each                    = data.talos_machine_configuration.final
  node                        = data.tailscale_device.this[each.key].addresses[0]
  client_configuration        = talos_machine_secrets.this.client_configuration
  machine_configuration_input = each.value.machine_configuration
  apply_mode                  = "reboot"
}

data "talos_client_configuration" "this" {
  cluster_name         = var.cluster.name
  client_configuration = talos_machine_secrets.this.client_configuration
  nodes                = local.node_ips
  endpoints            = local.controlplane_ips
}

data "talos_cluster_health" "this" {
  depends_on = [
    talos_machine_bootstrap.this,
  ]
  client_configuration = talos_machine_secrets.this.client_configuration
  control_plane_nodes  = local.controlplane_ips
  worker_nodes         = local.worker_ips
  endpoints            = local.node_ips
  timeouts = {
    read = "10m"
  }
}

resource "talos_cluster_kubeconfig" "this" {
  depends_on = [
    data.talos_cluster_health.this,
  ]
  node                 = local.bootstrap_node_ip
  endpoint             = local.bootstrap_node_ip
  client_configuration = talos_machine_secrets.this.client_configuration
  timeouts = {
    read = "1m"
  }
}
