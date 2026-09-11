package repl

import (
	"bufio"
	"fmt"
	"io"
	"lipi/lexer"
	"lipi/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		
		fmt.Printf("\n")
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("Token: %v\n", tok)
		}
	}
}
