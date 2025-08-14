# TaskMan CLI Makefile

# Build variables
BINARY_NAME=taskman
VERSION=1.0.0
BUILD_DIR=bin
MAIN_PACKAGE=./main.go

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt

# Build flags
LDFLAGS=-ldflags "-X github.com/vkhangstack/taskman/cmd.Version=$(VERSION)"

.PHONY: all build clean test deps fmt lint install uninstall help

# Default target
all: clean deps fmt test build

# Build the binary
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME).${VERSION} $(MAIN_PACKAGE)
	@echo "Built $(BINARY_NAME) v$(VERSION)"

# Build for multiple platforms
build-all: clean deps
	@echo "Building for multiple platforms..."
	@mkdir -p $(BUILD_DIR)

	# Linux AMD64
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-amd64 $(MAIN_PACKAGE)

	# Linux ARM64
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-linux-arm64 $(MAIN_PACKAGE)

	# macOS AMD64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64 $(MAIN_PACKAGE)

	# macOS ARM64 (Apple Silicon)
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 $(MAIN_PACKAGE)

	# Windows AMD64
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe $(MAIN_PACKAGE)

	@echo "Built binaries for all platforms in $(BUILD_DIR)/"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Download dependencies
deps:
	@echo "Downloading dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Install the binary to $GOPATH/bin
install: build
	@echo "Installing $(BINARY_NAME)..."
	$(GOCMD) install $(LDFLAGS) $(MAIN_PACKAGE)
	@echo "$(BINARY_NAME) installed to $(shell go env GOPATH)/bin/$(BINARY_NAME)"

# Uninstall the binary
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@rm -f $(shell go env GOPATH)/bin/$(BINARY_NAME)
	@echo "$(BINARY_NAME) uninstalled"

# Development build (with race detection)
dev:
	@echo "Building development version..."
	$(GOBUILD) -race $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME)-dev $(MAIN_PACKAGE)

# Run the application
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BUILD_DIR)/$(BINARY_NAME)

# Generate documentation
docs:
	@echo "Generating documentation..."
	@mkdir -p docs
	./$(BUILD_DIR)/$(BINARY_NAME) --help > docs/help.txt
	@echo "Documentation generated in docs/"

# Quick development cycle
dev-cycle: fmt test build

# Release preparation
release: clean deps fmt lint test build-all
	@echo "Release $(VERSION) prepared in $(BUILD_DIR)/"

# Show help
help:
	@echo "Available targets:"
	@echo "  all          - Clean, deps, fmt, test, and build"
	@echo "  build        - Build the binary"
	@echo "  build-all    - Build for multiple platforms"
	@echo "  clean        - Remove build artifacts"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  fmt          - Format code"
	@echo "  lint         - Lint code (requires golangci-lint)"
	@echo "  install      - Install binary to GOPATH/bin"
	@echo "  uninstall    - Remove binary from GOPATH/bin"
	@echo "  dev          - Development build with race detection"
	@echo "  run          - Build and run the application"
	@echo "  docs         - Generate documentation"
	@echo "  dev-cycle    - Quick development cycle (fmt, test, build)"
	@echo "  release      - Prepare release (all platforms)"
	@echo "  help         - Show this help message"