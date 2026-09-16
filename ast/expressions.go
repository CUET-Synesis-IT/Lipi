package ast

import (
	"fmt"
	"lipi/token"
)

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type IntLiteral struct {
	Token token.Token
	Value int64
}

func (i *IntLiteral) ExpressionNode() {}
func (i *IntLiteral) TokenLiteral() string {
	return i.Token.Literal
}
func (i *IntLiteral) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (p *PrefixExpression) ExpressionNode() {}
func (p *PrefixExpression) TokenLiteral() string {
	return p.Token.Literal
}
func (p *PrefixExpression) String() string {
	return fmt.Sprintf("(%s%s)", p.Operator, p.Right)
}

type InfixExpression struct {
	Token    token.Token
	Operator string
	Left     Expression
	Right    Expression
}

func (i *InfixExpression) ExpressionNode() {}
func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}
func (i *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left, i.Operator, i.Right)
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) ExpressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}
