// Package main runs Okra interpreter
// on user provided path of a .okr file

package main

import (
	"io/ioutil"
	"os"
	"strings"

	"Okra/okra/interpret"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".okr") {
		interpret.ReportErr(-1, interpret.NewOkraError(0, 0, "Must use \"okra [script]\" to run a .okr file"))
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	interpret.CheckErr(err, interpret.NewOkraError(0, 0, "Path not found"))
	scanner := interpret.NewScanner(string(bytes))
	tokens := scanner.ScanTokens()
	parser := interpret.NewParser(tokens)
	stmts, parseErrors := parser.Parse()
	if parseErrors {
		// FIXME: Add support for reporting parse errors
	}
	interpreter := interpret.NewInterpreter(stmts)
	interpreter.Interpret()
}
