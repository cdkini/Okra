package interpret

// FIXME: Add block statement
func (p *Parser) statement() Stmt {
	switch {
	case p.match(LeftBrace):
		stmts := p.blockStmt()
		return &BlockStmt{stmts}

	case p.match(Print):
		return p.printStmt()

	default:
		return p.expressionStmt()
	}
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
