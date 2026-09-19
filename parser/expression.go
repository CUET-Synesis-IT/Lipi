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
	case token.TRUE, token.FALSE:
		left, err = p.parseBoolean()
	case token.IF:
		left, err = p.parseIfExpression()
	case token.FUNCTION:
		left, err = p.parseFuncExpression()
	default:
		err = fmt.Errorf("unknown expression type: %s", p.currentToken.Type)
	}
	if err != nil {
		return nil, err
	}
	for {
		if p.peekToken.Type.IsInfix() && currPrecedence < p.peekToken.Type.Precedence() {
			p.nextToken()
			left, err = p.parseInfixExpression(left)
			if err != nil {
				return nil, err
			}
			continue
		}

		if p.peekToken.Type == token.LPAREN && currPrecedence < token.CALL {
			p.nextToken()
			left, err = p.parseCallExpression(left)
			if err != nil {
				return nil, err
			}
			continue
		}

		break
	}
	return left, nil
}

func (p *Parser) parseCallExpression(left ast.Expression) (ast.Expression, error) {
	ident, ok := left.(*ast.Identifier)
	if !ok {
		return nil, fmt.Errorf("expected identifier, got %s", left.TokenLiteral())
	}

	expr := &ast.CallExpression{Token: p.currentToken, Function: ident}

	var err error
	expr.Arguments, err = p.parseArguments()
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func (p *Parser) parseArguments() ([]ast.Expression, error) {
	var args []ast.Expression
	if err := p.expectToken(token.LPAREN); err != nil {
		return nil, err
	}
	p.nextToken()

	for p.currentToken.Type != token.RPAREN {
		arg, err := p.parseExpression(token.LOWEST)
		if err != nil {
			return nil, err
		}
		args = append(args, arg)
		p.nextToken()

		if err := p.expectToken(token.COMMA); err != nil {
			break
		}
		p.nextToken()
	}

	if err := p.expectToken(token.RPAREN); err != nil {
		return nil, err
	}

	return args, nil
}

func (p *Parser) parseFuncExpression() (ast.Expression, error) {
	expr := &ast.FuncExpression{Token: p.currentToken}
	p.nextToken()

	var err error
	expr.Parameters, err = p.parseParameters()
	if err != nil {
		return nil, err
	}

	expr.Body, err = p.parseBody()
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func (p *Parser) parseIfExpression() (ast.Expression, error) {
	expr := &ast.IfExpression{Token: p.currentToken}
	p.nextToken()

	if err := p.expectToken(token.LPAREN); err != nil {
		return nil, err
	}
	p.nextToken()

	var err error
	expr.Condition, err = p.parseExpression(token.LOWEST)
	if err != nil {
		return nil, err
	}
	p.nextToken()

	if err := p.expectToken(token.RPAREN); err != nil {
		return nil, err
	}
	p.nextToken()

	expr.Consequence, err = p.parseBody()
	if err != nil {
		return nil, err
	}

	if p.peekToken.Type != token.ELSE {
		return expr, nil
	}

	p.nextToken()
	p.nextToken()

	expr.Alternative, err = p.parseBody()
	if err != nil {
		return nil, err
	}

	return expr, nil
}

func (p *Parser) parseParameters() ([]*ast.Identifier, error) {
	var params []*ast.Identifier

	if err := p.expectToken(token.LPAREN); err != nil {
		return nil, err
	}
	p.nextToken()

	for {
		param, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		params = append(params, param)
		p.nextToken()

		if err := p.expectToken(token.COMMA); err != nil {
			if p.currentToken.Type == token.RPAREN {
				break
			}
			return nil, err
		}
		p.nextToken()
	}

	if err := p.expectToken(token.RPAREN); err != nil {
		return nil, err
	}
	p.nextToken()

	return params, nil
}

func (p *Parser) parseBody() ([]ast.Statement, error) {
	var statements []ast.Statement

	if err := p.expectToken(token.LBRACE); err != nil {
		return nil, err
	}
	p.nextToken()

	for p.currentToken.Type != token.RBRACE {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		statements = append(statements, stmt)
	}

	if err := p.expectToken(token.RBRACE); err != nil {
		return nil, err
	}

	return statements, nil
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

func (p *Parser) parseBoolean() (*ast.BooleanLiteral, error) {
	if p.currentToken.Type != token.TRUE && p.currentToken.Type != token.FALSE {
		return nil, fmt.Errorf("unknown boolean: %s", p.currentToken.Type)
	}
	return &ast.BooleanLiteral{Token: p.currentToken, Value: p.currentToken.Type == token.TRUE}, nil
}
