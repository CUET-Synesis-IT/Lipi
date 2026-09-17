package ast

import (
	"bytes"
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

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence []Statement
	Alternative []Statement
}

func (ie *IfExpression) ExpressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("IF ")
	out.WriteString(ie.Condition.String())
	out.WriteString(" { ")
	for _, stmt := range ie.Consequence {
		out.WriteString("\n  ")
		out.WriteString(stmt.String())
	}
	if ie.Alternative == nil {
		out.WriteString("\n}")
	} else {
		out.WriteString("\n} ELSE {")
		for _, stmt := range ie.Alternative {
			out.WriteString("\n  ")
			out.WriteString(stmt.String())
		}
		out.WriteString("\n}")
	}
	return out.String()
}

type FuncExpression struct {
	Token      token.Token
	Parameters []*Identifier
	Body       []Statement
}

func (fe *FuncExpression) ExpressionNode() {}
func (fe *FuncExpression) TokenLiteral() string {
	return fe.Token.Literal
}
func (fe *FuncExpression) String() string {
	var out bytes.Buffer
	out.WriteString("FUNC ")
	for _, param := range fe.Parameters {
		out.WriteString(param.String())
		out.WriteString(" ")
	}
	out.WriteString("{\n")
	for _, stmt := range fe.Body {
		out.WriteString("  ")
		out.WriteString(stmt.String())
		out.WriteString("\n")
	}
	out.WriteString("}")
	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  *Identifier
	Arguments []Expression
}

func (ce *CallExpression) ExpressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	for i, arg := range ce.Arguments {
		out.WriteString(arg.String())
		if i < len(ce.Arguments)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(")")
	return out.String()
}
