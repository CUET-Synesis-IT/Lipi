package parser

import (
	"fmt"
	"lipi/ast"
	"lipi/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	ধর সংখ্যা = ৫;
	ধর ক = ১৫;
	ধর খ = ১৫;
	`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if program == nil {
		t.Fatal("Parse() returned nil")
	}
	if len(program.Statements) != 3 {
		fmt.Printf("program.Statements: %v\n", program.Statements)
		t.Fatalf("expected 3 statements, got %d", len(program.Statements))
	}
	expectedIdent := []string{"সংখ্যা", "ক", "খ"}
	expectedValue := []int64{5, 15, 15}
	for i, stmt := range program.Statements {
		testLetStatement(t, stmt, expectedIdent[i], expectedValue[i])
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, expectedIdent string, expectedValue int64) {
	if stmt == nil {
		t.Fatal("stmt is nil")
	}
	if stmt.(*ast.LetStatement).Name.Value != expectedIdent {
		t.Fatalf("expected %s, got %s", expectedIdent, stmt.(*ast.LetStatement).Name.Value)
	}
	if stmt.(*ast.LetStatement).Value == nil {
		t.Fatal("value is nil")
	}
	if stmt.(*ast.LetStatement).Value.(*ast.IntLiteral).Value != expectedValue {
		t.Fatalf("expected %d, got %d", expectedValue, stmt.(*ast.LetStatement).Value.(*ast.IntLiteral).Value)
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
    ফেরাও ক + খ;
	ফেরাও ১৫;
	`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}
	if program == nil {
		t.Fatal("Parse() returned nil")
	}
	if len(program.Statements) != 2 {
		t.Fatalf("expected 2 statements, got %d", len(program.Statements))
	}
	for i, stmt := range program.Statements {
		if stmt == nil {
			t.Fatalf("stmt is nil at index %d", i)
		}
		if stmt.(*ast.ReturnStatement) == nil {
			t.Fatalf("stmt is not a ReturnStatement at index %d", i)
		}
	}
}

func TestParseBengaliDigits(t *testing.T) {
	testCases := []struct {
		input  string
		output int64
	}{
		{"১৫", 15},
		{"১০০", 100},
		{"০", 0},
		{"১০০০", 1000},
	}

	for _, tc := range testCases {
		result := parseBengaliDigits(tc.input)
		if result != tc.output {
			t.Fatalf("expected %d, got %d", tc.output, result)
		}
	}
}

func TestExpression(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"৫;", "5;"},
		{"৫ + ৫;", "(5 + 5);"},
		{"৫ != ৫;", "(5 != 5);"},
		{"৫ + ৫ * ২;", "(5 + (5 * 2));"},
		{"৫ * ২ + ৫;", "((5 * 2) + 5);"},
		{"-৫ * ২;", "((-5) * 2);"},
		{"৫ * -২;", "(5 * (-2));"},
		{"(৫ + ৫) * ২;", "((5 + 5) * 2);"},
		{"৫ - ২;", "(5 - 2);"},
		{"৫ / ২;", "(5 / 2);"},
		{"৫ + ৫ + ৫;", "((5 + 5) + 5);"},
		{"৫ * ২ * ৩;", "((5 * 2) * 3);"},
		{"(৫ + ৫) * (২ - ১);", "((5 + 5) * (2 - 1));"},
		{"৫ + ২ * ৩ - ১;", "((5 + (2 * 3)) - 1);"},
		{"সত্য;", "true;"},
		{"মিথ্যা;", "false;"},
		{"মিথ্যা == সত্য;", "(false == true);"},
		{"৫ + ৫ == ১০;", "((5 + 5) == 10);"},
		{"৫ != ৩ + ২;", "(5 != (3 + 2));"},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := New(l)
		program, err := p.Parse()
		if err != nil {
			t.Fatal(err)
		}
		if program == nil {
			t.Fatal("Parse() returned nil")
		}
		result := program.String()
		if result != tc.output {
			t.Fatalf("expected %s, got %s", tc.output, result)
		}
	}
}

func TestIfExpression(t *testing.T) {
	input := `যদি (খ < ৫) {
		খ;
	}`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if program == nil {
		t.Fatal("Parse() returned nil")
	}

	if len(program.Statements) != 1 {
		t.Fatalf(
			"expected 1 statement, got %d",
			len(program.Statements),
		)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf(
			"expected ExpressionStatement, got %T",
			program.Statements[0],
		)
	}

	expr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf(
			"expected IfExpression, got %T",
			stmt.Expression,
		)
	}

	if expr.Condition == nil {
		t.Fatal("if condition is nil")
	}

	if expr.Consequence == nil {
		t.Fatal("if consequence is nil")
	}

	if expr.Alternative != nil {
		t.Fatal("expected no alternative")
	}

	if len(expr.Consequence) != 1 {
		t.Fatalf(
			"expected 1 consequence statement, got %d",
			len(expr.Consequence),
		)
	}
}

func TestIfElseExpression(t *testing.T) {
	input := `যদি (খ < ৫) {
		খ;
	} নাহলে {
		১০;
	}`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf(
			"expected ExpressionStatement, got %T",
			program.Statements[0],
		)
	}

	expr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf(
			"expected IfExpression, got %T",
			stmt.Expression,
		)
	}

	if expr.Alternative == nil {
		t.Fatal("expected alternative block")
	}

	if len(expr.Alternative) != 1 {
		t.Fatalf(
			"expected 1 alternative statement, got %d",
			len(expr.Alternative),
		)
	}
}

func TestFuncExpression(t *testing.T) {
	input := `ধর যোগ_করো = ফাঙ্কশন(ক, খ) {
	    ফেরাও ক + খ;
	};`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if len(program.Statements) != 1 {
		t.Fatalf(
			"expected 1 statement, got %d",
			len(program.Statements),
		)
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf(
			"expected ExpressionStatement, got %T",
			program.Statements[0],
		)
	}

	expr, ok := stmt.Value.(*ast.FuncExpression)
	if !ok {
		t.Fatalf(
			"expected FuncExpression, got %T",
			stmt.Value,
		)
	}

	if expr.Parameters == nil {
		t.Fatal("expected parameters")
	}

	if len(expr.Parameters) != 2 {
		t.Fatalf(
			"expected 2 parameters, got %d",
			len(expr.Parameters),
		)
	}

	if expr.Body == nil {
		t.Fatal("expected body")
	}

	if len(expr.Body) != 1 {
		t.Fatalf(
			"expected 1 body statement, got %d",
			len(expr.Body),
		)
	}

	if expr.Body[0].String() != "RETURN (ক + খ);" {
		t.Fatalf(
			"expected body statement to be 'RETURN (ক + খ);', got %s",
			expr.Body[0].String(),
		)
	}
}

func TestCallExpression(t *testing.T) {
	input := `ধর যোগফল = যোগ(ক, খ + যোগ(ক, খ))`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if len(program.Statements) != 1 {
		t.Fatalf(
			"expected 1 statement, got %d",
			len(program.Statements),
		)
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf(
			"expected LetStatement, got %T",
			program.Statements[0],
		)
	}

	if stmt.Value.String() != "যোগ(ক, (খ + যোগ(ক, খ)))" {
		t.Fatalf(
			"expected value to be 'যোগ(ক, (খ + যোগ(ক, খ)))', got %s",
			stmt.Value.String(),
		)
	}
}

func TestFuncExpressionNoParams(t *testing.T) {
	input := `ধর পাই = ফাঙ্কশন() { ফেরাও ৫; };`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.LetStatement)
	if !ok {
		t.Fatalf("expected LetStatement, got %T", program.Statements[0])
	}

	expr, ok := stmt.Value.(*ast.FuncExpression)
	if !ok {
		t.Fatalf("expected FuncExpression, got %T", stmt.Value)
	}

	if len(expr.Parameters) != 0 {
		t.Fatalf("expected 0 parameters, got %d", len(expr.Parameters))
	}

	if len(expr.Body) != 1 {
		t.Fatalf("expected 1 body statement, got %d", len(expr.Body))
	}
}

func TestNestedIfStatements(t *testing.T) {
	input := `যদি (সত্য) {
		যদি (সত্য) {
			ফেরাও ২;
		}
		ফেরাও ১;
	}`

	l := lexer.New(input)
	p := New(l)

	program, err := p.Parse()
	if err != nil {
		t.Fatal(err)
	}

	if len(program.Statements) != 1 {
		t.Fatalf("expected 1 statement, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expected ExpressionStatement, got %T", program.Statements[0])
	}

	ifExpr, ok := stmt.Expression.(*ast.IfExpression)
	if !ok {
		t.Fatalf("expected IfExpression, got %T", stmt.Expression)
	}

	if len(ifExpr.Consequence) != 2 {
		t.Fatalf("expected 2 statements in consequence, got %d", len(ifExpr.Consequence))
	}
}
