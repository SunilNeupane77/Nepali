package lexer

import (
	"unicode"
)

// TokenType represents the type of token
type TokenType int

const (
	// Special tokens
	EOF TokenType = iota
	ILLEGAL

	// Identifiers and literals
	IDENT  // variable names
	NUMBER // numbers
	STRING // strings

	// Operators
	PLUS     // +
	MINUS    // -
	MULTIPLY // *
	DIVIDE   // /
	ASSIGN   // =

	// Keywords
	VAR   // संख्या
	PRINT // लेख्नुहोस्
	IF    // यदि
	ELSE  // अन्यथा
	LOOP  // लूप
	TRUE  // सत्य
	FALSE // असत्य
)

// Token represents a lexical token
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}

// Lexer represents the lexical analyzer
type Lexer struct {
	input        string
	position     int  // current position in input
	readPosition int  // current reading position in input
	ch           rune // current char under examination
	line         int  // current line number
	column       int  // current column number
}

// New creates a new Lexer
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

// readChar reads the next character and advances the position
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = rune(l.input[l.readPosition])
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++
}

// NextToken returns the next token
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = Token{Type: ASSIGN, Literal: "=", Line: l.line, Column: l.column}
	case '+':
		tok = Token{Type: PLUS, Literal: "+", Line: l.line, Column: l.column}
	case '-':
		tok = Token{Type: MINUS, Literal: "-", Line: l.line, Column: l.column}
	case '*':
		tok = Token{Type: MULTIPLY, Literal: "*", Line: l.line, Column: l.column}
	case '/':
		tok = Token{Type: DIVIDE, Literal: "/", Line: l.line, Column: l.column}
	case 0:
		tok = Token{Type: EOF, Literal: "", Line: l.line, Column: l.column}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.lookupIdentifier(tok.Literal)
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else if isDigit(l.ch) {
			tok.Type = NUMBER
			tok.Literal = l.readNumber()
			tok.Line = l.line
			tok.Column = l.column
			return tok
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Line: l.line, Column: l.column}
		}
	}

	l.readChar()
	return tok
}

// readIdentifier reads an identifier
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a number
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.column = 0
		}
		l.readChar()
	}
}

// lookupIdentifier looks up the identifier to determine if it's a keyword
func (l *Lexer) lookupIdentifier(ident string) TokenType {
	switch ident {
	case "संख्या":
		return VAR
	case "लेख्नुहोस्":
		return PRINT
	case "यदि":
		return IF
	case "अन्यथा":
		return ELSE
	case "लूप":
		return LOOP
	case "सत्य":
		return TRUE
	case "असत्य":
		return FALSE
	default:
		return IDENT
	}
}

// isLetter checks if the character is a letter
func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

// isDigit checks if the character is a digit
func isDigit(ch rune) bool {
	return unicode.IsDigit(ch)
}
