# Makefile for Password Manager

# Variables
BINARY_NAME=password-manager
BUILD_DIR=bin
SOURCE_DIR=.
MAIN_FILE=cmd/main.go
VERSION=1.0.0
BUILD_DATE=$(shell date +%Y-%m-%d_%H:%M:%S)
GIT_HASH=$(shell git rev-parse --short HEAD)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=$(GOCMD) fmt
GOVET=$(GOCMD) vet

# Build flags
LDFLAGS=-ldflags "\
	-X main.Version=$(VERSION) \
	-X main.BuildDate=$(BUILD_DATE) \
	-X main.GitHash=$(GIT_HASH) \
	-w -s"

# Default target
all: build

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BUILD_DIR)/$(BINARY_NAME)"

# Build for development (with debug information)
dev-build:
	@echo "Building for development..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME)-dev $(MAIN_FILE)
	@echo "Development build completed: $(BUILD_DIR)/$(BINARY_NAME)-dev"

# Build for Windows
build-windows:
	@echo "Building for Windows..."
	@mkdir -p $(BUILD_DIR)/windows
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/windows/$(BINARY_NAME).exe $(MAIN_FILE)

# Build for Linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p $(BUILD_DIR)/linux
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/linux/$(BINARY_NAME) $(MAIN_FILE)

# Build for macOS
build-darwin:
	@echo "Building for macOS..."
	@mkdir -p $(BUILD_DIR)/darwin
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o $(BUILD_DIR)/darwin/$(BINARY_NAME) $(MAIN_FILE)

# Build for all platforms
build-all: build-windows build-linux build-darwin
	@echo "Multi-platform build completed"

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GOGET) -d -v ./...
	$(GOMOD) download
	$(GOMOD) tidy

# Run the application
run:
	@echo "Running application..."
	$(GOCMD) run $(MAIN_FILE)

# Run in development mode
dev:
	@echo "Running in development mode..."
	$(GOCMD) run $(MAIN_FILE)

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "Running tests with coverage..."
	$(GOTEST) -cover -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html

# Format code
fmt:
	@echo "Formatting code..."
	$(GOFMT) ./...

# Vet code (static analysis)
vet:
	@echo "Running go vet..."
	$(GOVET) ./...

# Lint code (if golangci-lint is installed)
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed, skipping..."; \
		echo "Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Security audit
audit:
	@echo "Running security audit..."
	$(GOCMD) mod verify
	$(GOVET) ./...
	@if command -v gosec >/dev/null; then \
		gosec ./...; \
	else \
		echo "gosec not installed, skipping..."; \
		echo "Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Clean build artifacts
clean:
	@echo "Cleaning..."
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Install the application
install: build
	@echo "Installing $(BINARY_NAME) to /usr/local/bin..."
	@sudo cp $(BUILD_DIR)/$(BINARY_NAME) /usr/local/bin/
	@echo "Installation completed"

# Uninstall the application
uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstallation completed"

# Create distribution package
dist: build-all
	@echo "Creating distribution packages..."
	@mkdir -p dist
	@cd $(BUILD_DIR)/windows && zip ../../dist/$(BINARY_NAME)-$(VERSION)-windows-amd64.zip $(BINARY_NAME).exe
	@cd $(BUILD_DIR)/linux && tar -czf ../../dist/$(BINARY_NAME)-$(VERSION)-linux-amd64.tar.gz $(BINARY_NAME)
	@cd $(BUILD_DIR)/darwin && tar -czf ../../dist/$(BINARY_NAME)-$(VERSION)-darwin-amd64.tar.gz $(BINARY_NAME)
	@echo "Distribution packages created in dist/ directory"

# Show help
help:
	@echo "Available targets:"
	@echo "  build        - Build the application"
	@echo "  dev-build    - Build for development"
	@echo "  build-all    - Build for all platforms (Windows, Linux, macOS)"
	@echo "  run          - Run the application"
	@echo "  dev          - Run in development mode"
	@echo "  test         - Run tests"
	@echo "  test-coverage - Run tests with coverage report"
	@echo "  fmt          - Format code"
	@echo "  vet          - Run static analysis"
	@echo "  lint         - Run linter (if installed)"
	@echo "  audit        - Run security audit"
	@echo "  deps         - Install dependencies"
	@echo "  clean        - Clean build artifacts"
	@echo "  install      - Install to system"
	@echo "  uninstall    - Uninstall from system"
	@echo "  dist         - Create distribution packages"
	@echo "  help         - Show this help message"

# Default target
.DEFAULT_GOAL := build

# Phony targets
.PHONY: all build dev-build build-windows build-linux build-darwin build-all \
        deps run dev test test-coverage fmt vet lint audit clean install uninstall dist help