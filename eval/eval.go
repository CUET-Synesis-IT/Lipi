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

func Eval(node ast.Node, env *object.Environment) (object.Object, error) {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)
	case *ast.LetStatement:
		return evalLetStatement(node, env)
	case *ast.ReturnStatement:
		return evalReturnStatement(node, env)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.IfExpression:
		return evalIfExpression(node, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.InfixExpression:
		return evalInfixExpression(node, env)
	case *ast.PrefixExpression:
		return evalPrefixExpression(node, env)
	case *ast.IntLiteral:
		return &object.Int{Value: node.Value}, nil
	case *ast.BooleanLiteral:
		return evalBoolLiteral(node, env)
	default:
		return NIL, nil
	}
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) (object.Object, error) {
	value, ok := env.Get(node.Value)
	if !ok {
		return NIL, fmt.Errorf("identifier not found: %s", node.Value)
	}
	return value, nil
}

func evalLetStatement(node *ast.LetStatement, env *object.Environment) (object.Object, error) {
	value, err := Eval(node.Value, env)
	if err != nil {
		return NIL, err
	}
	env.Set(node.Name.Value, value)
	return NIL, nil
}

func evalProgram(node *ast.Program, env *object.Environment) (object.Object, error) {
	out, err := evalStatements(node.Statements, env)
	if err != nil {
		return NIL, err
	}
	if out, ok := out.(*object.ReturnValue); ok {
		return out.Value, nil
	}
	return out, nil
}

func evalReturnStatement(node *ast.ReturnStatement, env *object.Environment) (object.Object, error) {
	if node.Value == nil {
		return &object.ReturnValue{Value: NIL}, nil
	}
	value, err := Eval(node.Value, env)
	if err != nil {
		return NIL, err
	}
	return &object.ReturnValue{Value: value}, nil
}

func evalIfExpression(node *ast.IfExpression, env *object.Environment) (object.Object, error) {
	condition, err := Eval(node.Condition, env)
	if err != nil {
		return NIL, err
	}
	if isTruthy(condition) {
		return evalStatements(node.Consequence, env)
	}
	return evalStatements(node.Alternative, env)
}

func evalInfixExpression(node *ast.InfixExpression, env *object.Environment) (object.Object, error) {
	left, err := Eval(node.Left, env)
	if err != nil {
		return NIL, err
	}
	right, err := Eval(node.Right, env)
	if err != nil {
		return NIL, err
	}
	switch node.Token.Type {
	case token.PLUS:
		return evalPlusOperator(left, right, env)
	case token.MINUS:
		return evalMinusInfixOperator(left, right, env)
	case token.ASTERISK:
		return evalMultiplyOperator(left, right, env)
	case token.SLASH:
		return evalDivideOperator(left, right, env)
	case token.LT:
		return evalLessThanOperator(left, right, env)
	case token.GT:
		return evalGreaterThanOperator(left, right, env)
	case token.EQ:
		return evalEqualOperator(left, right, env)
	case token.NEQ:
		return evalNotEqualOperator(left, right, env)
	default:
		return NIL, nil
	}
}

func evalNotEqualOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
	eq, err := evalEqualOperator(left, right, env)
	if err != nil {
		return NIL, err
	}

	neq, err := evalBangOperator(eq, env)
	if err != nil {
		return NIL, err
	}
	return neq, nil
}

func evalEqualOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
	if left.Type() != right.Type() {
		return FALSE, nil
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

func evalGreaterThanOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
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

func evalLessThanOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
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

func evalMinusInfixOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
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

func evalPlusOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
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

func evalMultiplyOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
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

func evalDivideOperator(left, right object.Object, env *object.Environment) (object.Object, error) {
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
	if intRight.Value == 0 {
		return NIL, fmt.Errorf("division by zero")
	}
	result := intLeft.Value / intRight.Value
	return &object.Int{Value: result}, nil
}

func evalPrefixExpression(node *ast.PrefixExpression, env *object.Environment) (object.Object, error) {
	right, err := Eval(node.Right, env)
	if err != nil {
		return NIL, err
	}
	if right == nil {
		return NIL, nil
	}

	switch node.Token.Type {
	case token.BANG:
		return evalBangOperator(right, env)
	case token.MINUS:
		return evalMinusPrefixOperator(right, env)
	default:
		return NIL, nil
	}
}

func evalMinusPrefixOperator(right object.Object, env *object.Environment) (object.Object, error) {
	if right == NIL {
		return NIL, nil
	}
	intVal, ok := right.(*object.Int)
	if !ok {
		return NIL, fmt.Errorf("unsupported operand type for -: %s", right.Inspect())
	}
	return &object.Int{Value: -intVal.Value}, nil
}

func evalBangOperator(right object.Object, env *object.Environment) (object.Object, error) {
	if isTruthy(right) {
		return FALSE, nil
	}
	return TRUE, nil
}

func evalBoolLiteral(node *ast.BooleanLiteral, env *object.Environment) (object.Object, error) {
	if node.Value {
		return TRUE, nil
	}
	return FALSE, nil
}

func evalStatements(statements []ast.Statement, env *object.Environment) (object.Object, error) {
	var result object.Object = NIL
	var err error

	for _, stmt := range statements {
		result, err = Eval(stmt, env)
		if err != nil {
			return NIL, err
		}
		if _, ok := result.(*object.ReturnValue); ok {
			return result, nil
		}
	}
	return result, nil
}
