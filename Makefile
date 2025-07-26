# Makefile for hardcover-cli

APP := hardcover-cli
VERSION ?= $(shell git describe --tags --always --dirty)
GO := go
LINT := golangci-lint

BUILD_FLAGS := -ldflags "-X main.Version=${VERSION}"

.PHONY: all build test lint install clean help release fmt

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
	$(LINT) run

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

## Help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'