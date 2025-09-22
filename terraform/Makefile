.PHONY: check
check: check/fmt check/lint check/validate

.PHONY: check/fmt
check/fmt:
	@terraform fmt -recursive

.PHONY: check/lint
check/lint:
	@tflint --recursive

.PHONY: check/validate
check/validate:
	@terraform validate
