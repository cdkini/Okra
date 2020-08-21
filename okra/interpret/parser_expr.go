package interpret

func (p *Parser) Expression() Expr {
	return p.assignment()
}

func (p *Parser) assignment() Expr {
	expr := p.or()

	if p.match(Equal) {
		prev := p.prevToken()
		value := p.assignment()

		switch t := expr.(type) {
		case *VariableExpr:
			return &AssignmentExpr{t.identifier, value}
		default:
			curr := p.currToken()
			ReportErr(curr.line, curr.col, "Invalid assignment target: '"+prev.lexeme+"'")
		}

	}

	return expr
}

func (p *Parser) or() Expr {
	expr := p.and()
	for p.match(Or) {
		operator := p.prevToken()
		rightOperand := p.and()
		expr = &LogicalExpr{expr, operator, rightOperand}
	}
	return expr
}

func (p *Parser) and() Expr {
	expr := p.equality()

	for p.match(And) {
		operator := p.prevToken()
		rightOperand := p.equality()
		expr = &LogicalExpr{expr, operator, rightOperand}
	}
	return expr
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for p.match(BangEqual, EqualEqual) {
		operator := p.prevToken()
		rightOperand := p.comparison()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.additionOrSubtraction()

	for p.match(Greater, GreaterEqual, Less, LessEqual) {
		operator := p.prevToken()
		rightOperand := p.additionOrSubtraction()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) additionOrSubtraction() Expr {
	expr := p.multiplicationOrDivision()

	for p.match(Minus, Plus) {
		operator := p.prevToken()
		rightOperand := p.multiplicationOrDivision()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) multiplicationOrDivision() Expr {
	expr := p.unary()

	for p.match(Slash, Star) {
		operator := p.prevToken()
		rightOperand := p.unary()

		expr = &BinaryExpr{expr, operator, rightOperand}
	}

	return expr
}

func (p *Parser) unary() Expr {
	if p.match(Bang, Minus) {
		operator := p.prevToken()
		operand := p.unary()

		return &UnaryExpr{operator, operand}
	}

	return p.call()
}

func (p *Parser) call() Expr {
	expr := p.primary()

	for {
		if p.match(LeftParen) {
			expr = p.finishCall(expr)
		} else {
			break
		}
	}

	return expr
}

func (p *Parser) finishCall(callee Expr) Expr {
	var args []Expr
	if !p.check(RightParen) {
		for p.match(Comma) {
			args = append(args, p.Expression())
		}
	}
	paren := p.consume(RightParen, "Expect ')' after function arguments.")

	return &CallExpr{callee, paren, args}
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
		return &LiteralExpr{p.prevToken().literal}

	case p.match(String):
		return &LiteralExpr{p.prevToken().lexeme}

	case p.match(Identifier):
		return &VariableExpr{p.prevToken()}

	case p.match(LeftParen):
		expr := p.Expression()

		p.consume(RightParen, "Expect ')' after expression.")
		return &GroupingExpr{expr}

	default:
		curr := p.currToken()
		ReportErr(curr.line, curr.col, "Expect expression")
		return nil
	}
}
