# नेपाली प्रोग्रामिङ भाषा (Nepali Programming Language)

[![Go Report Card](https://goreportcard.com/badge/github.com/SunilNeupane77/nepali)](https://goreportcard.com/report/github.com/SunilNeupane77/nepali)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/SunilNeupane77/nepali?status.svg)](https://godoc.org/github.com/SunilNeupane77/nepali)

A Python-like programming language implementation in Go that allows writing code in Nepali script. This project aims to make programming more accessible to Nepali speakers by providing a familiar syntax in their native language.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Language Guide](#language-guide)
- [Examples](#examples)
- [Development](#development)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Features

- Python-like syntax with Nepali keywords
- Indentation-based block structure
- Support for functions and classes
- List comprehensions and dictionary support
- String formatting and methods
- Error handling with try-except
- Native Nepali number system support
- Modern programming constructs
- Advanced features like generators, decorators, and async programming

## Installation

### Prerequisites

- Go 1.21 or higher
- Git
- Make (optional, for build automation)

### From Source

1. Clone the repository:
```bash
git clone https://github.com/SunilNeupane77/nepali.git
cd nepali
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build and install:
```bash
make install
```

### Using Go Install

```bash
go install github.com/SunilNeupane77/nepali/cmd/nepali@latest
```

### Verifying Installation

After installation, verify that the `nepali` command is available:
```bash
nepali --version
```

## Quick Start

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

## Language Guide

### Basic Syntax

#### Variables and Data Types
```nepali
संख्या x = ५  # Number variable
पाठ नाम = "नेपाल"  # String variable
सूची = [१, २, ३]  # List
शब्दकोश = {"नाम": "नेपाल"}  # Dictionary
```

#### Functions
```nepali
कार्य जोड्नुहोस्(a, b):
    फिर्ता a + b
```

#### Classes
```nepali
वर्ग व्यक्ति:
    कार्य __init__(यो, नाम):
        यो.नाम = नाम
    
    कार्य परिचय(यो):
        लेख्नुहोस्(f"मेरो नाम {यो.नाम} हो")
```

### Advanced Features

#### Generators
```nepali
कार्य संख्या_जेनेरेटर(सुरु, अन्त):
    संख्या i = सुरु
    जबसम्म i < अन्त:
        फिर्ता i
        i = i + १
```

#### Decorators
```nepali
कार्य टाइमर(कार्य):
    कार्य wrapper(*args, **kwargs):
        सुरु_समय = time.time()
        परिणाम = कार्य(*args, **kwargs)
        अन्त_समय = time.time()
        लेख्नुहोस्(f"कार्य सम्पन्न भयो: {अन्त_समय - सुरु_समय} सेकेन्डमा")
        फिर्ता परिणाम
    फिर्ता wrapper
```

#### Async Programming
```nepali
कार्य async_कार्य():
    लेख्नुहोस्("एसिन्क्रोनस कार्य सुरु भयो")
    await asyncio.sleep(१)
    लेख्नुहोस्("एसिन्क्रोनस कार्य सम्पन्न भयो")
```

## Examples

Check out the `examples` directory for sample programs:
- `hello.nep`: Basic variable operations
- `student_management.nep`: Object-oriented programming example
- `advanced_features.nep`: Advanced language features demonstration

## Development

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

### Building from Source

```bash
make build
```

### Running Tests

```bash
make test
```

### Cleaning Build Artifacts

```bash
make clean
```

### Development Setup

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

## Troubleshooting

### Common Issues

1. **Installation Fails**
   ```bash
   go: downloading github.com/SunilNeupane77/nepali v0.0.0-20250424065828-ef2424ae4aba
   go: github.com/SunilNeupane77/nepali/cmd/nepali@latest: module github.com/SunilNeupane77/nepali@latest found, but does not contain package
   ```
   Solution:
   ```bash
   go mod tidy
   make clean
   make install
   ```

2. **Command Not Found**
   ```bash
   nepali: command not found
   ```
   Solution: Ensure your Go bin directory is in your PATH:
   ```bash
   echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
   source ~/.bashrc
   ```

3. **Build Errors**
   ```bash
   go build: no Go files in /path/to/nepali
   ```
   Solution: Ensure you're in the correct directory and all files are present:
   ```bash
   ls -la cmd/nepali/
   go mod tidy
   ```

### Debugging

1. Enable verbose output:
```bash
go build -v ./cmd/nepali
```

2. Check module status:
```bash
go mod verify
```

3. View module graph:
```bash
go mod graph
```

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and coding standards
- Write tests for new features
- Update documentation for any changes
- Use meaningful commit messages
- Keep the code clean and well-documented

### Code Style

- Use `gofmt` for code formatting
- Run `golangci-lint` for code quality checks
- Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Sunil Neupane - [@SunilNeupane77](https://github.com/SunilNeupane77)

Project Link: [https://github.com/SunilNeupane77/nepali](https://github.com/SunilNeupane77/nepali)

## Acknowledgments

- Inspired by Python's clean and readable syntax
- Thanks to all contributors who help make programming more accessible
- Special thanks to the Go community for their excellent tools and libraries 