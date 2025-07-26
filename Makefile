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

.PHONY: all build test lint install clean help release fmt graphql-fetch graphql-generate graphql-update

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

## Fetch latest GraphQL schema from remote endpoint
graphql-fetch:
	@echo "Fetching latest GraphQL schema from $(GRAPHQL_ENDPOINT)..."
	@if [ -z "$(GRAPHQL_ENDPOINT)" ]; then \
		echo "Error: GRAPHQL_ENDPOINT is not set. Please set it to the actual Hardcover API endpoint."; \
		echo "Example: make graphql-fetch GRAPHQL_ENDPOINT=https://api.hardcover.app/graphql"; \
		exit 1; \
	fi
	@echo "Downloading schema from $(GRAPHQL_ENDPOINT)..."
	@echo "Note: This requires the GraphQL endpoint to support introspection queries."
	@if [ -n "$(HARDCOVER_API_KEY)" ]; then \
		echo "Using API key for authentication..."; \
		curl -s -H "Content-Type: application/json" \
			-H "Authorization: Bearer $(HARDCOVER_API_KEY)" \
			-d '{"query":"query IntrospectionQuery{__schema{queryType{name}mutationType{name}subscriptionType{name}types{...FullType}directives{name description locations args{...InputValue}}}}fragment FullType on __Type{kind name description fields(includeDeprecated:true){name description args{...InputValue}type{...TypeRef}isDeprecated deprecationReason}inputFields{...InputValue}interfaces{...TypeRef}enumValues(includeDeprecated:true){name description isDeprecated deprecationReason}possibleTypes{...TypeRef}}fragment InputValue on __InputValue{name description type{...TypeRef}defaultValue}fragment TypeRef on __Type{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name}}}}}}}}"}' \
			$(GRAPHQL_ENDPOINT) > /tmp/schema_response.json; \
	else \
		echo "No API key found. Trying without authentication..."; \
		curl -s -H "Content-Type: application/json" \
			-d '{"query":"query IntrospectionQuery{__schema{queryType{name}mutationType{name}subscriptionType{name}types{...FullType}directives{name description locations args{...InputValue}}}}fragment FullType on __Type{kind name description fields(includeDeprecated:true){name description args{...InputValue}type{...TypeRef}isDeprecated deprecationReason}inputFields{...InputValue}interfaces{...TypeRef}enumValues(includeDeprecated:true){name description isDeprecated deprecationReason}possibleTypes{...TypeRef}}fragment InputValue on __InputValue{name description type{...TypeRef}defaultValue}fragment TypeRef on __Type{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name ofType{kind name}}}}}}}}"}' \
			$(GRAPHQL_ENDPOINT) > /tmp/schema_response.json; \
	fi
	@if [ ! -s /tmp/schema_response.json ]; then \
		echo "Error: Empty response from $(GRAPHQL_ENDPOINT)"; \
		exit 1; \
	fi
	@echo "Converting introspection result to GraphQL SDL..."
	@if command -v jq >/dev/null 2>&1; then \
		jq -r '.data.__schema | to_entries | map(select(.key != "__typename")) | from_entries | {data: {__schema: .}}' /tmp/schema_response.json > /tmp/schema_processed.json; \
	else \
		echo "Warning: jq not found, using raw response"; \
		cp /tmp/schema_response.json /tmp/schema_processed.json; \
	fi
	@if command -v gqlgen-introspect >/dev/null 2>&1; then \
		gqlgen-introspect /tmp/schema_processed.json > $(SCHEMA_FILE); \
	elif command -v go >/dev/null 2>&1; then \
		go run github.com/vektah/gqlparser/v2/cmd/gqlgen-introspect /tmp/schema_processed.json > $(SCHEMA_FILE); \
	else \
		echo "Error: Neither gqlgen-introspect nor go command found"; \
		echo "Please install gqlgen-introspect or ensure go is available"; \
		exit 1; \
	fi
	@rm -f /tmp/schema_response.json /tmp/schema_processed.json
	@echo "Schema updated successfully: $(SCHEMA_FILE)"

## Generate Go code from GraphQL schema and queries
graphql-generate:
	@echo "Generating Go code from GraphQL schema..."
	@if [ ! -f $(GENQLIENT_CONFIG) ]; then \
		echo "Error: genqlient config not found at $(GENQLIENT_CONFIG)"; \
		exit 1; \
	fi
	@echo "Ensuring genqlient dependencies are installed..."
	@$(GO) mod tidy
	@$(GO) get github.com/Khan/genqlient/generate@v0.8.1
	@$(GO) get github.com/vektah/gqlparser/v2/validator@v2.5.19
	$(GENQLIENT) $(GENQLIENT_CONFIG)
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