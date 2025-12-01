#!/bin/bash

set -euo pipefail

PROXMOX_DIR="$(git rev-parse --show-toplevel)/proxmox"
(cd "$PROXMOX_DIR" && make build)
PATH="$PATH:$PROXMOX_DIR/build"

echo "Cleaning up volumes..."
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-bazarr
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-jellyfin
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-prowlarr
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-qbittorrent
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-qui
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-radarr
proxmox volume delete --node=pve2 --storage=local-lvm --volume=vm-9999-sonarr