.PHONY: *

build:
	@go build -o build/automation main.go

lint:
	@golangci-lint run

lint-fix:
	@golangci-lint run --fix
