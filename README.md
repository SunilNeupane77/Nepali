# नेपाली प्रोग्रामिङ भाषा (Nepali Programming Language)

A programming language implementation in Go that allows writing code in Nepali script. This project aims to make programming more accessible to Nepali speakers by providing a familiar syntax in their native language.

## Features

- Lexical analysis of Nepali text
- Basic syntax parsing
- Support for Nepali keywords and operators
- Simple variable declarations and arithmetic operations
- Support for control structures (if-else, loops)
- Native Nepali number system support

## Prerequisites

- Go 1.16 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/SunilNeupane77/nepali.git
cd nepali
```

2. Install the package:
```bash
go install github.com/SunilNeupane77/nepali/cmd/nepali
```

3. Verify installation:
```bash
nepali --version
```

## Usage

### Writing Your First Program

Create a new file with `.nep` extension, for example `hello.nep`:

```nepali
# एक सरल कार्यक्रम
संख्या x = ५
संख्या y = १०
जोड = x + y
लेख्नुहोस्(जोड)
```

### Running Programs

To run a Nepali program:

```bash
nepali run hello.nep
```

### Language Syntax

#### Variables
```nepali
संख्या x = ५  # Number variable
पाठ नाम = "नेपाल"  # String variable
```

#### Control Structures
```nepali
# If-Else
यदि (x > y) {
    लेख्नुहोस्("x ठूलो छ")
} अन्यथा {
    लेख्नुहोस्("y ठूलो छ")
}

# Loops
लूप (i <= ५) {
    लेख्नुहोस्(i)
    i = i + १
}
```

#### Operators
- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Assignment: `=`

### Examples

Check out the `examples` directory for more sample programs:
- `hello.nep`: Basic variable operations
- `loop.nep`: Loop examples
- `conditional.nep`: If-else examples

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
└── docs/               # Documentation
```

### Running Tests
```bash
go test ./...
```

### Building from Source
```bash
go build -o nepali cmd/nepali/main.go
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by the need for programming languages in native scripts
- Thanks to all contributors who help make programming more accessible

## Contact

Sunil Neupane - [@SunilNeupane77](https://github.com/SunilNeupane77)

Project Link: [https://github.com/SunilNeupane77/nepali](https://github.com/SunilNeupane77/nepali) 