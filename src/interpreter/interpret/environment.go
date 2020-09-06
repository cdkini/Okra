package interpret

import (
	"github.com/cdkini/Okra/src/interpreter/ast"
	"github.com/cdkini/Okra/src/okraerr"
)

// An Environment is a wrapper around Go's map and is used to keep track of lexical scopes and state of Okra objects.
type Environment struct {
	enclosing *Environment           // Pointer to next scope up; nil by default as starting scope is global
	vals      map[string]interface{} // Key - value pairings of object names and associated value
}

func NewEnvironment(enclosing *Environment) *Environment {
	return &Environment{enclosing, make(map[string]interface{}, 0)}
}

// Define is a wrapper around a hashmap 'put' or storage of a key, value pairing.
// This method is to invoked upon the initialization of an object.
// Args: identifier (string)      - The name associated with the object to be stored
//       value      (interface{}) - The current value of the object
// Returns: nil
func (e *Environment) Define(identifier string, value interface{}) {
	e.vals[identifier] = value
}

// Assign alters the value of an already declared and existing object in the Environment.
// This method is to be invoked only upon previously defined objects.
// Args: token (Token)       - The object being assigned; arg is a Token to include column/line context
//       value (interface{}) - The current value of the object
// Returns: nil
// Raises: OkraError if object is declared prior to its usage in a program
func (e *Environment) Assign(token ast.Token, value interface{}) {
	if _, ok := e.vals[token.Lexeme]; ok {
		e.vals[token.Lexeme] = value
	} else if e.enclosing != nil {
		e.enclosing.Assign(token, value)
	} else {
		okraerr.ReportErr(token.Col, token.Line, "Variable not declared prior to usage.")
	}
}

// Define is a wrapper around a hashmap 'get' or retrieval of a key, value pairing.
// The current scope is first analyzed and if the variable is not contained therein, the method is recursively called
// using the Environment enclosing to check subsequent outer scopes.
// Args: token (Token) - The object being retrieved; arg is a Token to include column/line context
// Returns: nil
// Raises: OkraError if object is not in current Environment or any subsequent enclosings
func (e *Environment) Get(token ast.Token) interface{} {
	if val, ok := e.vals[token.Lexeme]; ok {
		return val
	}
	if e.enclosing != nil {
		return e.enclosing.Get(token)
	}
	okraerr.ReportErr(token.Line, token.Col, "Undefined variable '"+token.Lexeme+"'")
	return nil
}
