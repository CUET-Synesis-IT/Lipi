package parser

import (
	"lipi/ast"
	"lipi/token"
)

func (p *Parser) parseStatement() (ast.Statement, error) {
	switch p.currentToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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

	stmt.Value, err = p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}
	p.nextToken()

	return stmt, nil
}

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, error) {
	stmt := &ast.ReturnStatement{Token: p.currentToken}
	p.nextToken()

	if p.currentToken.Type == token.SEMICOLON {
		for p.peekToken.Type == token.SEMICOLON {
			p.nextToken()
		}
		p.nextToken()
		return stmt, nil
	}

	var err error
	stmt.Value, err = p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}
	p.nextToken()

	return stmt, nil
}

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, error) {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	var err error
	stmt.Expression, err = p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}

	for p.peekToken.Type == token.SEMICOLON {
		p.nextToken()
	}
	p.nextToken()

	return stmt, nil
}
