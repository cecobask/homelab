output "talosconfig" {
  value     = data.talos_client_configuration.this.talos_config
  sensitive = true
}

output "kubeconfig" {
  value     = talos_cluster_kubeconfig.this.kubeconfig_raw
  sensitive = true
}

output "secrets" {
  value     = yamlencode(talos_machine_secrets.this.machine_secrets)
  sensitive = true
}
