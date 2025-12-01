#!/bin/bash

set -euo pipefail

PROXMOX_DIR="$(git rev-parse --show-toplevel)/proxmox"
(cd "$PROXMOX_DIR" && make build)
PATH="$PATH:$PROXMOX_DIR/build"

echo "Bootstrapping volumes..."
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-bazarr --size=1G
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-jellyfin --size=10G
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-prowlarr --size=1G
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-qbittorrent --size=1G
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-qui --size=1G
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-radarr --size=5G
proxmox volume create --node=pve2 --storage=local-lvm --volume=vm-9999-sonarr --size=5G
