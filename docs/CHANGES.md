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

### 1. Build and Installation System
The project now uses a robust build and installation system with the following components:

#### Makefile Updates
```makefile
# Binary name
BINARY_NAME=nepali

# Build directory
BUILD_DIR=bin

# Installation directory
INSTALL_DIR=/usr/local/bin
```

Key features:
- Automated build process
- Proper installation to system directories
- Clean and uninstall capabilities
- Installation verification
- Error handling and checks

#### Build Process
The build process now follows these steps:
1. Creates a `bin` directory if it doesn't exist
2. Builds the binary with proper Go flags
3. Places the binary in the `bin` directory
4. Verifies the build was successful

#### Installation Process
The installation process includes:
1. Building the binary (if not already built)
2. Copying to system directory with proper permissions
3. Verifying the installation
4. Providing clear error messages if something fails

### 2. Git Configuration
Updated `.gitignore` to handle build artifacts properly:
```gitignore
# Allow main binary but ignore other bin contents
bin/*
!bin/.gitkeep
```

Changes include:
- Keeping the `bin` directory in git
- Ignoring build artifacts
- Maintaining a clean repository
- Allowing the main binary to be tracked if needed

### 3. Directory Structure
Added new directories and files:
- `bin/` directory for build outputs
- `.gitkeep` file to maintain directory structure
- Updated documentation in `docs/`

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

### Available Make Commands
- `make build`: Builds the project
- `make install`: Installs the program
- `make uninstall`: Removes the program
- `make clean`: Cleans build artifacts
- `make test`: Runs tests
- `make verify`: Verifies installation
- `make init`: Initializes project structure

### Running the Program
After installation:
```bash
nepali <filename>
```

Example:
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