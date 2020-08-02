package interpret

func (p *Parser) declaration() Stmt {
	switch {
	case p.match(Variable):
		return p.varDeclaration()

	default:
		return p.statement()
	}
}

func (p *Parser) varDeclaration() Stmt {
	identifier := p.consume(Identifier, "Expect variable name")

	var initializer Expr
	if p.match(Equal) {
		initializer = p.expression()
	}

	p.consume(Semicolon, "Expect ';' after variable declaration")
	return &VariableStmt{identifier, initializer}
}
