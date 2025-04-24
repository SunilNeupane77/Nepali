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
	EQ       // ==
	NEQ      // !=
	LT       // <
	GT       // >
	LTE      // <=
	GTE      // >=
	AND      // र
	OR       // वा
	NOT      // होइन

	// Delimiters
	LPAREN    // (
	RPAREN    // )
	LBRACE    // {
	RBRACE    // }
	LBRACKET  // [
	RBRACKET  // ]
	COMMA     // ,
	DOT       // .
	COLON     // :
	SEMICOLON // ;

	// Keywords
	VAR      // संख्या
	PRINT    // लेख्नुहोस्
	IF       // यदि
	ELSE     // अन्यथा
	LOOP     // लूप
	TRUE     // सत्य
	FALSE    // असत्य
	FUNCTION // कार्य
	RETURN   // फिर्ता
	BREAK    // ब्रेक
	CONTINUE // जारी
	IN       // मा
	FOR      // लागि
	WHILE    // जबसम्म
	CLASS    // वर्ग
	NEW      // नयाँ
	THIS     // यो
	SUPER    // माथिल्लो
	IMPORT   // आयात
	FROM     // बाट
	AS       // जस्तो
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
	indentLevel  int  // current indentation level
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
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: EQ, Literal: "==", Line: l.line, Column: l.column}
		} else {
			tok = Token{Type: ASSIGN, Literal: "=", Line: l.line, Column: l.column}
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: NEQ, Literal: "!=", Line: l.line, Column: l.column}
		} else {
			tok = Token{Type: NOT, Literal: "!", Line: l.line, Column: l.column}
		}
	case '<':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: LTE, Literal: "<=", Line: l.line, Column: l.column}
		} else {
			tok = Token{Type: LT, Literal: "<", Line: l.line, Column: l.column}
		}
	case '>':
		if l.peekChar() == '=' {
			l.readChar()
			tok = Token{Type: GTE, Literal: ">=", Line: l.line, Column: l.column}
		} else {
			tok = Token{Type: GT, Literal: ">", Line: l.line, Column: l.column}
		}
	case '+':
		tok = Token{Type: PLUS, Literal: "+", Line: l.line, Column: l.column}
	case '-':
		tok = Token{Type: MINUS, Literal: "-", Line: l.line, Column: l.column}
	case '*':
		tok = Token{Type: MULTIPLY, Literal: "*", Line: l.line, Column: l.column}
	case '/':
		tok = Token{Type: DIVIDE, Literal: "/", Line: l.line, Column: l.column}
	case '(':
		tok = Token{Type: LPAREN, Literal: "(", Line: l.line, Column: l.column}
	case ')':
		tok = Token{Type: RPAREN, Literal: ")", Line: l.line, Column: l.column}
	case '{':
		tok = Token{Type: LBRACE, Literal: "{", Line: l.line, Column: l.column}
	case '}':
		tok = Token{Type: RBRACE, Literal: "}", Line: l.line, Column: l.column}
	case '[':
		tok = Token{Type: LBRACKET, Literal: "[", Line: l.line, Column: l.column}
	case ']':
		tok = Token{Type: RBRACKET, Literal: "]", Line: l.line, Column: l.column}
	case ',':
		tok = Token{Type: COMMA, Literal: ",", Line: l.line, Column: l.column}
	case '.':
		tok = Token{Type: DOT, Literal: ".", Line: l.line, Column: l.column}
	case ':':
		tok = Token{Type: COLON, Literal: ":", Line: l.line, Column: l.column}
	case ';':
		tok = Token{Type: SEMICOLON, Literal: ";", Line: l.line, Column: l.column}
	case '"':
		tok.Literal = l.readString()
		tok.Type = STRING
		tok.Line = l.line
		tok.Column = l.column
		return tok
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

// peekChar looks at the next character without advancing
func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return rune(l.input[l.readPosition])
}

// readString reads a string literal
func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	return l.input[position:l.position]
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
			// Track indentation
			l.indentLevel = 0
		} else if l.ch == ' ' || l.ch == '\t' {
			l.indentLevel++
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
	case "कार्य":
		return FUNCTION
	case "फिर्ता":
		return RETURN
	case "ब्रेक":
		return BREAK
	case "जारी":
		return CONTINUE
	case "मा":
		return IN
	case "लागि":
		return FOR
	case "जबसम्म":
		return WHILE
	case "वर्ग":
		return CLASS
	case "नयाँ":
		return NEW
	case "यो":
		return THIS
	case "माथिल्लो":
		return SUPER
	case "आयात":
		return IMPORT
	case "बाट":
		return FROM
	case "जस्तो":
		return AS
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
