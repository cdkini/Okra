package interpret

import "errors"

type Parser struct {
	tokens []*Token
	curr   int
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.prevToken()
		right := p.comparison()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.addOrSubtract()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.prevToken()
		right := p.addOrSubtract()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) addOrSubtract() Expr {
	expr := p.multiplyOrDivide()

	for p.match(Plus, Minus) {
		operator := p.prevToken()
		right := p.multiplyOrDivide()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) multiplyOrDivide() Expr {
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.prevToken()
		right := p.unary()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) unary() Expr {
	expr := p.primary()

	for p.match(Bang, Minus) {
		operator := p.prevToken()
		right := p.primary()
		expr = Unary{operator, right}
	}

	return expr
}

func (p *Parser) primary() Expr {
	if p.match(True) {
		return Literal{true}
	} else if p.match(False) {
		return Literal{false}
	} else if p.match(Null) {
		return Literal{nil}
	}

	if p.match(Numeric, String) {
		return Literal{p.prevToken().literal}
	}

	if p.match(LeftParen) {
		expr := p.expression()
		p.consume(RightParen, "Expect ')' after expression.")
		return Grouping{expr}
	}

	return nil
}

func (p *Parser) match(tokens ...TokenType) bool {
	for _, t := range tokens {
		if p.currTokenType() == t {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) advance() Token {
	if p.currTokenType() != EOF {
		p.curr++
	}
	return p.prevToken()
}

func (p *Parser) consume(tokenType TokenType, msg string) (Token, error) {
	if p.currTokenType() == tokenType {
		return p.advance(), nil
	}
	return Token{}, errors.New(msg) // TODO: Implement error handling from err.go
}

func (p *Parser) currToken() Token {
	return *p.tokens[p.curr]
}

func (p *Parser) currTokenType() TokenType {
	return p.tokens[p.curr].tokenType
}

func (p *Parser) prevToken() Token {
	return *p.tokens[p.curr-1]
}

func (p *Parser) prevTokenType() TokenType {
	return p.tokens[p.curr-1].tokenType
}
