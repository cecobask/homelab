.PHONY: *

TALOS_DIR=$(HOME)/.talos

mkdir-talos:
	@mkdir -p $(TALOS_DIR)

docs:
	@terraform-docs .

fmt:
	@terraform fmt -recursive

lint:
	@tflint --recursive

all-checks: fmt lint docs

gen-talosconfig: mkdir-talos
	@terraform output -raw talosconfig > $(TALOS_DIR)/config

gen-kubeconfig: mkdir-talos
	@terraform output -raw kubeconfig > $(TALOS_DIR)/kubeconfig

gen-secrets: mkdir-talos
	@terraform output -raw secrets > $(TALOS_DIR)/secrets.yaml

all-configs: gen-talosconfig gen-kubeconfig gen-secrets

clean-state:
	@terraform state pull | jq '.serial=(.serial+1) | .resources=[] | .outputs={}' > terraform.tfstate.purged
	@terraform state push -lock=false terraform.tfstate.purged
	@rm -f terraform.tfstate.purged
	@terraform force-unlock -force cecobask/homelab