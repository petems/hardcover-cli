# Makefile for hardcover-cli

APP := hardcover-cli
VERSION ?= $(shell git describe --tags --always --dirty)
GO := go
LINT := golangci-lint

BUILD_FLAGS := -ldflags "-X main.Version=${VERSION}"

.PHONY: all build test lint install clean help release fmt generate-types

all: build

## Build the binary
build:
	@echo "Building $(APP) version $(VERSION)"
	$(GO) build -o bin/$(APP) $(BUILD_FLAGS) .

## Run tests
test:
	$(GO) test ./... -timeout 2m

## Run linter
lint:
	@echo "Checking golangci-lint version..."
	@$(LINT) version | grep -q "golangci-lint has version" || (echo "golangci-lint not found. Please install it first." && exit 1)
	@$(LINT) version | grep -oE "version [0-9]+\.[0-9]+\.[0-9]+" | cut -d' ' -f2 | awk -F. '{if ($$1 > 2 || ($$1 == 2 && $$2 >= 3)) exit 0; else exit 1}' || (echo "golangci-lint version 2.3.0 or higher required. Current version:" && $(LINT) version && exit 1)
	$(LINT) run

lint-fix:
	$(LINT) run --fix

## Format code
fmt:
	$(GO) fmt ./...

## Install binary into $GOBIN or default $GOPATH/bin
install: build
	@echo "Installing $(APP)..."
	install -m 0755 bin/$(APP) $(GOBIN) || install -m 0755 bin/$(APP) $(GOPATH)/bin

## Remove binaries and caches
clean:
	@echo "Cleaning..."
	rm -rf bin/
	$(GO) clean -cache ./...

## Generate Go types from GraphQL schema
generate-types: ## Generate Go types from remote GraphQL schema
	@echo "Generating Go types from GraphQL schema..."
	@echo "Generation started at: $(shell date '+%Y-%m-%d %H:%M:%S %Z')"
	@go run scripts/generate-types.go
	@echo "Generation completed at: $(shell date '+%Y-%m-%d %H:%M:%S %Z')"

## Help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'