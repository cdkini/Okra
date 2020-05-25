package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".okr") {
		fmt.Println("Error: Must use \"okra [script]\" to run a .okr file")
		os.Exit(-1)
	}
	runFile(os.Args[1])
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	checkErr(-1, err)
	scanner := NewScanner(string(bytes))
	tokens := scanner.scanTokens()
}

func reportError(line int, loc string, message string) {
	fmt.Println("[Line", line, "] Error", loc, ":", message)
}

func checkErr(code int, err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(code)
	}
}
