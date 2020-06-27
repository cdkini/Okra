package interpret

import "strings"

type ASTPrinter struct {
	Visitor
}

func (a *ASTPrinter) print(expr Expr) string {
	err := expr.accept(a)
	if err != nil {
		return ""
	}
	var sb strings.Builder
	switch t := expr.(type) {
	case Unary:

	case Binary:

	case Grouping:

	case Literal:
		sb += expr.val
	}

}
