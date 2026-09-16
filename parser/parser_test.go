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
