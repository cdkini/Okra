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
	if l, ok := e.localScope[token.lexeme]; !ok {
		if g, ok := e.globalScope[token.lexeme]; !ok {
			ReportErr(-1, NewOkraError(token.line, token.col, "Undefined variable '"+token.lexeme+"'"))
		} else {
			return g
		}
		return l
	}
	return nil
}

func (e *Environment) assignVar(token Token, value interface{}) {
	if _, ok := e.localScope[token.lexeme]; ok {
		e.localScope[token.lexeme] = value
		return
	}
	if _, ok := e.globalScope[token.lexeme]; ok {
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
