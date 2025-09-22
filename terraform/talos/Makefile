include ../Makefile

.ONESHELL:

TALOS_DIR := $(HOME)/.talos
$(TALOS_DIR):
	@mkdir -p "$@"

.PHONY: save
save: save/talosconfig save/kubeconfig save/secrets

.PHONY: save/talosconfig
save/talosconfig: | $(TALOS_DIR)
	@terraform output -raw talosconfig > $(TALOS_DIR)/config

.PHONY: save/kubeconfig
save/kubeconfig: | $(TALOS_DIR)
	@terraform output -raw kubeconfig > $(TALOS_DIR)/kubeconfig

.PHONY: save/secrets
save/secrets: | $(TALOS_DIR)
	@terraform output -raw secrets > $(TALOS_DIR)/secrets.yaml

.PHONY: purge
purge:
	@cd ../../automation
	@$(MAKE) build
	@./build/automation proxmox stop-vms --tags=talos
	@sleep 5
	@./build/automation proxmox destroy-vms --tags=talos --destroy-unreferenced-disks --purge
	@./build/automation proxmox delete-volumes --storages=local --volumes=local:iso/talos_v1.11.0_nocloud_amd64.iso

.PHONY: unlock
unlock:
	@terraform force-unlock -force cecobask/talos