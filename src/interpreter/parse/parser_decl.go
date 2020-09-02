package parse

import "Okra/src/interpreter/ast"

// declaration is a helper method that is at the very start of the recursive descent process. If the current item being
// parsed fits the requirements of a structure, function, or variable declaration, subsequent helper methods are
// triggered to return an appropriate Stmt instance. Otherwise, the parser jumps to non-declarative Stmts.
// Args: nil
// Returns: Instance of Stmt that fits the ENBF / context-free grammar rules as set by Okra
func (p *Parser) declaration() ast.Stmt {

	switch {

	case p.match(ast.Struct):
		return p.structure()

	case p.match(ast.Func):
		return p.function()

	case p.match(ast.Variable):
		return p.variable()

	default:
		return p.statement()
	}
}

func (p *Parser) structure() ast.Stmt {
	name := p.consume(ast.Identifier, "Expect struct name.")
	p.consume(ast.LeftBrace, "Expect '{' before struct body.")

	var methods []ast.FuncStmt
	for !p.check(ast.RightBrace) && !p.isAtEOF() {
		if f, ok := p.function().(*ast.FuncStmt); ok {
			methods = append(methods, *f)
		}
	}
	p.consume(ast.RightBrace, "Expect '}' after struct body.")
	return ast.NewStructStmt(name, methods)
}

func (p *Parser) function() ast.Stmt {
	identifier := p.consume(ast.Identifier, "Expect valid identifier.")
	p.consume(ast.Colon, "Expect ':' after identifier.")
	var params []ast.Token
	if !p.check(ast.Colon) {
		for {
			params = append(params, p.consume(ast.Identifier, "Expect ':' after parameters."))
			if !p.match(ast.Comma) {
				break
			}
		}
	}
	p.consume(ast.Colon, "Expect ':' after parameters.")
	p.consume(ast.LeftBrace, "Expect '{' before func body.")
	body := p.blockStmt()
	return ast.NewFuncStmt(identifier, params, body)
}

func (p *Parser) variable() ast.Stmt {
	identifier := p.consume(ast.Identifier, "Expect variable name.")

	var initializer ast.Expr
	if p.match(ast.Colon) {
		initializer = p.Expression()
	}

	p.consume(ast.Semicolon, "Expect ';' after variable declaration.")
	return ast.NewVariableStmt(identifier, initializer)
}
