// Package main is the entry point for the Nepali programming language interpreter
package main

import (
	"fmt"
	"os"

	"github.com/SunilNeupane77/nepali/internal/lexer"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: nepali <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	l := lexer.New(string(content))
	for {
		tok := l.NextToken()
		if tok.Type == lexer.EOF {
			break
		}
		fmt.Printf("Token{Type: %s, Literal: %q}\n", tok.Type, tok.Literal)
	}
}
