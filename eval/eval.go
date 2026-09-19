package eval

import (
	"fmt"
	"lipi/ast"
	"lipi/object"
	"lipi/token"
)

var (
	TRUE  = &object.Bool{Value: true}
	FALSE = &object.Bool{Value: false}
	NIL   = &object.Nil{}
)

func Eval(node ast.Node) (object.Object, error) {
	switch node := node.(type) {
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)
	case *ast.InfixExpression:
		return evalInfixExpression(node)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node)
	case *ast.IntLiteral:
		return &object.Int{Value: node.Value}, nil
	case *ast.BooleanLiteral:
		return evalBoolLiteral(node)
	default:
		return NIL, nil
	}
}

func evalInfixExpression(node *ast.InfixExpression) (object.Object, error) {
	left, err := Eval(node.Left)
	if err != nil {
		return NIL, err
	}
	right, err := Eval(node.Right)
	if err != nil {
		return NIL, err
	}
	switch node.Token.Type {
	case token.PLUS:
		return evalPlusOperator(left, right)
	case token.MINUS:
		return evalMinusInfixOperator(left, right)
	case token.ASTERISK:
		return evalMultiplyOperator(left, right)
	case token.SLASH:
		return evalDivideOperator(left, right)
	case token.LT:
		return evalLessThanOperator(left, right)
	case token.GT:
		return evalGreaterThanOperator(left, right)
	case token.EQ:
		return evalEqualOperator(left, right)
	case token.NEQ:
		return evalNotEqualOperator(left, right)
	default:
		return NIL, nil
	}
}

func evalNotEqualOperator(left, right object.Object) (object.Object, error) {
	eq, err := evalEqualOperator(left, right)
	if err != nil {
		return NIL, err
	}

	neq, err := evalBangOperator(eq)
	if err != nil {
		return NIL, err
	}
	return neq, nil
}

func evalEqualOperator(left, right object.Object) (object.Object, error) {
	if left.Type() != right.Type() {
		return NIL, nil
	}
	switch left.(type) {
	case *object.Int:
		intLeft := left.(*object.Int)
		intRight := right.(*object.Int)
		if intLeft.Value == intRight.Value {
			return TRUE, nil
		}
		return FALSE, nil
	case *object.Bool:
		boolLeft := left.(*object.Bool)
		boolRight := right.(*object.Bool)
		if boolLeft.Value == boolRight.Value {
			return TRUE, nil
		}
		return FALSE, nil
	case *object.Nil:
		return TRUE, nil
	default:
		return FALSE, nil
	}
}

func evalGreaterThanOperator(left, right object.Object) (object.Object, error) {
	if left == NIL || right == NIL {
		return NIL, fmt.Errorf("unexpected nil operand")
	}
	intLeft, ok := left.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	intRight, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	if intLeft.Value > intRight.Value {
		return TRUE, nil
	}
	return FALSE, nil
}

func evalLessThanOperator(left, right object.Object) (object.Object, error) {
	if left == NIL || right == NIL {
		return NIL, fmt.Errorf("unexpected nil operand")
	}
	intLeft, ok := left.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	intRight, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	if intLeft.Value < intRight.Value {
		return TRUE, nil
	}
	return FALSE, nil
}

func evalMinusInfixOperator(left, right object.Object) (object.Object, error) {
	if left == NIL || right == NIL {
		return NIL, fmt.Errorf("unexpected nil operand")
	}
	intLeft, ok := left.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	intRight, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	result := intLeft.Value - intRight.Value
	return &object.Int{Value: result}, nil
}

func evalPlusOperator(left, right object.Object) (object.Object, error) {
	if left == NIL || right == NIL {
		return NIL, fmt.Errorf("unexpected nil operand")
	}
	intLeft, ok := left.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	intRight, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	result := intLeft.Value + intRight.Value
	return &object.Int{Value: result}, nil
}

func evalMultiplyOperator(left, right object.Object) (object.Object, error) {
	if left == NIL || right == NIL {
		return NIL, fmt.Errorf("unexpected nil operand")
	}
	intLeft, ok := left.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	intRight, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	result := intLeft.Value * intRight.Value
	return &object.Int{Value: result}, nil
}

func evalDivideOperator(left, right object.Object) (object.Object, error) {
	if left == NIL || right == NIL {
		return NIL, fmt.Errorf("unexpected nil operand")
	}
	intLeft, ok := left.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	intRight, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unexpected non-int operand")
	}
	result := intLeft.Value / intRight.Value
	return &object.Int{Value: result}, nil
}

func evalPrefixExpression(node *ast.PrefixExpression) (object.Object, error) {
	right, err := Eval(node.Right)
	if err != nil {
		return NIL, err
	}
	if right == nil {
		return NIL, nil
	}

	switch node.Token.Type {
	case token.BANG:
		return evalBangOperator(right)
	case token.MINUS:
		return evalMinusPrefixOperator(right)
	default:
		return NIL, nil
	}
}

func evalMinusPrefixOperator(right object.Object) (object.Object, error) {
	if right == NIL {
		return NIL, nil
	}
	intVal, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unsupported operand type for -: %s", right.Inspect())
	}
	return &object.Int{Value: -intVal.Value}, nil
}

func evalBangOperator(right object.Object) (object.Object, error) {
	switch right := right.(type) {
	case *object.Bool:
		if right.Value {
			return FALSE, nil
		}
		return TRUE, nil
	case *object.Nil:
		return TRUE, nil
	case *object.Int:
		if right.Value != 0 {
			return FALSE, nil
		}
		return TRUE, nil
	default:
		return NIL, nil
	}
}

func evalBoolLiteral(node *ast.BooleanLiteral) (object.Object, error) {
	if node.Value {
		return TRUE, nil
	}
	return FALSE, nil
}

func evalStatements(statements []ast.Statement) (object.Object, error) {
	var result object.Object
	var err error

	for _, stmt := range statements {
		result, err = Eval(stmt)
		if err != nil {
			return NIL, err
		}
	}
	return result, nil
}
