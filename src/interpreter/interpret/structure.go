package interpret

import (
	"github.com/cdkini/Okra/src/interpreter/ast"
	"github.com/cdkini/Okra/src/okraerr"
)

// A Structure encapsulates a set of variables and methods to represent a user-created object.
// Structure successfully fulfills all of the Callable interface's methods.
type Structure struct {
	name    string
	methods map[string]*Function
}

func NewStructure(name string, methods map[string]*Function) *Structure {
	return &Structure{name, methods}
}

func (s *Structure) Arity() int {
	if init := s.findMethod("construct"); init != nil {
		return init.Arity()
	}
	return 0
}

func (s *Structure) Call(i *Interpreter, args []interface{}) interface{} {
	instance := NewInstance(*s)
	if init := s.findMethod("construct"); init != nil {
		init.bind(instance).Call(i, args)
	}
	return NewInstance(*s)
}

func (s *Structure) findMethod(method string) *Function {
	if _, ok := s.methods[method]; ok {
		return s.methods[method]
	}
	return nil
}

// Instance represents a specific instance of a previous defined Structure.
// It's usage is limited to the retrieval and storage of instance fields ('this').
type Instance struct {
	class  Structure
	fields map[string]interface{}
}

func NewInstance(s Structure) *Instance {
	return &Instance{s, make(map[string]interface{})}
}

func (i *Instance) get(property ast.Token) interface{} {
	if val, ok := i.fields[property.Lexeme]; ok {
		return val
	}
	if method := i.class.findMethod(property.Lexeme); method != nil {
		return method.bind(i)
	}
	okraerr.ReportErr(0, 0, "Undefined property "+property.Lexeme+".")
	return nil
}

func (i *Instance) set(property ast.Token, val interface{}) {
	i.fields[property.Lexeme] = val
}

// A ReturnValue is a wrapper around a literal expression and is used to represent the resulting value of a function
// or method invokation. In establishing a struct to represent such values, we can use type checking in our
// interpreter's evaluation to determine appropriate control flow.
type ReturnValue struct {
	*ast.LiteralExpr
}

func NewReturnValue(literal *ast.LiteralExpr) *ReturnValue {
	return &ReturnValue{literal}
}
