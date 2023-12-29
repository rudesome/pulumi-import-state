all: compile

target ?= pulumi-import-state

compile: ## local build to binary
	@echo "Compiling..."
	go build -o $(target) cmd/*.go

docker:
	@echo "build docker with nix"
	$(eval OUTPUT=$(shell sh -c "nix build .#docker --json --no-link --print-build-logs | jq -r \".[0].outputs.out\""))
	docker load < ${OUTPUT}
	docker run --env-file .env -it localhost/$(target)

develop:
	nix develop

clean:
	docker image rm localhost/$(target) -f
	docker system prune -f

.PHONY: help
help:  ## this help messages
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
