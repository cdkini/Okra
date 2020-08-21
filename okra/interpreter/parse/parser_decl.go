package parse

import "Okra/okra/interpreter/ast"

func (p *Parser) declaration() ast.Stmt {
	switch {
	case p.match(ast.Variable):
		return p.varDeclaration()

	default:
		return p.statement()
	}
}

func (p *Parser) varDeclaration() ast.Stmt {
	identifier := p.consume(ast.Identifier, "Expect variable name")

	var initializer ast.Expr
	if p.match(ast.Equal) {
		initializer = p.Expression()
	}

	p.consume(ast.Semicolon, "Expect ';' after variable declaration")
	return ast.NewVariableStmt(identifier, initializer)
}
