include ../Makefile

.ONESHELL:

.PHONY: up
up: apply save bootstrap

.PHONY: apply
apply:
	@terraform apply -auto-approve

.PHONY: bootstrap
bootstrap:
	@bash hack/bootstrap.sh

.PHONY: save
save: save/kubeconfig save/talosconfig

.PHONY: save/kubeconfig
save/kubeconfig:
	@terraform output -raw kubeconfig > outputs/kubeconfig

.PHONY: save/talosconfig
save/talosconfig:
	@terraform output -raw talosconfig > outputs/talosconfig

.PHONY: update-argocd-creds
update-argocd-creds:
	@bash hack/update-argocd-creds.sh

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