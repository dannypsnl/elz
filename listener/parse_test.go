package listener

import (
	"testing"

	"github.com/elz-lang/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"llvm.org/llvm/bindings/go/llvm"
)

func TestParse(t *testing.T) {
	res := NewParse(`
	a = 10
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

@a = global i32 10
`

	if res != expected {
		t.Errorf("expected: `%s`\nactual: `%s`", expected, res)
	}
}

func TestVerifyModule(t *testing.T) {
	input := antlr.NewInputStream(`a = 10`)
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	listener := /*listener.*/ New()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if err := llvm.VerifyModule(listener.Module(), llvm.PrintMessageAction); err != nil {
		t.Error("VerifyModule fail")
	}
}
