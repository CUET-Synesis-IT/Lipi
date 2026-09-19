package eval

import (
	"fmt"
	"lipi/lexer"
	"lipi/object"
	"lipi/parser"
	"testing"
)

func TestIntEval(t *testing.T) {
	testCases := []struct {
		input    string
		expected int64
	}{
		{"৫;", 5},
		{"০;", 0},
		{"১০;", 10},
		{"-৫;", -5},
		{"-০;", 0},
		{"-১০;", -10},
		{"৫ + ৫ + ৫ + ৫ - ১০", 10},
		{"২ * ২ * ২ * ২ * ২", 32},
		{"-৫০ + ১০০ + -৫০", 0},
		{"৫ * ২ + ১০", 20},
		{"৫ + ২ * ১০", 25},
		{"২০ + ২ * -১০", 0},
		{"৫০ / ২ * ২ + ১০", 60},
		{"২ * (৫ + ১০)", 30},
		{"৩ * ৩ * ৩ + ১০", 37},
		{"৩ * (৩ * ৩) + ১০", 37},
		{"(৫ + ১০ * ২ + ১৫ / ৩) * ২ + -১০", 50},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := parser.New(l)
		program, err := p.Parse()
		if err != nil {
			t.Errorf("parse error: %v", err)
		}
		result, err := Eval(program)
		if err != nil {
			t.Errorf("eval error: %v", err)
		}

		if result == nil {
			t.Errorf("result is nil")
		}

		intResult, ok := result.(*object.Int)
		if !ok {
			t.Errorf("result is not an int")
		}

		if intResult.Value != tc.expected {
			t.Errorf("expected %d, got %d", tc.expected, intResult.Value)
		}
	}
}

func TestBoolEval(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"সত্য;", true},
		{"মিথ্যা;", false},
		{"!সত্য;", false},
		{"!মিথ্যা;", true},
		{"!৫;", false},
		{"!-৫;", false},
		{"!০;", true},
		{"!!৫;", true},
		{"!!০;", false},
		{"১ < ২", true},
		{"১ > ২", false},
		{"১ < ১", false},
		{"১ > ১", false},
		{"১ == ১", true},
		{"১ != ১", false},
		{"১ == ২", false},
		{"১ != ২", true},
		{"সত্য == সত্য", true},
		{"মিথ্যা == মিথ্যা", true},
		{"সত্য == মিথ্যা", false},
		{"সত্য != মিথ্যা", true},
		{"মিথ্যা != সত্য", true},
		{"(১ < ২) == সত্য", true},
		{"(১ < ২) == মিথ্যা", false},
		{"(১ > ২) == সত্য", false},
		{"(১ > ২) == মিথ্যা", true},
	}

	for _, tc := range testCases {
		l := lexer.New(tc.input)
		p := parser.New(l)
		program, err := p.Parse()
		fmt.Printf("program: %v\n", program)
		if err != nil {
			t.Errorf("parse error: %v", err)
		}
		result, err := Eval(program)
		if err != nil {
			t.Errorf("eval error: %v", err)
		}

		if result == nil {
			t.Errorf("result is nil")
		}

		intResult, ok := result.(*object.Bool)
		if !ok {
			t.Errorf("result is not a bool")
		}

		if intResult.Value != tc.expected {
			t.Errorf("expected %t, got %t", tc.expected, intResult.Value)
		}
	}
}
