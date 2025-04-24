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
└── go.mod             # Go module definition
```

## Recent Changes

### 1. Package Structure Fix
- Fixed the package structure to follow Go's standard layout
- Corrected the main package declaration in `cmd/nepali/main.go`
- Aligned import paths with the module name in `go.mod`

### 2. Module Configuration
The project now uses the following module configuration in `go.mod`:
```go
module github.com/SunilNeupane77/nepali

go 1.21

require (
    github.com/stretchr/testify v1.8.4
)

replace github.com/SunilNeupane77/nepali => ./
```

### 3. Main Package Updates
The main package (`cmd/nepali/main.go`) has been updated to:
- Use the correct package name (`package main`)
- Import the lexer package with the correct path
- Provide a simple command-line interface for the interpreter

### 4. Build Process
The project can now be built using:
```bash
go build -o nepali cmd/nepali/main.go
```

## How to Use

### Building the Project
1. Clone the repository:
```bash
git clone https://github.com/SunilNeupane77/nepali.git
cd nepali
```

2. Build the project:
```bash
go build -o nepali cmd/nepali/main.go
```

### Running the Interpreter
```bash
./nepali <filename>
```

Example:
```bash
./nepali examples/hello.nep
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
go test ./...
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