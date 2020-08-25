package parse

import "Okra/okra/interpreter/ast"

func (p *Parser) statement() ast.Stmt {
	switch {

	case p.match(ast.If):
		return p.ifStmt()

	case p.match(ast.For):
		return p.forStmt()

	case p.match(ast.LeftBrace):
		stmts := p.block()
		return ast.NewBlockStmt(stmts)

	case p.match(ast.Print):
		return p.printStmt()

	default:
		return p.expressionStmt()
	}
}

func (p *Parser) ifStmt() ast.Stmt {
	p.consume(ast.LeftParen, "Expect '(' after 'if'.")
	condition := p.Expression()
	p.consume(ast.RightParen, "Expect ')' after condition.")

	thenBranch := p.statement()
	var elseBranch ast.Stmt
	if p.match(ast.Else) {
		elseBranch = p.statement()
	}

	return ast.NewIfStmt(condition, thenBranch, elseBranch)
}

func (p *Parser) forStmt() ast.Stmt {
	p.consume(ast.LeftParen, "Expect '(' after 'for'.")
	condition := p.Expression()
	p.consume(ast.RightParen, "Expect ')' after condition.")
	body := p.statement()

	return ast.NewForStmt(condition, body)
}

func (p *Parser) block() []ast.Stmt {
	stmts := []ast.Stmt{}

	for !p.check(ast.RightBrace) && !p.isAtEOF() {
		stmts = append(stmts, p.declaration())
	}

	p.consume(ast.RightBrace, "Expect '}' at end of block")

	return stmts
}

func (p *Parser) printStmt() ast.Stmt {
	expr := p.Expression()
	p.consume(ast.Semicolon, "Expect ';' after value")

	return ast.NewPrintStmt(expr)
}

func (p *Parser) expressionStmt() ast.Stmt {
	expr := p.Expression()
	p.consume(ast.Semicolon, "Expect ';' after expression")

	return ast.NewExpressionStmt(expr)
}
