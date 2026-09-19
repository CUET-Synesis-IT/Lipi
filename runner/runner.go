package runner

import (
	"fmt"
	"os"

	"lipi/eval"
	"lipi/lexer"
	"lipi/object"
	"lipi/parser"
)

func RunFile(filename string) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	input := string(content)

	l := lexer.New(input)
	p := parser.New(l)

	program, err := p.Parse()
	if err != nil {
		return err
	}

	env := object.NewEnvironment()

	result, err := eval.Eval(program, env)
	if err != nil {
		return err
	}

	if result != nil {
		fmt.Println(result.Inspect())
	}

	return nil
}