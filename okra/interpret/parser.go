package interpret

type Parser struct {
	tokens []Token
	curr   int
}

func (p *Parser) evaluate() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.tokens[p.curr]
		right := p.comparison()
		expr = Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.addOrSubtract()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.tokens[p.curr]
		right := p.addOrSubtract()
		expr = Binary{expr, operator, right}
	}

	return expr
}

func (p *Parser) addOrSubtract() Expr {
	expr := p.multiplyOrDivide()

	for p.match(Plus, Minus) {
		operator := p.tokens[p.curr]
		right := p.multiplyOrDivide()
		expr = Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) multiplyOrDivide() Expr {
	return nil
}

func (p *Parser) unary() Expr {
	return nil
}

func (p *Parser) primary() Expr {
	return nil
}

func (p *Parser) match(tokens ...TokenType) bool {
	for _, t := range tokens {
		if p.currTokenType() == t && p.currTokenType() != EOF {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) advance() TokenType {
	if p.currTokenType() != EOF {
		p.curr++
	}
	return p.prevTokenType()
}

func (p *Parser) currTokenType() TokenType {
	return p.tokens[p.curr].tokenType
}

func (p *Parser) prevTokenType() TokenType {
	return p.tokens[p.curr-1].tokenType
}
