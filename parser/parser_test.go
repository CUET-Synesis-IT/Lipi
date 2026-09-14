package parser

import (
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

	program := p.Parse()
	if program == nil {
		t.Fatal("Parse() returned nil")
	}
	if len(program.Statements) != 3 {
		t.Fatalf("expected 3 statements, got %d", len(program.Statements))
	}
	expectedIdent := []string{"সংখ্যা", "ক", "খ"}
	expectedValue := []int{5, 15, 15}
	for i, stmt := range program.Statements {
		testLetStatement(t, stmt, expectedIdent[i], expectedValue[i])
	}
}

func testLetStatement(t *testing.T, stmt ast.Statement, expectedIdent string, expectedValue int) {
	if stmt == nil {
		t.Fatal("stmt is nil")
	}
	if stmt.(*ast.LetStatement).Name.Value != expectedIdent {
		t.Fatalf("expected %s, got %s", expectedIdent, stmt.(*ast.LetStatement).Name.Value)
	}
	if stmt.(*ast.LetStatement).Value == nil {
		t.Fatal("value is nil")
	}
	// if stmt.(*ast.LetStatement).Value.(*ast.IntegerLiteral).Value != expectedValue {
	// 	t.Fatalf("expected %d, got %d", expectedValue, stmt.(*ast.LetStatement).Value.(*ast.IntegerLiteral).Value)
	// }
}