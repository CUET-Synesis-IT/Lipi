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

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil, nil
	}
}

func (p *Parser) parseLetStatement() (*ast.LetStatement, error) {
	stmt := &ast.LetStatement{Token: p.currentToken}
	p.nextToken()

	var err error
	stmt.Name, err = p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	p.nextToken()

	if err := p.expectToken(token.ASSIGN); err != nil {
		return nil, err
	}
	p.nextToken()

	// TODO: fix later
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	p.nextToken()

	return stmt, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.currentToken}
	p.nextToken()

	// TODO: fix later
	for p.currentToken.Type != token.SEMICOLON {
		p.nextToken()
	}
	p.nextToken()

	return stmt, nil
}

func (p *Parser) parseIdentifier() (*ast.Identifier, error) {
	if err := p.expectToken(token.IDENT); err != nil {
		return nil, err
	}
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}, nil
}

func (p *Parser) expectToken(t token.TokenType) error {
	if p.currentToken.Type != t {
		return fmt.Errorf("expected %s, got %s", t, p.currentToken.Type)
	}
	return nil
}