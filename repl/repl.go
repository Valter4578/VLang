package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/valter4578/vlang/lexer"
	"github.com/valter4578/vlang/token"
)

const char = "@"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(char)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
