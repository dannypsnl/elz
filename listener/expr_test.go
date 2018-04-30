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

	if listener.context.Reporter.ErrMsgs[0] != "global var: x, it's type is: i32, but receive: f32" {
		t.Error("var type is not equal to expression's type should be an error")
	}
}

func TestIntSuffix(t *testing.T) {
	res := NewParse(`
	a = 3'i8
	b = 3'i16
	c = 3'i32
	d = 3'i64
	e = 3'f32
	f = 3'f64
	`)

	expected := `; ModuleID = 'main'
source_filename = "main"

@a = global i8 3
@b = global i16 3
@c = global i32 3
@d = global i64 3
@e = global float 3.000000e+00
@f = global double 3.000000e+00
`

	if expected != res {
		t.Errorf("suffix parse Error; expected: `%s`\nactual: `%s`", expected, res)
	}
}
