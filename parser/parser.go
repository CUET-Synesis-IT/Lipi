package parser

import (
	"fmt"
	"lipi/ast"
	"lipi/lexer"
	"lipi/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekToken    token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Parse() (*ast.Program, error) {
	program := &ast.Program{}
	program.Statements = make([]ast.Statement, 0)

	for p.currentToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return program, err
		}
		program.Statements = append(program.Statements, stmt)
	}
	return program, nil
}

func (p *Parser) expectToken(t token.TokenType) error {
	if p.currentToken.Type != t {
		return fmt.Errorf("expected %s, got %s", t, p.currentToken)
	}
	return nil
}
