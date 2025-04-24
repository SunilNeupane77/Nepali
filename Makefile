.PHONY: all build clean install uninstall test

# Binary name
BINARY_NAME=nepali

# Build directory
BUILD_DIR=bin

# Installation directory
INSTALL_DIR=/usr/local/bin

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) cmd/nepali/main.go

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@go clean

install: build
	@echo "Installing $(BINARY_NAME)..."
	@if [ -f $(BUILD_DIR)/$(BINARY_NAME) ]; then \
		sudo cp $(BUILD_DIR)/$(BINARY_NAME) $(INSTALL_DIR)/; \
		echo "Installation complete. You can now run '$(BINARY_NAME)' from anywhere."; \
	else \
		echo "Error: Binary not found. Run 'make build' first."; \
		exit 1; \
	fi

uninstall:
	@echo "Uninstalling $(BINARY_NAME)..."
	@sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)
	@echo "Uninstallation complete."

test:
	@echo "Running tests..."
	@go test ./...

# Create bin directory and .gitkeep file
init:
	@mkdir -p $(BUILD_DIR)
	@touch $(BUILD_DIR)/.gitkeep

# Verify installation
verify:
	@if command -v $(BINARY_NAME) >/dev/null 2>&1; then \
		echo "$(BINARY_NAME) is installed and available in PATH"; \
	else \
		echo "$(BINARY_NAME) is not installed or not in PATH"; \
		exit 1; \
	fi 