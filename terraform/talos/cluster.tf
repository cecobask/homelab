locals {
  bootstrap_node_ip = one([for vm_name, vm in var.vms : vm.ipv4 if vm_name == var.cluster.bootstrap_node])
  controlplane_ips  = [for vm_name, vm in var.vms : vm.ipv4 if vm.machine_type == "controlplane"]
  node_ips          = [for vm_name, vm in var.vms : vm.ipv4]
}

resource "talos_machine_secrets" "this" {
  talos_version = var.image.version
}

data "talos_machine_configuration" "this" {
  for_each           = var.vms
  cluster_name       = var.cluster.name
  cluster_endpoint   = format("https://%s:6443", var.cluster.load_balancer_ip)
  talos_version      = var.image.version
  kubernetes_version = var.cluster.kubernetes_version
  machine_type       = each.value.machine_type
  machine_secrets    = talos_machine_secrets.this.machine_secrets
  config_patches = flatten([
    templatefile("config.tftpl", {
      cluster_name     = var.cluster.name
      hostname         = each.key
      installer_url    = data.talos_image_factory_urls.this.urls.installer
      load_balancer_ip = var.cluster.load_balancer_ip
      node_name        = each.value.node_name
    })
  ])
}

resource "talos_machine_configuration_apply" "this" {
  for_each                    = data.talos_machine_configuration.this
  node                        = var.vms[each.key].ipv4
  client_configuration        = talos_machine_secrets.this.client_configuration
  machine_configuration_input = each.value.machine_configuration
  lifecycle {
    replace_triggered_by = [
      proxmox_virtual_environment_vm.this[each.key]
    ]
  }
}

resource "talos_machine_bootstrap" "this" {
  depends_on           = [talos_machine_configuration_apply.this]
  node                 = local.bootstrap_node_ip
  client_configuration = talos_machine_secrets.this.client_configuration
}

data "talos_client_configuration" "this" {
  cluster_name         = var.cluster.name
  client_configuration = talos_machine_secrets.this.client_configuration
  nodes                = local.node_ips
  endpoints            = local.controlplane_ips
}

resource "talos_cluster_kubeconfig" "this" {
  depends_on           = [talos_machine_bootstrap.this]
  node                 = var.cluster.load_balancer_ip
  client_configuration = talos_machine_secrets.this.client_configuration
  timeouts = {
    read = "5m"
  }
}
