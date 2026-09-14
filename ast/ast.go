package ast

import "lipi/token"

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	ExpressionNode()
}

type Statement interface {
	Node
	StatementNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	return "ROOT"
}

type LetStatement struct {
	Token     token.Token
	Name      *Identifier
	Value     Expression
}

func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) ExpressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}