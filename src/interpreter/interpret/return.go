package interpret

import "Okra/src/interpreter/ast"

type ReturnValue struct {
	*ast.LiteralExpr
}

func NewReturnValue(literal *ast.LiteralExpr) *ReturnValue {
	return &ReturnValue{literal}
}
