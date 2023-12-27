all: compile

target ?= pulumi-import-state

compile: ## local build to binary
	@echo "Compiling..."
	go build -o $(target) cmd/*.go

docker:
	@echo "build docker with nix"
	nix build .#docker --json --no-link --print-build-logs

.PHONY: help
help:  ## this help messages
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
