package interpret

import "Okra/okra/interpreter/ast"

type ReturnValue struct {
	*ast.LiteralExpr
}

func NewReturnValue(literal *ast.LiteralExpr) *ReturnValue {
	return &ReturnValue{literal}
}
