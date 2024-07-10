.PHONY: help
help: ## Show this help
	@echo "Available makefile targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[93m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build go-factory
	go build . -o ./dist

.PHONY: run
run: ## Run go-factory
	go run .

.PHONY: install
install: ## Install go-factory
	go install github.com/jotadrilo/go-factory
