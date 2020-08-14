package interpret

func (p *Parser) statement() Stmt {
	switch {

	case p.match(If):
		return p.ifStmt()
	case p.match(LeftBrace):
		stmts := p.blockStmt()
		return &BlockStmt{stmts}

	case p.match(Print):
		return p.printStmt()

	default:
		return p.expressionStmt()
	}
}

func (p *Parser) ifStmt() Stmt {
	p.consume(LeftParen, "Expect '(' after 'if'.")
	condition := p.Expression()
	p.consume(RightParen, "Expect ')' after condition.")

	thenBranch := p.statement()
	var elseBranch Stmt
	if p.match(Else) {
		elseBranch = p.statement()
	}

	return &IfStmt{condition, thenBranch, elseBranch}
}

func (p *Parser) blockStmt() []Stmt {
	stmts := []Stmt{}

	for !p.check(RightBrace) && !p.isAtEOF() {
		stmts = append(stmts, p.declaration())
	}

	p.consume(RightBrace, "Expect '}' at end of block")

	return stmts
}

func (p *Parser) printStmt() Stmt {
	expr := p.Expression()
	p.consume(Semicolon, "Expect ';' after value")

	return &PrintStmt{expr}
}

func (p *Parser) expressionStmt() Stmt {
	expr := p.Expression()
	p.consume(Semicolon, "Expect ';' after expression")

	return &ExpressionStmt{expr}
}
