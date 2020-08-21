package stdlib

import "Okra/okra/interpreter/interpret"

func BuildStdlib() (stdlib map[string]interpret.Callable) {
	stdlib["hello"] = &StdStr{}
	return stdlib
}
