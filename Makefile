.PHONY: help
help: ## Show this help
	@echo "Available makefile targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

VERSION := $(shell cat VERSION)

.PHONY: build
build: ## Build go-factory
	go build -ldflags="-X github.com/jotadrilo/go-factory/cmd.version=$(VERSION)" -o ./dist

.PHONY: run
run: ## Run go-factory
	go run -ldflags="-X github.com/jotadrilo/go-factory/cmd.version=$(VERSION)" .

.PHONY: install
install: ## Install go-factory
	go install github.com/jotadrilo/go-factory
