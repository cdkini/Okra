package parse

import "Okra/okra/interpreter/ast"

func (p *Parser) declaration() ast.Stmt {

	switch {

	case p.match(ast.Func):
		return p.function()

	case p.match(ast.Variable):
		return p.varDeclaration()

	default:
		return p.statement()
	}
}

func (p *Parser) function() ast.Stmt {
	identifier := p.consume(ast.Identifier, "Expect valid identifier.")
	p.consume(ast.LeftParen, "Expect '(' after identifier.")
	var params []ast.Token
	if !p.check(ast.RightParen) {
		for p.match(ast.Comma) {
			params = append(params, p.consume(ast.Identifier, "Expect parameter name."))
		}
	}
	p.consume(ast.RightParen, "Expect ')' after parameters.")
	body := p.block()
	return ast.NewFuncStmt(identifier, params, body)
}

func (p *Parser) varDeclaration() ast.Stmt {
	identifier := p.consume(ast.Identifier, "Expect variable name.")

	var initializer ast.Expr
	if p.match(ast.Equal) {
		initializer = p.Expression()
	}

	p.consume(ast.Semicolon, "Expect ';' after variable declaration")
	return ast.NewVariableStmt(identifier, initializer)
}
