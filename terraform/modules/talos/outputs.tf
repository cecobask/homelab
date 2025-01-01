output "talosconfig" {
  value       = data.talos_client_configuration.this.talos_config
  description = "Talos configuration file for usage with talosctl"
  sensitive   = true
}

output "kubeconfig" {
  value       = talos_cluster_kubeconfig.this.kubeconfig_raw
  description = "Admin Kubernetes configuration file"
  sensitive   = true
}

output "secrets" {
  value       = yamlencode(talos_machine_secrets.this.machine_secrets)
  description = "Secrets bundle file which can later be used to generate a config"
  sensitive   = true
}
