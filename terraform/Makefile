.PHONY: *

all: fmt lint docs

docs:
	@terraform-docs .

fmt:
	@terraform fmt -recursive

lint:
	@tflint --recursive
