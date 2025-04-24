# Nepali Programming Language

[![Go Report Card](https://goreportcard.com/badge/github.com/SunilNeupane77/nepali)](https://goreportcard.com/report/github.com/SunilNeupane77/nepali)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/SunilNeupane77/nepali?status.svg)](https://godoc.org/github.com/SunilNeupane77/nepali)

A programming language designed for Nepali speakers, making coding more accessible and natural in Nepali.

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

- Nepali keywords and syntax
- Support for Unicode characters
- Modern programming constructs
- Easy to learn for Nepali speakers

## Installation

### Prerequisites
- Go 1.21 or later
- Git

### Steps

1. Clone the repository:
```bash
git clone https://github.com/SunilNeupane77/nepali.git
cd nepali
```

2. Install dependencies:
```bash
go mod download
```

3. Build the compiler:
```bash
go build -o nepali cmd/nepali/main.go
```

4. Add to PATH:
```bash
sudo mv nepali /usr/local/bin/
```

## Quick Start

Let's write our first program:

```nepali
# Hello World program
लेख्नुहोस्("नमस्ते संसार!")

# Variable declaration
संख्या x = ५
संख्या y = १०

# Function definition
कार्य जोड्नुहोस्(a, b):
    फिर्ता a + b

# Using the function
जोड = जोड्नुहोस्(x, y)
लेख्नुहोस्(जोड)
```

Save this as `hello.nep` and run it:
```bash
nepali hello.nep
```

## Language Guide

### Basic Syntax

#### Variables and Data Types
```nepali
# Numbers
संख्या x = ५
संख्या y = १०.५

# Strings
पाठ नाम = "नेपाल"

# Lists
सूची = [१, २, ३, ४, ५]

# Dictionaries
शब्दकोश = {
    "नाम": "नेपाल",
    "राजधानी": "काठमाडौं"
}
```

#### Control Structures
```nepali
# If-else
यदि x > ५:
    लेख्नुहोस्("x ५ भन्दा ठूलो छ")
अन्यथा:
    लेख्नुहोस्("x ५ भन्दा सानो छ")

# Loops
# For loop
लागि i मा range(५):
    लेख्नुहोस्(i)

# While loop
जबसम्म x < १०:
    लेख्नुहोस्(x)
    x = x + १
```

#### Functions
```nepali
# Function definition
कार्य जोड्नुहोस्(a, b):
    फिर्ता a + b

# Function with multiple parameters
कार्य परिचय(नाम, उमेर):
    लेख्नुहोस्(f"मेरो नाम {नाम} हो र म {उमेर} वर्षको छु")
```

## Examples

Check the `examples` directory for sample programs:
- `hello.nep`: Basic "Hello World" program
- `python_features.nep`: Python-like features in Nepali
- `calculator.nep`: Simple calculator implementation

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

### Development Commands

```bash
# Initialize project
make init

# Build project
make build

# Install package
make install

# Run tests
make test

# Clean build artifacts
make clean

# Verify installation
make verify
```

## Troubleshooting

### Common Issues

1. **Directory Not Found**
   ```bash
   stat /path/to/nepali/cmd/nepali: directory not found
   make: *** [Makefile:7: install] Error 1
   ```
   Solution:
   ```bash
   make init
   make clean
   make build
   make install
   ```

2. **Installation Fails**
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

## Contributing

We welcome contributions! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

Sunil Neupane - [@SunilNeupane77](https://github.com/SunilNeupane77)

Project Link: [https://github.com/SunilNeupane77/nepali](https://github.com/SunilNeupane77/nepali)

## Acknowledgments

- All contributors who have helped shape this language
- The Nepali programming community
- The Go programming language team 