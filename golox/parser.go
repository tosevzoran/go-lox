package golox

import "fmt"

/*
expression     → equality ;
equality       → comparison ( ( "!=" | "==" ) comparison )* ;
comparison     → term ( ( ">" | ">=" | "<" | "<=" ) term )* ;
term           → factor ( ( "-" | "+" ) factor )* ;
factor         → unary ( ( "/" | "*" ) unary )* ;
unary          → ( "!" | "-" ) unary | primary ;
primary        → NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")" ;


The parser is represented with a struce

Each rule becomes a function
*/

type Parser struct {
	tokens  []Token
	current int
}

func NewParser(t []Token) Parser {
	return Parser{t, 0}
}

func (p *Parser) parse() (IExpr, error) {
	return p.expression()
}

func (p *Parser) expression() (IExpr, error) {
	return p.equality()
}

func (p *Parser) equality() (IExpr, error) {
	expr, err := p.comparison()

	if err != nil {
		return nil, err
	}

	for p.match(BANG_EQUAL, EQUAL_EQUAL) {
		operator := p.prevoius()
		right, err := p.comparison()

		if err != nil {
			return nil, err
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) comparison() (IExpr, error) {
	expr, err := p.term()

	if err != nil {
		return nil, err
	}

	for p.match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
		operator := p.prevoius()
		right, err := p.term()

		if err != nil {
			return nil, err
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) term() (IExpr, error) {
	expr, err := p.factor()

	if err != nil {
		return nil, err
	}

	for p.match(MINUS, PLUS) {
		operator := p.prevoius()
		right, err := p.factor()

		if err != nil {
			return nil, err
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, err
}

func (p *Parser) factor() (IExpr, error) {
	expr, err := p.unary()

	if err != nil {
		return nil, err
	}

	for p.match(SLASH, STAR) {
		operator := p.prevoius()
		right, err := p.unary()

		if err != nil {
			return nil, err
		}

		expr = NewBinaryExpr(expr, operator, right)
	}

	return expr, nil
}

func (p *Parser) unary() (IExpr, error) {
	if p.match(BANG, BANG_EQUAL) {
		operator := p.prevoius()
		right, err := p.unary()

		if err != nil {
			return nil, err
		}

		return NewUnaryExpr(operator, right), nil
	}

	return p.primary()
}

func (p *Parser) primary() (IExpr, error) {
	if p.match(FALSE) {
		return NewLiteralExpr(false), nil
	}

	if p.match(TRUE) {
		return NewLiteralExpr(true), nil
	}

	if p.match(NIL) {
		return NewLiteralExpr(nil), nil
	}

	if p.match(NUMBER, STRING) {
		return NewLiteralExpr(p.prevoius().literal), nil
	}

	if p.match(LEFT_PAREN) {
		expr, err := p.expression()

		if err != nil {
			return nil, err
		}

		p.consume(RIGHT_PAREN, "expect ')' after expression.")

		return NewGroupingExpr(expr), nil
	}

	return nil, fmt.Errorf("expected expression")
}

func (p *Parser) consume(t TokenType, msg string) (*Token, error) {
	if p.check(t) {
		token := p.advance()

		return &token, nil
	}

	return nil, fmt.Errorf("error in line %d: %s", p.peek().line, msg)
}

func (p *Parser) match(tokens ...TokenType) bool {
	for _, t := range tokens {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p Parser) check(tt TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return p.peek().tokenType == tt
}

func (p *Parser) advance() Token {
	if !p.isAtEnd() {
		p.current++
	}

	return p.prevoius()
}

func (p Parser) peek() Token {
	return p.tokens[p.current]
}

func (p Parser) isAtEnd() bool {
	return p.peek().tokenType == EOF
}

func (p Parser) prevoius() Token {
	return p.tokens[p.current-1]
}
