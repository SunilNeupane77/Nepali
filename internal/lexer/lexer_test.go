package lexer

import (
	"testing"
)

func TestNew(t *testing.T) {
	input := "संख्या x = ५"
	l := New(input)
	if l == nil {
		t.Fatal("New() returned nil")
	}
}

func TestNextToken(t *testing.T) {
	input := `संख्या x = ५
लेख्नुहोस्("नमस्ते संसार!")`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{VAR, "संख्या"},
		{IDENT, "x"},
		{ASSIGN, "="},
		{INT, "५"},
		{PRINT, "लेख्नुहोस्"},
		{LPAREN, "("},
		{STRING, "नमस्ते संसार!"},
		{RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNepaliKeywords(t *testing.T) {
	input := `संख्या x = ५
लेख्नुहोस्(x)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{VAR, "संख्या"},
		{IDENT, "x"},
		{ASSIGN, "="},
		{INT, "५"},
		{PRINT, "लेख्नुहोस्"},
		{LPAREN, "("},
		{IDENT, "x"},
		{RPAREN, ")"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
