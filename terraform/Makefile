.PHONY: *

all-checks: fmt lint docs

docs:
	@terraform-docs .

fmt:
	@terraform fmt -recursive

lint:
	@tflint --recursive

clean-state:
	@terraform state pull | jq '.serial=(.serial+1) | .resources=[] | .outputs={}' > terraform.tfstate.purged
	@terraform state push -lock=false terraform.tfstate.purged
	@rm -f terraform.tfstate.purged
	@terraform force-unlock -force cecobask/homelab