# Installation Guide

This guide provides detailed instructions for installing and setting up the Nepali Programming Language.

## Prerequisites

- Go 1.21 or higher
- Git
- Make
- VS Code (recommended for development)

## Installation Methods

### Method 1: Using Make (Recommended)

1. **Clone the Repository**
   ```bash
   git clone https://github.com/SunilNeupane77/nepali.git
   cd nepali
   ```

2. **Set Up Development Environment**
   ```bash
   make setup
   ```
   This command will:
   - Create necessary directories
   - Initialize Go module
   - Set up project structure

3. **Build the Project**
   ```bash
   make build
   ```
   This will create the executable in the `bin` directory.

4. **Install Globally**
   ```bash
   make install
   ```
   This will install the `nepali` command globally.

5. **Verify Installation**
   ```bash
   make verify
   ```
   This will check if the installation was successful.

### Method 2: Manual Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/SunilNeupane77/nepali.git
   cd nepali
   ```

2. **Initialize Go Module**
   ```bash
   go mod init github.com/SunilNeupane77/nepali
   go mod tidy
   ```

3. **Build and Install**
   ```bash
   go build -o bin/nepali cmd/nepali/main.go
   go install ./cmd/nepali
   ```

### Method 3: Using Go Install

```bash
go install github.com/SunilNeupane77/nepali/cmd/nepali@latest
```

## Post-Installation Setup

### Adding to PATH

Add Go's bin directory to your PATH:

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

### VS Code Setup

1. Install the Go extension in VS Code
2. Open the project in VS Code:
   ```bash
   code .
   ```
3. Install Go tools when prompted

## Verifying Installation

1. Check if the command is available:
   ```bash
   which nepali
   ```

2. Check the version:
   ```bash
   nepali --version
   ```

3. Run a test program:
   ```bash
   nepali examples/hello.nep
   ```

## Troubleshooting

### Common Issues

1. **Command Not Found**
   ```bash
   nepali: command not found
   ```
   Solution:
   ```bash
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

2. **Directory Not Found**
   ```bash
   stat /path/to/nepali/cmd/nepali: directory not found
   ```
   Solution:
   ```bash
   make setup
   make build
   make install
   ```

3. **Module Not Found**
   ```bash
   go: cannot find module providing package
   ```
   Solution:
   ```bash
   go mod tidy
   make clean
   make install
   ```

### Debugging

1. Check Go environment:
   ```bash
   go env
   ```

2. Verify module status:
   ```bash
   go mod verify
   ```

3. Check installation path:
   ```bash
   echo $GOPATH
   which nepali
   ```

## Development Setup

### Required Tools

1. Install development tools:
   ```bash
   go install golang.org/x/tools/cmd/godoc@latest
   go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
   ```

2. Set up pre-commit hooks:
   ```bash
   cp .git/hooks/pre-commit.sample .git/hooks/pre-commit
   chmod +x .git/hooks/pre-commit
   ```

### Project Structure

```
.
├── cmd/
│   └── nepali/           # Main executable
├── internal/
│   ├── lexer/           # Lexical analyzer
│   ├── parser/          # Parser implementation
│   └── interpreter/     # Interpreter/compiler
├── examples/            # Example programs
├── docs/               # Documentation
├── tests/              # Test files
└── Makefile           # Build automation
```

## Usage

### Basic Usage

1. Create a new file with `.nep` extension:
   ```nepali
   # hello.nep
   संख्या x = ५
   संख्या y = १०
   जोड = x + y
   लेख्नुहोस्(जोड)
   ```

2. Run the program:
   ```bash
   nepali hello.nep
   ```

### Advanced Usage

See the [Language Guide](../README.md#language-guide) for more examples and features.

## Support

If you encounter any issues during installation or usage, please:
1. Check the [Troubleshooting](#troubleshooting) section
2. Open an issue on GitHub
3. Contact the maintainers 