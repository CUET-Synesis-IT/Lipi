package repl

import (
	"bufio"
	"fmt"
	"io"
	"lipi/lexer"
	"lipi/token"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		if !scanner.Scan() {
			return
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "exit" {
			fmt.Fprintln(out, "Goodbye!")
			return
		}

		input := line

		for strings.Count(input, "{") > strings.Count(input, "}") {
			if !scanner.Scan() {
				return
			}
			input += "\n" + scanner.Text()
		}

		l := lexer.New(input)

		fmt.Fprintln(out)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "Token: %v\n", tok)
		}
	}
}
