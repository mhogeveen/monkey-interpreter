package evaluator

import (
	"testing"

	"monkey-interpreter/object"
)

func TestBuiltinFunctions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{`len("")`, 0},
		{`len("four")`, 4},
		{`len("hello world")`, 11},
		{`len(1)`, "argument to 'len' not supported, received INTEGER"},
		{`len("one", "two")`, "wrong number of arguments. received=2, want=1"},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {

		case int:
			testIntegerObject(t, evaluated, int64(expected))

		case string:
			errObj, ok := evaluated.(*object.Error)
			if !ok {
				t.Errorf("object is not Error. received=%T (%+v)", evaluated, evaluated)
				continue
			}

			if errObj.Message != expected {
				t.Errorf("wrong error message. expected=%q, received=%q", expected, errObj.Message)
			}
		}
	}
}
