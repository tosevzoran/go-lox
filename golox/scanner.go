package golox

import (
	"errors"
	"strconv"
)


var Keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}


type Scanner struct {
	source  string
	tokens  []Token
	start   int
	current int
	line    int
}

func NewScanner(s string) Scanner {
	return Scanner{
		source:  s,
		tokens:  make([]Token, 0),
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() ([]Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()

		if err != nil {
			return nil, err
		}
	}

	s.tokens = append(s.tokens, NewToken(EOF, "", nil, s.line))

	return s.tokens, nil
}

func (s *Scanner) scanToken() error {
	char := s.advance()
	switch char {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)

	// check for the second characters
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			// a comment goes until the end of the line
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
		case ' ', '\r', '\t':
			// ignore whitespace
		case '\n':
			s.line++
		case '"':
			err := s.string()

			if err != nil {
				return err
			}
		default:
			if isDigit(char) {
				err := s.number()

				if err != nil {
					return err
				}
			} else if isAlpha(char) {
				s.identifier()
			} else {
				return errors.New("unexpected character")
			}
	}

	return nil
}

func (s *Scanner) addToken(tokenType TokenType) {
	text := s.source[s.start:s.current]
	s.addTokenWithLiteral(tokenType, text, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, lexeme string, literal any) {
	s.tokens = append(s.tokens, NewToken(tokenType, lexeme, literal, s.line))
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

/* match is like a conditional advance - it consumes the next char only if matches the expected value */
func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}

	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func (s *Scanner) string() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.isAtEnd() {
		return errors.New("unterminated string")
	}

	// the closing "
	s.advance()

	// trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]

	s.addTokenWithLiteral(STRING, value, value)

	return nil
}

func (s *Scanner) number() error {
	for isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && isDigit(s.peekNext()) {
		// consume the "."
		s.advance()

		for isDigit(s.peek()) {
			s.advance()
		}
	}

	stringVal := s.source[s.start:s.current]

	value, err := strconv.ParseFloat(stringVal, 32)

	if err != nil {
		return err
	}

	s.addTokenWithLiteral(NUMBER, stringVal, value)

	return nil
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	text := s.source[s.start:s.current]
	tokenType := Keywords[text]

	if tokenType == 0 {
		tokenType = IDENTIFIER
	}

	s.addTokenWithLiteral(tokenType, text, text)
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isAlphaNumeric(char byte) bool {
	return isDigit(char) || isAlpha(char)
}
