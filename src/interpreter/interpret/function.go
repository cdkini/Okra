package interpret

import (
	"github.com/cdkini/Okra/src/interpreter/ast"
)

// A Function is a wrapper around a FuncStmt and is instantiated upon the interpreter
// visiting a parsed function or structure declaration.
// Function successfully fulfills all of the Callable interface's methods.
type Function struct {
	declaration   ast.FuncStmt
	closure       *Environment // Key-value mapping to enable appropriate lexical scoping for closures
	isConstructor bool         // Only fulfilled by 'construct' method
}

func NewFunction(declaration ast.FuncStmt, closure *Environment, isConstructor bool) *Function {
	return &Function{declaration, closure, isConstructor}
}

func (f *Function) Arity() int {
	return len(f.declaration.Params)
}

func (f *Function) Call(i *Interpreter, args []interface{}) interface{} {
	env := NewEnvironment(f.closure)
	for i, token := range f.declaration.Params {
		env.Define(token.Lexeme, args[i])
	}

	block := i.executeBlock(f.declaration.Body, env)
	if r, ok := block.(*ReturnValue); ok {
		return r.LiteralExpr.Val
	}
	if f.isConstructor {
		return f.closure.vals["this"]
	}
	return nil
}

func (f *Function) bind(instance *Instance) *Function {
	env := NewEnvironment(f.closure)
	env.Define("this", instance)
	return NewFunction(f.declaration, env, f.isConstructor)
}
