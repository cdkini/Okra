package parse

import (
	"github.com/cdkini/Okra/src/interpreter/ast"
	"github.com/cdkini/Okra/src/okraerr"
)

// Expression is triggered to parse an Expr or the contents of a Stmt that includes an Expr as a property.
// The method uses recursive descent like the rest of the parser, moving down parser helper methods in order of
// precedence until a set of grammar rules are met.
// Args: nil
// Returns: Instance of Expr that fits the ENBF / context-free grammar rules as set by Okra
func (p *Parser) Expression() ast.Expr {
	return p.assignment()
}

func (p *Parser) assignment() ast.Expr {
	expr := p.or()

	if p.match(ast.Colon) {
		prev := p.prevToken()
		value := p.assignment()

		switch expr := expr.(type) {
		case *ast.VariableExpr:
			return ast.NewAssignmentExpr(expr.Identifier, value)
		case *ast.GetExpr:
			return ast.NewSetExpr(expr.Object, expr.Property, value)
		default:
			curr := p.currToken()
			okraerr.ReportErr(curr.Line, curr.Col, "Invalid assignment target: '"+prev.Lexeme+"'.")
		}
	}
	return expr
}

func (p *Parser) or() ast.Expr {
	expr := p.and()
	for p.match(ast.Or) {
		operator := p.prevToken()
		rightOperand := p.and()
		expr = ast.NewLogicalExpr(expr, operator, rightOperand)
	}
	return expr
}

func (p *Parser) and() ast.Expr {
	expr := p.equality()

	for p.match(ast.And) {
		operator := p.prevToken()
		rightOperand := p.equality()
		expr = ast.NewLogicalExpr(expr, operator, rightOperand)
	}
	return expr
}

func (p *Parser) equality() ast.Expr {
	expr := p.comparison()

	for p.match(ast.BangEqual, ast.Equal) {
		operator := p.prevToken()
		rightOperand := p.comparison()

		expr = ast.NewBinaryExpr(expr, operator, rightOperand)
	}
	return expr
}

func (p *Parser) comparison() ast.Expr {
	expr := p.additionOrSubtraction()

	for p.match(ast.Greater, ast.GreaterEqual, ast.Less, ast.LessEqual) {
		operator := p.prevToken()
		rightOperand := p.additionOrSubtraction()

		expr = ast.NewBinaryExpr(expr, operator, rightOperand)
	}
	return expr
}

func (p *Parser) additionOrSubtraction() ast.Expr {
	expr := p.multiplicationOrDivision()

	for p.match(ast.Minus, ast.Plus) {
		operator := p.prevToken()
		rightOperand := p.multiplicationOrDivision()

		expr = ast.NewBinaryExpr(expr, operator, rightOperand)
	}
	return expr
}

func (p *Parser) multiplicationOrDivision() ast.Expr {
	expr := p.unary()

	for p.match(ast.Slash, ast.Star) {
		operator := p.prevToken()
		rightOperand := p.unary()

		expr = ast.NewBinaryExpr(expr, operator, rightOperand)
	}
	return expr
}

func (p *Parser) unary() ast.Expr {
	if p.match(ast.Bang, ast.Minus) {
		operator := p.prevToken()
		operand := p.unary()

		return ast.NewUnaryExpr(operator, operand)
	}
	return p.call()
}

func (p *Parser) call() ast.Expr {
	expr := p.primary()

	for {
		if p.match(ast.LeftParen) {
			expr = p.finishCall(expr)
		} else if p.match(ast.Dot) {
			property := p.consume(ast.Identifier, "Expect property name after '.'.")
			expr = ast.NewGetExpr(expr, property)
		} else {
			break
		}
	}
	return expr
}

func (p *Parser) finishCall(callee ast.Expr) ast.Expr {
	var args []ast.Expr
	if !p.check(ast.RightParen) {
		for {
			args = append(args, p.Expression())
			if !p.match(ast.Comma) {
				break
			}
		}
	}
	paren := p.consume(ast.RightParen, "Expect ')' after function arguments.")

	return ast.NewCallExpr(callee, paren, args)
}

func (p *Parser) primary() ast.Expr {
	switch {

	case p.match(ast.False):
		return ast.NewLiteralExpr(false)

	case p.match(ast.True):
		return ast.NewLiteralExpr(false)

	case p.match(ast.Null):
		return ast.NewLiteralExpr(nil)

	case p.match(ast.Numeric):
		return ast.NewLiteralExpr(p.prevToken().Literal)

	case p.match(ast.String):
		return ast.NewLiteralExpr(p.prevToken().Lexeme)

	case p.match(ast.Identifier):
		return ast.NewVariableExpr(p.prevToken())

	case p.match(ast.This):
		return ast.NewThisExpr(p.prevToken())

	case p.match(ast.LeftParen):
		expr := p.Expression()

		p.consume(ast.RightParen, "Expect ')' after expression.")
		return ast.NewGroupingExpr(expr)

	default:
		curr := p.currToken()
		okraerr.ReportErr(curr.Line, curr.Col, "Expect expression.")
		return nil
	}
}
