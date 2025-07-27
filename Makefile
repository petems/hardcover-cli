# Makefile for hardcover-cli

APP := hardcover-cli
VERSION ?= $(shell git describe --tags --always --dirty)
GO := go
LINT := golangci-lint
GENQLIENT := go run github.com/Khan/genqlient

BUILD_FLAGS := -ldflags "-X main.Version=${VERSION}"

# GraphQL configuration
GRAPHQL_DIR := internal/client
SCHEMA_FILE := $(GRAPHQL_DIR)/schema.graphql
GENQLIENT_CONFIG := $(GRAPHQL_DIR)/genqlient.yaml
GENERATED_FILE := $(GRAPHQL_DIR)/generated.go

# Default GraphQL endpoint (update this with the actual Hardcover API endpoint)
GRAPHQL_ENDPOINT ?= https://api.hardcover.app/v1/graphql

# Use latest version of genqlient
GENQLIENT_LATEST := go run github.com/Khan/genqlient@latest

.PHONY: all build build-dev build-prod test lint install clean help release fmt graphql-fetch graphql-generate graphql-update

all: build-prod

## Build the binary (production build - no development features)
build: build-prod

## Build the binary with development features (includes schema command and generated code)
build-dev:
	@echo "Building $(APP) version $(VERSION) with development features"
	$(GO) build -tags dev -o bin/$(APP)-dev $(BUILD_FLAGS) .

## Build the binary for production (excludes development features)
build-prod:
	@echo "Building $(APP) version $(VERSION) for production"
	$(GO) build -o bin/$(APP) $(BUILD_FLAGS) .

## Run tests
test:
	$(GO) test ./... -timeout 2m

## Run linter
lint:
	$(LINT) run

## Format code
fmt:
	$(GO) fmt ./...

## Install binary into $GOBIN or default $GOPATH/bin
install: build-prod
	@echo "Installing $(APP)..."
	install -m 0755 bin/$(APP) $(GOBIN) || install -m 0755 bin/$(APP) $(GOPATH)/bin

## Remove binaries and caches
clean:
	@echo "Cleaning..."
	rm -rf bin/
	$(GO) clean -cache ./...

## Fetch latest GraphQL schema from remote endpoint using the CLI
graphql-fetch:
	@echo "Fetching latest GraphQL schema using hardcover-cli..."
	@if [ ! -f bin/$(APP)-dev ]; then \
		echo "Building hardcover-cli with dev features first..."; \
		$(MAKE) build-dev; \
	fi
	@bin/$(APP)-dev schema --output $(SCHEMA_FILE)
	@echo "Schema updated successfully: $(SCHEMA_FILE)"

## Generate Go code from GraphQL schema and queries using latest version
graphql-generate:
	@echo "Generating Go code from GraphQL schema using latest version of genqlient..."
	@if [ ! -f $(GENQLIENT_CONFIG) ]; then \
		echo "Error: genqlient config not found at $(GENQLIENT_CONFIG)"; \
		exit 1; \
	fi
	@if [ ! -f $(SCHEMA_FILE) ]; then \
		echo "Error: Schema file not found at $(SCHEMA_FILE)"; \
		echo "Run 'make graphql-fetch' first to download the schema"; \
		exit 1; \
	fi
	$(GENQLIENT_LATEST) $(GENQLIENT_CONFIG)
	@echo "Generated Go code: $(GENERATED_FILE)"

## Fetch latest schema and regenerate code (complete update)
graphql-update: graphql-fetch graphql-generate
	@echo "GraphQL schema and code generation completed successfully!"

## Show GraphQL schema information
graphql-info:
	@echo "GraphQL Configuration:"
	@echo "  Schema file: $(SCHEMA_FILE)"
	@echo "  Config file: $(GENQLIENT_CONFIG)"
	@echo "  Generated file: $(GENERATED_FILE)"
	@echo "  Endpoint: $(GRAPHQL_ENDPOINT)"
	@echo ""
	@if [ -f $(SCHEMA_FILE) ]; then \
		echo "Current schema types:"; \
		grep "^type " $(SCHEMA_FILE) | sed 's/^type /  - /' | sed 's/ {.*//'; \
	else \
		echo "No schema file found at $(SCHEMA_FILE)"; \
	fi

## Help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'