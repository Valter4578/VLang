package main

import (
	"os"

	"github.com/valter4578/vlang/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
