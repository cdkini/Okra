package interpret

// TODO: Add docstring
type Environment struct {
	varMap map[string]interface{}
}

func NewEnvironment() *Environment {
	return &Environment{make(map[string]interface{})}
}

func (e *Environment) putVar(k string, v interface{}) {
	e.varMap[k] = v
}

func (e *Environment) getVar(k Token) interface{} {
	val, ok := e.varMap[k.lexeme]
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
	return val
}
