package main

import (
	"fmt"
	"lipi/lexer"
	"lipi/token"
	"os"
)

func main() {
	content, _ := os.ReadFile("example/01.lipi")
	l := lexer.New(string(content))
	for {
		tok := l.NextToken()
		if tok.Type == token.EOF {
			break
		}
		fmt.Printf("tok: %v\n", tok)
	}
}
