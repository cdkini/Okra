package main

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/cdkini/Okra/src/interpreter/interpret"
	"github.com/cdkini/Okra/src/interpreter/parse"
	"github.com/cdkini/Okra/src/interpreter/scan"
	"github.com/cdkini/Okra/src/okraerr"
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
	// Check for valid path
	bytes, err := ioutil.ReadFile(path)
	okraerr.CheckErr(err, 0, 0, "Path not found")

	// Lex or tokenize input stream
	scanner := scan.NewScanner(string(bytes))
	tokens := scanner.ScanTokens()

	// Parse tokens into AST
	parser := parse.NewParser(tokens)
	stmts := parser.Parse()

	// Traverse AST to generate output
	interpreter := interpret.NewInterpreter()
	// interpreter.LoadStdlib(stdlib.BuildStdlib()) // TODO: Open to add stdlib import
	interpreter.Interpret(stmts)
}

func fmtFile(path string) {
	// TODO: Implement formatter!
}
