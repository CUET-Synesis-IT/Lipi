package parser

import (
	"lipi/lexer"
	"lipi/token"
	"lipi/ast"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{l: l}
	
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() *ast.Program {
	return nil
}
