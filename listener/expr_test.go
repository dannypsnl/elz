package listener

import (
	"testing"

	"github.com/elz-lang/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestF32(t *testing.T) {
	res := NewParse(`
	x = 3.2
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

@x = global float 0x40099999A0000000
`

	if expected != res {
		t.Errorf("f32 parse Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}

func TestTypeErrorInVarDef(t *testing.T) {
	input := antlr.NewInputStream(`x: i32 = 3.2`)
	lexer := parser.NewElzLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewElzParser(stream)
	p.BuildParseTrees = true
	tree := p.Prog()
	listener := /*listener.*/ New()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)

	if listener.context.Reporter.HasNoError() {
		t.Error("fkngberbjobdfbnj;dfb;")
	}
}
