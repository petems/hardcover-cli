# Define variables
GO_CMD := go
BIN_NAME := hardcover-cli # Replace with your desired binary name
BUILD_DIR := bin
SRC_DIR := . # Assuming your main package is in the root directory

# Default target
all: build

# Build the application
build:
	@echo "Building $(BIN_NAME)..."
	mkdir -p $(BUILD_DIR)
	$(GO_CMD) build -o $(BUILD_DIR)/$(BIN_NAME) $(SRC_DIR)

# Run the application
run: build
	@echo "Running $(BIN_NAME)..."
	$(BUILD_DIR)/$(BIN_NAME)

# Run tests
test:
	@echo "Running tests..."
	$(GO_CMD) test ./...

# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Format Go code
fmt:
	@echo "Formatting Go code..."
	$(GO_CMD) fmt ./...

# Lint Go code (requires golangci-lint or similar)
lint:
	@echo "Linting Go code..."
	# Install golangci-lint if not present: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

.PHONY: all build run test clean fmt lint