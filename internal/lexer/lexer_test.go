package lexer

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `संख्या x = ५
लेख्नुहोस्(x)`

	tests := []struct {
		expectedType    TokenType
		expectedLiteral string
	}{
		{VAR, "संख्या"},
		{IDENT, "x"},
		{ASSIGN, "="},
		{NUMBER, "५"},
		{PRINT, "लेख्नुहोस्"},
		{IDENT, "x"},
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
