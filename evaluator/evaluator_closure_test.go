package evaluator

import "testing"

func TestClosures(t *testing.T) {
	input := `
    let newAddr = fn(x) {
      fn(y) { x + y };
    };

    let addTwo = newAddr(2);
    addTwo(2);
  `

	testIntegerObject(t, testEval(input), 4)
}
