package interpret

struct Parser {
	tokens Tokens[]
	curr int
}

func (p *Parser) evaluate() Expr {
	return equality()
}

func (p *Parser) equality() Expr {

}

func (p *Parser) comparsion() Expr {

}

func (p *Parser) addOrSubtract() Expr {

}

func (p *Parser) multiplyOrDivide() Expr {

}

func (p *Parser) unary() Expr {

}

func (p *Parser) primary() Expr {
	
}