// Package main runs Okra interpreter
// on user provided path of a .okr file

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/cdkini/Okra/interpret"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".okr") {
		interpret.ThrowErr(-1, interpret.NewOkraError(0, 0, "Must use \"okra [script]\" to run a .okr file"))
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	interpret.CheckErr(-1, err, interpret.NewOkraError(0, 0, "Path not found"))
	scanner := interpret.NewScanner(string(bytes))
	tokens := scanner.ScanTokens()
	parser := interpret.NewParser(tokens)
	expr := parser.Parse()
	interpreter := interpret.NewInterpreter()
	interpreter.Interpret(expr)

	fmt.Println(tokens)
}
