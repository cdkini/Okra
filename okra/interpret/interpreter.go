package interpret

type Interpreter struct{}

// TODO: Update with environment variables and other factors
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Evaluate(expr Expr) interface{} {
	switch t := expr.(type) {
	case Unary:
		return i.interpretUnary(t)
	case Binary:
		return i.interpretBinary(t)
	case Grouping:
		return i.interpretGrouping(t)
	case Literal:
		return i.interpretLiteral(t)
	default:
		return nil
	}
}

func (i *Interpreter) interpretLiteral(l Literal) interface{} {
	return l.val
}

func (i *Interpreter) interpretGrouping(g Grouping) interface{} {
	return i.Evaluate(g.expression)
}

func (i *Interpreter) interpretUnary(u Unary) interface{} {
	operand := i.Evaluate(u.operand)

	switch u.operator.tokenType {
	case Minus:
		return -evaluateNumeric(operand)
	case Bang:
		return !isTruthy(operand)
	}
	return nil
}

func (i *Interpreter) interpretBinary(b Binary) interface{} {
	leftOperand := i.Evaluate(b.leftOperand)
	rightOperand := i.Evaluate(b.rightOperand)

	switch b.operator.tokenType {
	case Minus:
		return evaluateNumeric(leftOperand) - evaluateNumeric(rightOperand)
	case Plus:
		// TODO: Add in concatenation of strings
		return evaluateNumeric(leftOperand) + evaluateNumeric(rightOperand)
	case Slash:
		return evaluateNumeric(leftOperand) / evaluateNumeric(rightOperand)
	case Star:
		return evaluateNumeric(leftOperand) * evaluateNumeric(rightOperand)
	case Greater:
		return evaluateNumeric(leftOperand) > evaluateNumeric(rightOperand)
	case Less:
		return evaluateNumeric(leftOperand) < evaluateNumeric(rightOperand)
	case GreaterEqual:
		return evaluateNumeric(leftOperand) >= evaluateNumeric(rightOperand)
	case LessEqual:
		return evaluateNumeric(leftOperand) <= evaluateNumeric(rightOperand)
	case EqualEqual:
		return leftOperand == rightOperand
	case BangEqual:
		return leftOperand != rightOperand
	}
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
