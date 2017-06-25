package parser

import (
	"testing"
	"simplexSimia/ast"
	"simplexSimia/lexer"
)

func TestSimStatements(t *testing.T) {
	input := `sim x = 5; sim y = 10; sim foobar = 838383;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSimStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testSimStatement(t *testing.T, s ast.Statement, name string)bool {
	if s.TokenLiteral() != "sim" {
		t.Errorf("s.TokenLiteral not 'sim'. got=%d", s.TokenLiteral())
		return false
	}

	simStmt, ok := s.(*ast.SimStatement)
	if !ok {
		t.Errorf("s not *ast.SimStatement. got=%T", s)
		return false
	}

	if simStmt.Name.Value != name {
		t.Errorf("simStmt.Name.Value not '%s'. got=%s", name, simStmt.Name.Value)
		return false
	}

	if simStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, simStmt.Name)
		return false
	}
	return true
}