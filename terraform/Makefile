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

clean/state:
	@terraform state pull | jq '.resources=null | .outputs=null | .check_results=null' > terraform.tfstate.clean
	@terraform state push -force terraform.tfstate.clean
	@rm -f terraform.tfstate.clean
	@terraform force-unlock -force cecobask/homelab || true

clean/resources:
	@cd ../automation
	@$(MAKE) build
	@./build/automation tailscale delete-devices --tags talos
	@./build/automation proxmox stop-vms --node=pve1 --vmids=101,104
	@./build/automation proxmox stop-vms --node=pve2 --vmids=102
	@./build/automation proxmox stop-vms --node=pve3 --vmids=103
	@sleep 5
	@./build/automation proxmox destroy-vms --node=pve1 --vmids=101,104 --destroy-unreferenced-disks --purge
	@./build/automation proxmox destroy-vms --node=pve2 --vmids=102 --destroy-unreferenced-disks --purge
	@./build/automation proxmox destroy-vms --node=pve3 --vmids=103 --destroy-unreferenced-disks --purge
	@./build/automation proxmox delete-volume --node=pve1 --storage=local --volume=local:iso/talos_v1.9.0_nocloud_amd64.iso
	@./build/automation proxmox delete-volume --node=pve2 --storage=local --volume=local:iso/talos_v1.9.0_nocloud_amd64.iso
	@./build/automation proxmox delete-volume --node=pve3 --storage=local --volume=local:iso/talos_v1.9.0_nocloud_amd64.iso
