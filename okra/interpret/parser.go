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
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.tokens[p.curr]
		right := p.unary()
		expr = Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) unary() Expr {
	expr := p.primary()

	for p.match(Bang, Minus) {
		operator := p.tokens[p.curr]
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
	} else if p.match(Numeric, String) {
		return Literal{p.tokens[p.curr-1].literal}
	} else if p.match(LeftParen) {
		expr := p.evaluate()
		p.consume(RightParen, "Expect ')' after expression.")
		return Grouping{expr}
	} else if p.match(LeftBracket) {
		expr := p.evaluate()
		p.consume(RightBracket, "Expect ']' after expression.")
		return Grouping{expr}
	} else if p.match(LeftBrace) {
		expr := p.evaluate()
		p.consume(RightBrace, "Expect '}' after expression.")
		return Grouping{expr}
	}
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

func (p *Parser) consume(tokenType TokenType, message string) (Token, error) {
	if p.currTokenType() == EOF {
	}
}
