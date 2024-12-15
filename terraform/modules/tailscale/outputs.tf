output "tailnet_key" {
  value       = tailscale_tailnet_key.this.key
  sensitive   = true
  description = "The tailnet key allows you to register new nodes without needing to sign in via a web browser"
}
