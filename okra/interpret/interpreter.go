package interpret

import (
	"fmt"
	"strings"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct{}

// TODO: Update with environment variables and other factors
func NewInterpreter() *Interpreter {
	return &Interpreter{}
}

// Interpret evaluates an expression and returns the result to the user
// Args: expr [Expr]: The expression we wish to break down
// Returns: String representation of evaluated expression
func (i *Interpreter) Interpret(expr Expr) string {
	val := fmt.Sprintf("%v", expr.accept(i))
	fmt.Println(val) // TODO: Add in stringify method to displayed proper output to console
	return val
}

func (i *Interpreter) visitLiteral(l Literal) interface{} {
	return l.val
}

func (i *Interpreter) visitGrouping(g Grouping) interface{} {
	return g.expression.accept(i)
}

func (i *Interpreter) visitUnary(u Unary) interface{} {
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

func (i *Interpreter) visitBinary(b Binary) interface{} {
	leftOperand := b.leftOperand.accept(i)
	rightOperand := b.rightOperand.accept(i)

	switch b.operator.tokenType {
	case Minus:
		checkNumericValidity("Runtime Error => \"-\" used on non-numeric operands", leftOperand, rightOperand)
		return evaluateNumeric(leftOperand) - evaluateNumeric(rightOperand)
	case Plus:
		if isString(leftOperand) && isString(rightOperand) {
			return concatenateString(leftOperand, rightOperand) // FIXME: Currently does not work!
		}
		if isNumeric(leftOperand) && isNumeric(rightOperand) {
			checkNumericValidity("Runtime Error => \"+\" used on non-numeric operands", leftOperand, rightOperand)
			return evaluateNumeric(leftOperand) + evaluateNumeric(rightOperand)
		}
		ReportErr(-1, NewOkraError(0, 0, "Runtime Error => \"+\" used on incompatible operands"))
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

func isNumeric(num interface{}) bool {
	_, ok := num.(float64)
	if !ok {
		return false
	}
	return true
}

func isString(str interface{}) bool {
	_, ok := str.(string)
	if !ok {
		return false
	}
	return true
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

func concatenateString(strs ...interface{}) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(fmt.Sprintf("%v", str))
	}
	return sb.String()
}
