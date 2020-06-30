package interpret

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct{}

// TODO: Update with environment variables and other factors
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret(stmts []Stmt) {
	for _, stmt := range stmts {
		stmt.accept(i)
	}
}

func (i *Interpreter) visitExpressionStmt(stmt ExpressionStmt) {
	stmt.expr.accept(i)
}

func (i *Interpreter) visitLiteralExpr(l LiteralExpr) interface{} {
	return l.val
}

func (i *Interpreter) visitGroupingExpr(g GroupingExpr) interface{} {
	return g.expression.accept(i)
}

func (i *Interpreter) visitUnaryExpr(u UnaryExpr) interface{} {
	operand := u.operand.accept(i)

	switch u.operator.tokenType {
	case Minus:
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operand", operand)
		return -evaluateNumeric(operand)
	case Bang:
		return !isTruthy(operand)
	}
	return nil
}

func (i *Interpreter) visitBinaryExpr(b BinaryExpr) interface{} {
	leftOperand := b.leftOperand.accept(i)
	rightOperand := b.rightOperand.accept(i)

	switch b.operator.tokenType {
	case Minus:
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) - evaluateNumeric(rightOperand)
	case Plus:
		checkNumericValidity("Runtime Error => \"+\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) + evaluateNumeric(rightOperand)
	case Slash:
		checkNumericValidity("Runtime Error => \"/\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) / evaluateNumeric(rightOperand)
	case Star:
		checkNumericValidity("Runtime Error => \"*\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) * evaluateNumeric(rightOperand)
	case Greater:
		checkNumericValidity("Runtime Error => \">\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) > evaluateNumeric(rightOperand)
	case Less:
		checkNumericValidity("Runtime Error => \"<\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) < evaluateNumeric(rightOperand)
	case GreaterEqual:
		checkNumericValidity("Runtime Error => \">=\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) >= evaluateNumeric(rightOperand)
	case LessEqual:
		checkNumericValidity("Runtime Error => \"<=\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) <= evaluateNumeric(rightOperand)
	case EqualEqual:
		return leftOperand == rightOperand
	case BangEqual:
		return leftOperand != rightOperand
	}
	return nil
}

func isTruthy(i interface{}) bool {
	switch val := i.(type) {
	case nil:
		return false
	case bool:
		return val
	default:
		return true
	}
}

func evaluateNumeric(i interface{}) float64 {
	t, ok := i.(float64)
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
	return t
}

func evaluateString(i interface{}) string {
	t, ok := i.(string)
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
	return t
}

func checkNumericValidity(msg string, i ...interface{}) {
	for _, n := range i {
		switch n.(type) {
		case float64:
			continue
		default:
			ReportErr(-1, NewOkraError(0, 0, msg))
		}
	}
}
