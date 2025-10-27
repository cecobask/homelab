include ../Makefile

.ONESHELL:

.PHONY: apply
apply: save
	@terraform apply -auto-approve

.PHONY: destroy
destroy:
	@terraform state rm $$(terraform state list | grep -E '(kubernetes_secret|kustomization_build|kustomization_resource)')
	@terraform destroy -auto-approve

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

.PHONY: unlock
unlock:
	@terraform force-unlock -force cecobask/talos