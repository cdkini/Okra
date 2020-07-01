package interpret

type Environment struct {
	varMap map[string]interface{}
}

func (e *Environment) putVar(k string, v interface{}) {
	e.varMap[k] = v
}

func (e *Environment) getVar(k string) interface{} {
	val, ok := e.varMap[k]
	if !ok {
		ReportErr(-1, NewOkraError(0, 0, "Placeholder"))
	}
	return val
}
