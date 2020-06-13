package interpret

import "errors"

type Parser struct {
	tokens []*Token
	curr   int
	errors []OkraError
}

func NewParser(tokens []*Token) *Parser {
	return &Parser{tokens, 0, make([]OkraError, 0)}
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.getPrevToken()
		right := p.comparison()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.addOrSubtract()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.getPrevToken()
		right := p.addOrSubtract()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) addOrSubtract() Expr {
	expr := p.multiplyOrDivide()

	for p.match(Plus, Minus) {
		operator := p.getPrevToken()
		right := p.multiplyOrDivide()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) multiplyOrDivide() Expr {
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.getPrevToken()
		right := p.unary()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) unary() Expr {
	expr := p.primary()

	for p.match(Bang, Minus) {
		operator := p.getPrevToken()
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
		return Literal{p.getPrevToken().literal}
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
		if p.getCurrTokenType() == t {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) advance() Token {
	if p.getCurrTokenType() != EOF {
		p.curr++
	}
	return p.getPrevToken()
}

func (p *Parser) consume(tokenType TokenType, msg string) (Token, error) {
	if p.getCurrTokenType() == tokenType {
		return p.advance(), nil
	}
	return Token{}, errors.New(msg) // TODO: Implement error handling from err.go
}

func (p *Parser) getCurrToken() Token {
	return *p.tokens[p.curr]
}

func (p *Parser) getCurrTokenType() TokenType {
	return p.tokens[p.curr].tokenType
}

func (p *Parser) getPrevToken() Token {
	return *p.tokens[p.curr-1]
}

func (p *Parser) getPrevTokenType() TokenType {
	return p.tokens[p.curr-1].tokenType
}
