.PHONY: help
help: ## show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

fmt: ## go format
	gofmt -l -w .

test: ## run test
	go test ./... -v

build: ## go build
	go build -o main ./cmd/checkdeps