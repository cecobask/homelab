terraform {
  required_version = "~> 1.9.8"
  required_providers {
    tailscale = {
      source  = "tailscale/tailscale"
      version = "0.17.2"
    }
  }
}
