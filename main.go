package main

import (
	"flag"
	"github.com/valter4578/vlang/repl"
	"os"
)

func main() {
	fileFlag := flag.String("file", "", "Enter file name to evaluate file")
	flag.Parse()

	if *fileFlag == "" {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		in := parseFile(*fileFlag)
		evalFile(in)
	}

}
