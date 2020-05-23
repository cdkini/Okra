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
	checkErr(-1, err) // Error if file path invalid
	scanner := NewScanner(string(bytes))
	tokens := scanner.scanTokens()
}

func reportError(line int, loc string, message string) {
	fmt.Println("[Line", line, "] Error", loc, ":", message)
}

func checkErr(code int, err error) {
	// TODO: Update this to include a specific error type / message
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(code)
	}
}
