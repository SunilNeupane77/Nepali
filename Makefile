.PHONY: build install test clean init verify setup

# Build the project
build:
	@echo "Building nepali..."
	@mkdir -p bin
	go build -o bin/nepali cmd/nepali/main.go

# Install the project
install:
	@echo "Installing nepali..."
	@mkdir -p $(shell go env GOPATH)/bin
	go install ./cmd/nepali

# Run tests
test:
	@echo "Running tests..."
	go test ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -rf bin/
	go clean

# Initialize the project
init:
	@echo "Initializing project..."
	go mod init github.com/SunilNeupane77/nepali
	go mod tidy

# Verify installation
verify:
	@echo "Verifying installation..."
	which nepali || echo "nepali not found in PATH"
	nepali --version || echo "nepali command not working"

# Setup development environment
setup:
	@echo "Setting up development environment..."
	@mkdir -p cmd/nepali internal/lexer internal/parser internal/interpreter examples docs tests bin
	@go mod init github.com/SunilNeupane77/nepali
	@go mod tidy
	@echo "Development environment setup complete" 