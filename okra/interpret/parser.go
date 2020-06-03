package interpret

struct Parser {
	tokens Tokens[]
	curr int
}

func (p *Parser) evaluate() Expr {
	return equality()
}

func (p *Parser) equality() Expr {
	expr := comparison()

	for tokenMatch(BangEqual, EqualEqual) {
		operator := prev()
		right := comparison()
		expr = Binary{expr, operator, right}
	}
	return expr
}

func (p *Parser) comparison() Expr {
	expr := addition()

	for tokenMatch(Greater, GreaterEqual, Less, LessEqual) {
		operator := prev()
		right := addition()
		expr = Binary{expr, operator, right}
	}
	return expr 
}

func (p *Parser) addOrSubtract() Expr {

}

func (p *Parser) multiplyOrDivide() Expr {

}

func (p *Parser) unary() Expr {

}

func (p *Parser) primary() Expr {

}

func (p *Parser) match(tokens ...TokenType) bool {
	for i, token := range tokens {
		if p.tokens[p.curr].tokenType == token && p.tokens[p.curr].tokenType != EOF {
			p.advance()
		}
	}
}

func (p *Parser) advance() TokenType {
	if p.tokens[p.curr].tokenType != EOF {
		p.curr++
	}
	return p.tokens[p.curr - 1]
}
