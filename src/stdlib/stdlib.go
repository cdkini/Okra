package stdlib

import "Okra/src/interpreter/interpret"

func BuildStdlib() (stdlib map[string]interpret.Callable) {
	stdlib["hello"] = &StdStr{}
	return stdlib
}
