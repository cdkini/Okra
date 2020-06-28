package interpret

import "fmt"

type Interpreter struct{}

// TODO: Update with environment variables and other factors
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

func (i *Interpreter) Interpret(expr Expr) {
	val := i.Evaluate(expr)
	fmt.Println(val) // TODO: Add in stringify method to displayed proper output to console
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
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operand", operand)
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
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) - evaluateNumeric(rightOperand)
	case Plus:
		// TODO: Add in concatenation of strings and error handling for non-numeric/non-string addition
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
