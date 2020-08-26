package interpret

import (
	"Okra/okra/interpreter/ast"
	"Okra/okra/okraerr"
)

// TODO: Add docstring
type Environment struct {
	enclosing *Environment
	vars      map[string]interface{}
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing, make(map[string]interface{})}
}

func (e *Environment) Define(identifier string, value interface{}) {
	e.vars[identifier] = value
}

func (e *Environment) Assign(token ast.Token, value interface{}) {
	if _, ok := e.vars[token.Lexeme]; ok {
		e.vars[token.Lexeme] = value
	} else if e.enclosing != nil {
		e.enclosing.Assign(token, value)
	} else {
		okraerr.ReportErr(token.Col, token.Line, "Variable not declared prior to usage")
	}
}

func (e *Environment) Get(token ast.Token) interface{} {
	if val, ok := e.vars[token.Lexeme]; ok {
		return val
	}
	if e.enclosing != nil {
		return e.enclosing.Get(token)
	}
	okraerr.ReportErr(token.Line, token.Col, "Undefined variable '"+token.Lexeme+"'")
	return nil
}
