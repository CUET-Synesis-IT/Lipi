package ast

import (
	"bytes"
)

type Node interface {
	TokenLiteral() string
	String() string
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
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
