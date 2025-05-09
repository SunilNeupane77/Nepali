package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SunilNeupane77/nepali/internal/evaluator"
	"github.com/SunilNeupane77/nepali/internal/lexer"
	"github.com/SunilNeupane77/nepali/internal/parser"
	"github.com/SunilNeupane77/nepali/internal/object"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		source, err := os.ReadFile(filename)
		if err != nil {
			fmt.Printf("Error reading file: %s\n", err)
			os.Exit(1)
		}
		run(string(source))
	} else {
		fmt.Printf("नेपाली प्रोग्रामिङ भाषा\n")
		fmt.Printf("त्याहाँ लाइन टाइप गर्नुहोस् `अन्त्य` लाइन टाइप गर्नुहोस्\n")
		fmt.Printf("> ")
		
		for {
			var input string
			if _, err := fmt.Scanln(&input); err != nil {
				fmt.Printf("Error reading input: %s\n", err)
				continue
			}
			
			if strings.TrimSpace(input) == "अन्त्य" {
				break
			}
			
			run(input)
			fmt.Printf("> ")
		}
	}
}

func run(source string) {
	l := lexer.New(source)
	p := parser.New(l)
	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		printParserErrors(p.Errors())
		os.Exit(1)
	}

	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		fmt.Printf("%s\n", evaluated.Inspect())
	}
}

func printParserErrors(errors []string) {
	fmt.Printf("कोडमा त्रुटि छन्:\n")
	for _, msg := range errors {
		fmt.Printf("\t%s\n", msg)
	}
}
