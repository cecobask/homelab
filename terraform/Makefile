.PHONY: *
.SHELLFLAGS := -e -c -o pipefail
.ONESHELL:
TALOS_DIR := $(HOME)/.talos

check: check/fmt check/lint check/validate check/docs

check/fmt:
	@terraform fmt -recursive

check/lint:
	@tflint --recursive

check/validate:
	@terraform validate

check/docs:
	@terraform-docs .

mkdir/talos:
	@mkdir -p $(TALOS_DIR)

save: save/talosconfig save/kubeconfig save/secrets

save/talosconfig: mkdir/talos
	@terraform output -raw talosconfig > $(TALOS_DIR)/config

save/kubeconfig: mkdir/talos
	@terraform output -raw kubeconfig > $(TALOS_DIR)/kubeconfig

save/secrets: mkdir/talos
	@terraform output -raw secrets > $(TALOS_DIR)/secrets.yaml

purge:
	@cd ../automation
	@$(MAKE) build
	@./build/automation tailscale delete-devices --tags=talos
	@./build/automation proxmox stop-vms --tags=talos
	@sleep 5
	@./build/automation proxmox destroy-vms --tags=talos --destroy-unreferenced-disks --purge
	@./build/automation proxmox delete-volumes --storages=local --volumes=local:iso/talos_v1.9.0_nocloud_amd64.iso

unlock:
	@terraform force-unlock -force cecobask/homelab