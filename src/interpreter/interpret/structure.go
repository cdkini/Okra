package interpret

import (
	"Okra/src/interpreter/ast"
	"Okra/src/okraerr"
)

// TODO:
type Struct struct {
	name    string
	methods map[string]*Function
}

func NewStruct(name string, methods map[string]*Function) *Struct {
	return &Struct{name, methods}
}

func (s *Struct) Arity() int {
	if init := s.findMethod("construct"); init != nil {
		return init.Arity()
	}
	return 0
}

func (s *Struct) Call(i *Interpreter, args []interface{}) interface{} {
	instance := NewInstance(*s)
	if init := s.findMethod("construct"); init != nil {
		init.bind(instance).Call(i, args)
	}
	return NewInstance(*s)
}

func (s *Struct) findMethod(method string) *Function {
	if _, ok := s.methods[method]; ok {
		return s.methods[method]
	}
	return nil
}

// TODO:
type Instance struct {
	class  Struct
	fields map[string]interface{}
}

func NewInstance(class Struct) *Instance {
	return &Instance{class, make(map[string]interface{})}
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
