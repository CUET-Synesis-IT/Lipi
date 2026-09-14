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
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "exit" {
			return
		}

		input := line

		for strings.Count(input, "{") > strings.Count(input, "}") {
			if !scanner.Scan() {
				break
			}
			input += "\n" + scanner.Text()
		}

		l := lexer.New(input)

		fmt.Fprintln(out)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "Token: %v\n", tok)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(out, "Error reading input: %v\n", err)
	}
}
