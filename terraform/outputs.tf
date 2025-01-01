output "talosconfig" {
  value       = module.talos.talosconfig
  description = "Talos configuration file for usage with talosctl"
  sensitive   = true
}

output "kubeconfig" {
  value       = module.talos.kubeconfig
  description = "Admin Kubernetes configuration file"
  sensitive   = true
}

output "secrets" {
  value       = module.talos.secrets
  description = "Secrets bundle file which can later be used to generate a config"
  sensitive   = true
}
