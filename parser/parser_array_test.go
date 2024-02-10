package parser

import (
	"testing"

	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
)

func TestParsingArrayLiterals(t *testing.T) {
	input := "[1, 2 * 2, 3 + 3]"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
	}

	if len(array.Elements) != 3 {
		t.Fatalf("len(array.Elements) not 3. got=%d", len(array.Elements))
	}

	testIntegerLiteral(t, array.Elements[0], 1)
	testInfixExpression(t, array.Elements[1], 2, "*", 2)
	testInfixExpression(t, array.Elements[2], 3, "+", 3)
}

func TestParsingEmptyArrayLiterals(t *testing.T) {
	input := "[]"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	array, ok := stmt.Expression.(*ast.ArrayLiteral)
	if !ok {
		t.Fatalf("exp not ast.ArrayLiteral. got=%T", stmt.Expression)
	}

	if len(array.Elements) != 0 {
		t.Errorf("len(array.Elements) not 0. got=%d", len(array.Elements))
	}
}
