GOPATH:=$(shell go env GOPATH)
PWD:=$(shell pwd)
TOOLS:="micromaniacs/tools:v3"

all:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo

.PHONY: install
install: ## Install dependencies
	@if [ ! -f go.mod ]; then go mod init github.com/micromaniacs/gcloud-grpc-boilerplate; fi
	@go mod vendor

.PHONY: upgrade
upgrade:
	@rm go.mod && rm go.sum && rm -rf vendor
	@go mod init github.com/micromaniacs/gcloud-grpc-boilerplate
	@go mod vendor

.PHONY: proto
proto: ## Generate protobuf entities
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) protoc -I . -I /usr/include/google/protobuf --go_out=plugins=grpc:. ./service.proto

.PHONY: mock_message
mock_message:
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) mockgen -package=mocks -source=internal/message/repository.go -destination=internal/message/mocks/repository.go
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) mockgen -package=mocks -source=internal/message/serializer.go -destination=internal/message/mocks/serializer.go
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) mockgen -package=mocks -source=internal/message/validator.go -destination=internal/message/mocks/validator.go

.PHONY: mock
mock: mock_message ## Generate mocks

.PHONY: fmt
fmt: ## Format source code
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) gofumpt -w -s -extra -l .
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) goimports -w -l .

.PHONY: lint
lint: ## Lint source code
	@docker run --rm -v $(PWD):/app -w /app $(TOOLS) golangci-lint --exclude-use-default=false run ./...

.PHONY: test
test: ## Run tests
	@if [ -f coverage.out ]; then rm coverage.out; fi
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -func=coverage.out
	@rm coverage.out
