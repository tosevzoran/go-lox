package golox

import "fmt"

type TokenType int

const (
	// single character tokens
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// one or two character tokens
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// literals
	IDENTIFIER
	STRING
	NUMBER

	// keywords
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line int
}

func NewToken(tokenType TokenType, lexeme string, literal any, line int) Token {
	return Token{tokenType, lexeme, literal, line}
}

func (t Token) String() string {
	return fmt.Sprintf("(%v %v %v)", t.tokenType, t.lexeme, t.literal)
}
