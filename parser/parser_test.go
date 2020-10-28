package parser

import (
	"testing"

	"github.com/valter4578/vlang/ast"
	"github.com/valter4578/vlang/lexer"
)

func TestLetStatments(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let c = 838383; 
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatalf("Error in ParseProgram(): returned nil")
	}

	if len(program.Statments) != 3 {
		t.Fatalf("Error: program.Statments don't containt 3 statements, got:%d", len(program.Statments))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"c"},
	}

	for i, tt := range tests {
		statment := program.Statements[i]
		if !testLetStatement(t, statment, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	// letStatement, ok := s.(*ast.LetStatement)
	letStatment, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatment. got=%T", s)
		return false
	}

	if letStatment.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStatment.Name.Value)
		return false
	}

	if letStatment.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStatment.Name)
		return false
	}

	return true
}
