include ../Makefile

.ONESHELL:

.PHONY: cilium
cilium:
	@bash hack/cilium.sh

.PHONY: up
up: apply save

.PHONY: apply
apply:
	@terraform apply -auto-approve

.PHONY: save
save: save/kubeconfig save/talosconfig

.PHONY: save/kubeconfig
save/kubeconfig:
	@terraform output -raw kubeconfig > outputs/kubeconfig

.PHONY: save/talosconfig
save/talosconfig:
	@terraform output -raw talosconfig > outputs/talosconfig

.PHONY: down
down: cleanup destroy

.PHONY: cleanup
cleanup:
	@bash hack/cleanup.sh

.PHONY: destroy
destroy:
	@terraform destroy -auto-approve -refresh=false

.PHONY: unlock
unlock:
	@terraform force-unlock -force cecobask/talos
