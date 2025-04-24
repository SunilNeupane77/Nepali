# Recent Changes and Project Structure

## Project Structure
```
.
├── cmd/
│   └── nepali/           # Main executable
│       └── main.go       # Entry point for the interpreter
├── internal/
│   ├── lexer/           # Lexical analyzer
│   ├── parser/          # Parser implementation
│   └── ast/             # Abstract Syntax Tree definitions
├── examples/            # Example programs
├── docs/               # Documentation
├── tests/              # Test files
├── bin/                # Build output directory
├── Makefile           # Build automation
└── go.mod             # Go module definition
```

## Recent Changes

### 1. Lexer Improvements
The lexer has been updated to properly handle Nepali characters:

#### Unicode Support
- Added proper Unicode character handling
- Fixed Nepali keyword recognition
- Improved string literal handling
- Added support for Nepali numbers

#### Token Types
```go
const (
    // Keywords
    VAR      = "VAR"      // संख्या
    PRINT    = "PRINT"    // लेख्नुहोस्
)
```

#### Keyword Mapping
```go
var keywords = map[string]TokenType{
    "संख्या":  VAR,
    "लेख्नुहोस्": PRINT,
}
```

### 2. Testing
Added comprehensive tests for Nepali language features:

#### Test Cases
- Nepali keywords
- Nepali string literals
- Nepali numbers
- Unicode character handling

#### Example Test
```go
func TestNextToken(t *testing.T) {
    input := `संख्या x = ५
लेख्नुहोस्("नमस्ते संसार!")`
    // ... test cases ...
}
```

### 3. Example Programs
Added sample programs to demonstrate language features:

#### Hello World
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

## How to Use

### Building the Project
1. Initialize the project:
```bash
make init
```

2. Build the project:
```bash
make build
```

3. Install the program:
```bash
make install
```

### Running Programs
```bash
nepali examples/hello.nep
```

## Development Guidelines

### Adding New Features
1. Create new packages in the `internal` directory
2. Update imports to use the correct module path
3. Add tests in the corresponding test files
4. Update documentation as needed

### Code Organization
- Keep the main package simple and focused on the entry point
- Place implementation details in the `internal` directory
- Use proper Go package naming conventions
- Follow Go's standard project layout

### Testing
- Write tests for new features
- Run tests using:
```bash
make test
```

## Future Improvements
1. Add proper error handling
2. Implement the parser
3. Add support for more language features
4. Improve documentation
5. Add more examples

## Contributing
Please read [CONTRIBUTING.md](../CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License
This project is licensed under the MIT License - see the [LICENSE](../LICENSE) file for details. 