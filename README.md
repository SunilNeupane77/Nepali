# Nepali Programming Language (नेपाली प्रोग्रामिङ भाषा)

[![Go Report Card](https://goreportcard.com/badge/github.com/SunilNeupane77/nepali)](https://goreportcard.com/report/github.com/SunilNeupane77/nepali)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/SunilNeupane77/nepali?status.svg)](https://godoc.org/github.com/SunilNeupane77/nepali)

A programming language designed for Nepali speakers, making coding more accessible and natural in Nepali. यो प्रोग्रामिङ भाषा नेपाली भाषीहरूको लागि डिजाइन गरिएको छ, जसले कोडिङलाई अधिक सजिलो र प्राकृतिक बनाउँछ।

## Table of Contents (सामग्री)

- [Features (विशेषताहरू)](#features)
- [Installation (स्थापना)](#installation)
- [Quick Start (द्रुत सुरुवात)](#quick-start)
- [Language Guide (भाषा गाइड)](#language-guide)
- [Examples (उदाहरणहरू)](#examples)
- [Development (विकास)](#development)
- [Troubleshooting (समस्या समाधान)](#troubleshooting)
- [Contributing (योगदान)](#contributing)
- [License (लाइसेन्स)](#license)
- [Contact (सम्पर्क)](#contact)

## Features (विशेषताहरू)

- Nepali keywords and syntax (नेपाली कीवर्ड र सिन्ट्याक्स)
- Support for Unicode characters (युनिकोड अक्षरहरूको समर्थन)
- Modern programming constructs (आधुनिक प्रोग्रामिङ संरचनाहरू)
- Easy to learn for Nepali speakers (नेपाली भाषीहरूको लागि सजिलो)

## Installation (स्थापना)

### Prerequisites (आवश्यकताहरू)
- Go 1.21 or later (Go 1.21 वा त्यसभन्दा माथिको संस्करण)
- Git

### Steps (चरणहरू)

1. Clone the repository (रेपोजिटरी क्लोन गर्नुहोस्):
```bash
git clone https://github.com/SunilNeupane77/nepali.git
cd nepali
```

2. Install dependencies (डिपेन्डेन्सीहरू स्थापना गर्नुहोस्):
```bash
go mod download
```

3. Build the compiler (कम्पाइलर बिल्ड गर्नुहोस्):
```bash
go build -o nepali cmd/nepali/main.go
```

4. Add to PATH (पाथमा थप्नुहोस्):
```bash
sudo mv nepali /usr/local/bin/
```

## Quick Start (द्रुत सुरुवात)

Let's write our first program in Nepali (चल्नुहोस्, नेपालीमा हाम्रो पहिलो प्रोग्राम लेख्न):

```nepali
# Hello World program (हलो वर्ल्ड प्रोग्राम)
लेख्नुहोस्("नमस्ते संसार!")

# Variable declaration (चर घोषणा)
संख्या x = ५
संख्या y = १०

# Function definition (कार्य परिभाषा)
कार्य जोड्नुहोस्(a, b):
    फिर्ता a + b

# Using the function (कार्य प्रयोग गर्नु)
जोड = जोड्नुहोस्(x, y)
लेख्नुहोस्(जोड)
```

Save this as `hello.nep` and run it (यसलाई `hello.nep` को रूपमा सेभ गर्नुहोस् र चलाउनुहोस्):
```bash
nepali hello.nep
```

## Language Guide (भाषा गाइड)

### Basic Syntax (आधारभूत सिन्ट्याक्स)

#### Variables and Data Types (चर र डाटा प्रकारहरू)
```nepali
# Numbers (संख्याहरू)
संख्या x = ५
संख्या y = १०.५

# Strings (पाठहरू)
पाठ नाम = "नेपाल"

# Lists (सूचीहरू)
सूची = [१, २, ३, ४, ५]

# Dictionaries (शब्दकोशहरू)
शब्दकोश = {
    "नाम": "नेपाल",
    "राजधानी": "काठमाडौं"
}
```

#### Control Structures (नियन्त्रण संरचनाहरू)
```nepali
# If-else (यदि-अन्यथा)
यदि x > ५:
    लेख्नुहोस्("x ५ भन्दा ठूलो छ")
अन्यथा:
    लेख्नुहोस्("x ५ भन्दा सानो छ")

# Loops (लुपहरू)
# For loop (फर लुप)
लागि i मा range(५):
    लेख्नुहोस्(i)

# While loop (वाइल लुप)
जबसम्म x < १०:
    लेख्नुहोस्(x)
    x = x + १
```

#### Functions (कार्यहरू)
```nepali
# Function definition (कार्य परिभाषा)
कार्य जोड्नुहोस्(a, b):
    फिर्ता a + b

# Function with multiple parameters (धेरै पारामिटर भएको कार्य)
कार्य परिचय(नाम, उमेर):
    लेख्नुहोस्(f"मेरो नाम {नाम} हो र म {उमेर} वर्षको छु")
```

## Examples (उदाहरणहरू)

Check the `examples` directory for sample programs (नमुना प्रोग्रामहरूको लागि `examples` डिरेक्टरी हेर्नुहोस्):
- `hello.nep`: Basic "Hello World" program (आधारभूत "Hello World" प्रोग्राम)
- `python_features.nep`: Python-like features in Nepali (नेपालीमा Python जस्ता विशेषताहरू)
- `calculator.nep`: Simple calculator implementation (सरल क्यालकुलेटर कार्यान्वयन)

## Development (विकास)

### Project Structure (प्रोजेक्ट संरचना)
```
.
├── cmd/
│   └── nepali/           # Main executable (मुख्य एक्जिक्युटेबल)
├── internal/
│   ├── lexer/           # Lexical analyzer (लेक्सिकल एनालाइजर)
│   ├── parser/          # Parser implementation (पार्सर कार्यान्वयन)
│   └── interpreter/     # Interpreter/compiler (इन्टरप्रिटर/कम्पाइलर)
├── examples/            # Example programs (उदाहरण प्रोग्रामहरू)
├── docs/               # Documentation (दस्तावेजीकरण)
├── tests/              # Test files (परीक्षण फाइलहरू)
└── Makefile           # Build automation (बिल्ड स्वचालन)
```

### Development Commands (विकास आदेशहरू)

```bash
# Initialize project (प्रोजेक्ट सुरु गर्नु)
make init

# Build project (प्रोजेक्ट बिल्ड गर्नु)
make build

# Install package (प्याकेज स्थापना गर्नु)
make install

# Run tests (परीक्षणहरू चलाउनु)
make test

# Clean build artifacts (बिल्ड आर्टिफ्याक्टहरू सफा गर्नु)
make clean

# Verify installation (स्थापना जाँच गर्नु)
make verify
```

## Troubleshooting (समस्या समाधान)

### Common Issues (सामान्य समस्याहरू)

1. **Directory Not Found (डिरेक्टरी भेटिएन)**
   ```bash
   stat /path/to/nepali/cmd/nepali: directory not found
   make: *** [Makefile:7: install] Error 1
   ```
   Solution (समाधान):
   ```bash
   make init
   make clean
   make build
   make install
   ```

2. **Installation Fails (स्थापना असफल)**
   ```bash
   go: downloading github.com/SunilNeupane77/nepali v0.0.0-20250424065828-ef2424ae4aba
   go: github.com/SunilNeupane77/nepali/cmd/nepali@latest: module github.com/SunilNeupane77/nepali@latest found, but does not contain package
   ```
   Solution (समाधान):
   ```bash
   go mod tidy
   make clean
   make install
   ```

## Contributing (योगदान)

We welcome contributions! Please read our [Contributing Guidelines](CONTRIBUTING.md) for details. हामी योगदानलाई स्वागत गर्छौं! कृपया विवरणहरूको लागि [योगदान दिशानिर्देशहरू](CONTRIBUTING.md) पढ्नुहोस्।

## License (लाइसेन्स)

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. यो प्रोजेक्ट MIT लाइसेन्स अन्तर्गत लाइसेन्स गरिएको छ - विवरणहरूको लागि [LICENSE](LICENSE) फाइल हेर्नुहोस्।

## Contact (सम्पर्क)

Sunil Neupane - [@SunilNeupane77](https://github.com/SunilNeupane77)

Project Link: [https://github.com/SunilNeupane77/nepali](https://github.com/SunilNeupane77/nepali)

## Acknowledgments (आभार)

- All contributors who have helped shape this language (यो भाषालाई आकार दिन मद्दत गर्ने सबै योगदानकर्ताहरू)
- The Nepali programming community (नेपाली प्रोग्रामिङ समुदाय)
- The Go programming language team (Go प्रोग्रामिङ भाषा टिम) 