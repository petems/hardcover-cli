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

## Run performance benchmarks
bench:
	$(GO) test -bench=. -benchmem ./... -run=^$$

## Run specific performance benchmarks for string operations
bench-strings:
	$(GO) test -bench=BenchmarkStringBuilding -benchmem ./cmd -run=^$$

## Run client performance benchmarks
bench-client:
	$(GO) test -bench=. -benchmem ./internal/client -run=^$$

## Run and save benchmark results for comparison
bench-save:
	$(GO) test -bench=. -benchmem ./... -run=^$$ > benchmarks_$(shell date +%Y%m%d_%H%M%S).txt

## Run CPU profiling
profile-cpu:
	$(GO) test -bench=BenchmarkClient_Execute_JSONMarshaling -cpuprofile cpu.prof ./internal/client -run=^$$
	$(GO) tool pprof cpu.prof

## Run memory profiling  
profile-mem:
	$(GO) test -bench=BenchmarkClient_Execute_JSONMarshaling -memprofile mem.prof ./internal/client -run=^$$
	$(GO) tool pprof mem.prof

## Help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'