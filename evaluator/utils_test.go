package evaluator

import (
	"testing"

	"monkey-interpreter/lexer"
	"monkey-interpreter/object"
	"monkey-interpreter/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testNullObject(t *testing.T, obj object.Object) bool {
	if obj != NULL {
		t.Errorf("object is not NULL. received=%T (%+v)", obj, obj)
		return false
	}

	return true
}
