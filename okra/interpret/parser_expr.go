package interpret

import "fmt"

func (p *Parser) Expression() Expr {
	return p.assignment()
}

func (p *Parser) assignment() Expr {
	expr := p.equality()

	if p.match(Equal) {
		equals := p.previous()
		value := p.assignment()

		exprVar, ok := value.(*VariableExpr)
		if ok {
			return &AssignmentExpr{exprVar.identifier, value}
		}

		// FIXME: Add OkraError instance
		// panic(&ParseError{equals, "Invalid assignment target"})
		fmt.Print(equals) // TODO: Remove upon fixing method
	}

	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.previous()
		rightOperand := p.comparison()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.additionOrSubtraction()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.previous()
		rightOperand := p.additionOrSubtraction()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) additionOrSubtraction() Expr {
	expr := p.multiplicationOrDivision()

	for p.match(Minus, Plus) {
		operator := p.previous()
		rightOperand := p.multiplicationOrDivision()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) multiplicationOrDivision() Expr {
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.previous()
		rightOperand := p.unary()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(Bang, Minus) {
		operator := p.previous()
		operand := p.unary()

		return &UnaryExpr{operator, operand}
	}

	return p.primary()
}

func (p *Parser) primary() Expr {
	switch {
	case p.match(False):
		return &LiteralExpr{false}

	case p.match(True):
		return &LiteralExpr{true}

	case p.match(Null):
		return &LiteralExpr{nil}

	case p.match(Numeric):
		return &LiteralExpr{p.previous().literal}

	case p.match(String):
		return &LiteralExpr{p.previous().lexeme}

	case p.match(Identifier):
		return &VariableExpr{p.previous()}

	case p.match(LeftParen):
		expr := p.Expression()

		p.consume(RightParen, "Expect ')' after expression.")
		return &GroupingExpr{expr}

	default:
		// FIXME: Add OkraError instancereturn nil
		// panic(&ParseError{p.peek(), "Expect expression"})
		return nil // TODO: Remove upon fixing method
	}
}
