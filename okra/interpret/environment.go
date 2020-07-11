package interpret

// TODO: Add docstring
type Environment struct {
	localScope  map[string]interface{}
	globalScope map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{make(map[string]interface{}), make(map[string]interface{})}
}

func (e *Environment) putVar(identifier string, value interface{}, isLocal bool) {
	if isLocal {
		e.localScope[identifier] = value
	} else {
		e.globalScope[identifier] = value
	}
}

func (e *Environment) getVar(token Token) interface{} {
	val1, ok1 := e.localScope[token.lexeme]
	if !ok1 {
		val2, ok2 := e.globalScope[token.lexeme]
		if !ok2 {
			ReportErr(-1, NewOkraError(token.col, token.line, "Variable not declared prior to usage"))
		}
		return val2
	}
	return val1
}

func (e *Environment) assignVar(token Token, value interface{}) {
	if _, ok := e.localScope[token.lexeme]; ok {
		e.localScope[token.lexeme] = value
		return
	}
	if _, ok2 := e.globalScope[token.lexeme]; ok2 {
		e.globalScope[token.lexeme] = value
		return
	}
	ReportErr(-1, NewOkraError(token.col, token.line, "Variable not declared prior to usage"))
}

func (e *Environment) clearLocalScope() {
	for k := range e.localScope {
		delete(e.localScope, k)
	}
}
