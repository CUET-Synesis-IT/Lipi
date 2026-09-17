package repl

import (
	"bufio"
	"fmt"
	"io"
	"lipi/lexer"
	"lipi/parser"
	"lipi/token"
	"strings"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, username string) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)

		if !scanner.Scan() {
			break
		}

		line := strings.TrimSpace(scanner.Text())

		if line == "exit" {
			fmt.Fprintf(out, "Goodbye, %s!\n", username)
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
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "Token: %v\n", tok)
		}

		l = lexer.New(input)
		p := parser.New(l)
		program, err := p.Parse()
		if err != nil {
			fmt.Fprintln(out, err)
		}
		fmt.Fprintln(out, program)
		fmt.Fprintln(out)

	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(out, "Error reading input: %v\n", err)
	}
}
