package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey-go/lexer"
	"monkey-go/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	Scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := Scanner.Scan()

		if !scanned {
			return
		}

		line := Scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
