include ../Makefile

.ONESHELL:

.PHONY: up
up: apply save bootstrap

.PHONY: apply
apply:
	@terraform apply -auto-approve

.PHONY: bootstrap
bootstrap: bootstrap/volumes bootstrap/applications

.PHONY: bootstrap/volumes
bootstrap/volumes:
	@bash hack/bootstrap/volumes.sh

.PHONY: bootstrap/applications
bootstrap/applications:
	@bash hack/bootstrap/applications.sh

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

.PHONY: down
down: cleanup destroy

.PHONY: cleanup
cleanup: cleanup/kubernetes

.PHONY: cleanup/kubernetes
cleanup/kubernetes:
	@bash hack/cleanup/kubernetes.sh

.PHONY: cleanup/proxmox
cleanup/proxmox:
	@bash hack/cleanup/proxmox.sh

.PHONY: destroy
destroy:
	@terraform destroy -auto-approve -refresh=false

.PHONY: unlock
unlock:
	@terraform force-unlock -force cecobask/talos