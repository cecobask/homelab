output "talosconfig" {
  value     = module.talos.talosconfig
  sensitive = true
}

output "kubeconfig" {
  value     = module.talos.kubeconfig
  sensitive = true
}

output "secrets" {
  value     = module.talos.secrets
  sensitive = true
}