package parser

import (
	"monkey-interpreter/ast"
	"monkey-interpreter/lexer"
	"testing"
)

func TestReturnStatement(t *testing.T) {
	input := `
    return 5;
    return 10;
    return 993322;
  `

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. received=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. received=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. received %q", returnStmt.TokenLiteral())
		}
	}
}
