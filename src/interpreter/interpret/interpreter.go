package interpret

import (
	"Okra/src/interpreter/ast"
	"Okra/src/okraerr"
)

// An Interpreter takes in a given expression and evaluates it into its most basic literal form.
// Interpreter inherits from the Visitor interface, allowing it interact with all Expr types.
type Interpreter struct {
	env     *Environment
	globals *Environment // stdlib imports or other preprocessed objects
}

func NewInterpreter() *Interpreter {
	// TODO: Open to add standard library methods as part of global
	return &Interpreter{NewEnvironment(nil), NewEnvironment(nil)}
}

// TODO:
// func (i *Interpreter) LoadStdlib(stdlib map[string]Callable) {
// 	for k, v := range stdlib {
// 		i.globalEnv.Define(k, v)
// 	}
// }

// TODO: Update docstring after changes from stmt
func (i *Interpreter) Interpret(stmts []ast.Stmt) {
	for _, stmt := range stmts {
		i.interpretStmt(stmt)
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

func evalNumeric(i interface{}) float64 {
	t, ok := i.(float64)
	if !ok {
		okraerr.ReportErr(0, 0, "Expect numeric")
	}
	return t
}

func evalString(i interface{}) string {
	t, ok := i.(string)
	if !ok {
		okraerr.ReportErr(0, 0, "Expect string")
	}
	return t
}

func checkNumericValidity(msg string, i ...interface{}) {
	for _, n := range i {
		switch n.(type) {
		case float64:
			continue
		default:
			okraerr.ReportErr(0, 0, msg)
		}
	}
}
