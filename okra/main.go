// Package main runs Okra interpreter
// on user provided path of a .okr file

package main

import (
	"Okra/okra/interpreter/interpret"
	"Okra/okra/interpreter/parse"
	"Okra/okra/interpreter/scan"
	"Okra/okra/okraerr"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 || !(os.Args[1] == "run" || os.Args[1] == "fmt") {
		okraerr.ReportErr(0, 0, "Must use one of the following:\n"+
			"  ~ 'okra run [script]' => Runs the Okra interpreter on a .okr file \n"+
			"  ~ 'okra fmt [script]' => Runs the Okra formatter on a .okr file")
	}

	if os.Args[1] == "run" {
		if !strings.HasSuffix(os.Args[2], ".okr") {
			okraerr.ReportErr(0, 0, "File type not supported; please pass a .okr file")
		}
		runFile(os.Args[2])
	}

	if os.Args[1] == "fmt" {
		if !strings.HasSuffix(os.Args[2], ".okr") {
			okraerr.ReportErr(0, 0, "File type not supported; please pass a .okr file")
		}
		fmtFile(os.Args[2])
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	okraerr.CheckErr(err, 0, 0, "Path not found")
	scanner := scan.NewScanner(string(bytes))
	tokens := scanner.ScanTokens()
	parser := parse.NewParser(tokens)
	stmts, _ := parser.Parse()
	interpreter := interpret.NewInterpreter(stmts)
	interpreter.Interpret()
}

func fmtFile(path string) {
	// TODO: Implement formatter!
}
