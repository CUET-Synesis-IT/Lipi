package ast

import (
	"bytes"
	"lipi/token"
)

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString("LET ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	} else {
		out.WriteString("todo")
	}
	out.WriteString(";")
	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (rs *ReturnStatement) StatementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString("RETURN ")
	if rs.Value != nil {
		out.WriteString(rs.Value.String())
	} else {
		out.WriteString("todo")
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) StatementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (es *ExpressionStatement) String() string {
	var out bytes.Buffer
	if es.Expression != nil {
		out.WriteString(es.Expression.String())
	} else {
		out.WriteString("todo")
	}
	out.WriteString(";")
	return out.String()
}
