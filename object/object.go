package object

import "fmt"

type ObjectType string

const (
	IntType         ObjectType = "INT"
	BoolType        ObjectType = "BOOL"
	NilType         ObjectType = "NIL"
	ReturnValueType ObjectType = "RETURN_VALUE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Int struct {
	Value int64
}

func (i *Int) Type() ObjectType {
	return IntType
}
func (i *Int) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

type Bool struct {
	Value bool
}

func (b *Bool) Type() ObjectType {
	return BoolType
}
func (b *Bool) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

type Nil struct{}

func (n *Nil) Type() ObjectType {
	return NilType
}
func (n *Nil) Inspect() string {
	return "nil"
}

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() ObjectType {
	return ReturnValueType
}
func (r *ReturnValue) Inspect() string {
	return r.Value.Inspect()
}
