include ../Makefile

.ONESHELL:

.PHONY: up
up: apply save bootstrap

.PHONY: apply
apply:
	@terraform apply -auto-approve

.PHONY: bootstrap
bootstrap:
	@kubectl apply -k bootstrap

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
cleanup:
	@kubens argocd
	@argocd login --core
	@argocd app delete media --yes --wait

.PHONY: destroy
destroy:
	@terraform destroy -auto-approve -refresh=false

.PHONY: unlock
unlock:
	@terraform force-unlock -force cecobask/talos