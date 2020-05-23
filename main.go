package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: Must use \"okra [script]\" to run a .okr file")
		os.Exit(-1)
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	errorCheck(err, -1)
	scanner := NewScanner(string(bytes))
	tokens, err := scanner.scanTokens()
	errorCheck(err, -1)
}

func reportError(line int, loc string, message string) {
	fmt.Println("[Line", line, "] Error", loc, ":", message)
}

func errorCheck(e error, code int) {
	if e != nil {
		fmt.Println("Error:", e)
		os.Exit(code)
	}
}
