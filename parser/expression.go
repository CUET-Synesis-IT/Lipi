package parser

import (
	"fmt"
	"lipi/ast"
	"lipi/token"
)

func (p *Parser) parseExpression(currPrecedence int) (ast.Expression, error) {
	var left ast.Expression
	var err error

	switch p.currentToken.Type {
	case token.INT:
		left, err = p.parseIntLiteral()
	case token.IDENT:
		left, err = p.parseIdentifier()
	case token.MINUS, token.BANG:
		left, err = p.parsePrefixExpression()
	case token.LPAREN:
		left, err = p.parseGroupExpression()
	default:
		err = fmt.Errorf("unknown expression type: %s", p.currentToken.Type)
	}
	if err != nil {
		return nil, err
	}
	for p.peekToken.Type.IsInfix() && currPrecedence < p.peekToken.Type.Precedence() {
		p.nextToken()
		left, err = p.parseInfixExpression(left)
		if err != nil {
			return nil, err
		}
	}
	return left, nil
}

func (p *Parser) parseIdentifier() (*ast.Identifier, error) {
	if err := p.expectToken(token.IDENT); err != nil {
		return nil, err
	}
	return &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}, nil
}

func (p *Parser) parseIntLiteral() (*ast.IntLiteral, error) {
	lit := &ast.IntLiteral{Token: p.currentToken}
	if err := p.expectToken(token.INT); err != nil {
		return nil, err
	}
	lit.Value = parseBengaliDigits(p.currentToken.Literal)
	return lit, nil
}

func (p *Parser) parsePrefixExpression() (*ast.PrefixExpression, error) {
	if p.currentToken.Type != token.MINUS && p.currentToken.Type != token.BANG {
		return nil, fmt.Errorf("unknown prefix operator: %s", p.currentToken.Type)
	}
	expr := &ast.PrefixExpression{Token: p.currentToken, Operator: p.currentToken.Literal}
	p.nextToken()

	var err error
	expr.Right, err = p.parseExpression(token.PREFIX)
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseInfixExpression(left ast.Expression) (*ast.InfixExpression, error) {
	expr := &ast.InfixExpression{Token: p.currentToken, Operator: p.currentToken.Literal, Left: left}
	p.nextToken()
	var err error
	expr.Right, err = p.parseExpression(expr.Token.Type.Precedence())
	if err != nil {
		return nil, err
	}
	return expr, nil
}

func (p *Parser) parseGroupExpression() (ast.Expression, error) {
	p.nextToken()

	expr, err := p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}
	p.nextToken()

	if err := p.expectToken(token.RPAREN); err != nil {
		return nil, err
	}
	return expr, nil
}
