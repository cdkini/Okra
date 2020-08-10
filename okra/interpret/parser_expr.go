package interpret

func (p *Parser) Expression() Expr {
	return p.assignment()
}

func (p *Parser) assignment() Expr {
	expr := p.equality()

	if p.match(Equal) {
		value := p.assignment()

		exprVar, ok := value.(*VariableExpr)
		if ok {
			return &AssignmentExpr{exprVar.identifier, value}
		}

		curr := p.currentToken()
		ReportErr(NewOkraError(curr.line, curr.col, "Invalid assignment target"))
	}

	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.previousToken()
		rightOperand := p.comparison()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.additionOrSubtraction()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.previousToken()
		rightOperand := p.additionOrSubtraction()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) additionOrSubtraction() Expr {
	expr := p.multiplicationOrDivision()

	for p.match(Minus, Plus) {
		operator := p.previousToken()
		rightOperand := p.multiplicationOrDivision()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) multiplicationOrDivision() Expr {
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.previousToken()
		rightOperand := p.unary()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(Bang, Minus) {
		operator := p.previousToken()
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
		return &LiteralExpr{p.previousToken().literal}

	case p.match(String):
		return &LiteralExpr{p.previousToken().lexeme}

	case p.match(Identifier):
		return &VariableExpr{p.previousToken()}

	case p.match(LeftParen):
		expr := p.Expression()

		p.consume(RightParen, "Expect ')' after expression.")
		return &GroupingExpr{expr}

	default:
		curr := p.currentToken()
		ReportErr(NewOkraError(curr.line, curr.col, "Expect expression"))
		return nil
	}
}
