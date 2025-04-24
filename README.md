# नेपाली प्रोग्रामिङ भाषा (Nepali Programming Language)

A Python-like programming language implementation in Go that allows writing code in Nepali script. This project aims to make programming more accessible to Nepali speakers by providing a familiar syntax in their native language.

## Features

- Python-like syntax with Nepali keywords
- Indentation-based block structure
- Support for functions and classes
- List comprehensions and dictionary support
- String formatting and methods
- Error handling with try-except
- Native Nepali number system support
- Modern programming constructs

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

#### Control Structures
```nepali
# If-Else
यदि (x > y):
    लेख्नुहोस्("x ठूलो छ")
अन्यथा:
    लेख्नुहोस्("y ठूलो छ")

# For Loop
लागि i मा range(५):
    लेख्नुहोस्(i)

# While Loop
जबसम्म i < ५:
    लेख्नुहोस्(i)
    i = i + १
```

#### List Comprehensions
```nepali
संख्या सूची = [x लागि x मा range(१, ११)]
संख्या वर्ग = [x * x लागि x मा सूची]
```

#### String Formatting
```nepali
नाम = "नेपाल"
लेख्नुहोस्(f"मेरो देश {नाम} हो")
```

### Examples

Check out the `examples` directory for more sample programs:
- `hello.nep`: Basic variable operations
- `python_features.nep`: Python-like features demonstration
- `classes.nep`: Object-oriented programming examples

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

- Inspired by Python's clean and readable syntax
- Thanks to all contributors who help make programming more accessible

## Contact

Sunil Neupane - [@SunilNeupane77](https://github.com/SunilNeupane77)

Project Link: [https://github.com/SunilNeupane77/nepali](https://github.com/SunilNeupane77/nepali) 