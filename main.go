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

func runFile(path string) (bytes []byte, err error) {
	bytes, err = ioutil.ReadFile(path)
	check(err)
	run(string(bytes))
	return bytes, err
}

func run(source string) {

}

func check(e error) {
	if e != nil {
		fmt.Println("Error:", e)
	}
}
