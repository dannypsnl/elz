package listener

import (
	"testing"

	"github.com/elz-lang/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestF32(t *testing.T) {
	src := `
	x = 3.2
	`

	expected := `
@x = global float 0x40099999A0000000
`

	hasTestTemplate(t, src, expected)
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
	src := `
	a = 3'i8
	b = 3'i16
	c = 3'i32
	d = 3'i64
	e = 3'f32
	f = 3'f64
	`

	expected := `
@a = global i8 3
@b = global i16 3
@c = global i32 3
@d = global i64 3
@e = global float 3.000000e+00
@f = global double 3.000000e+00
`

	hasTestTemplate(t, src, expected)
}

func TestFloatSuffix(t *testing.T) {
	src := `
	a = 3.14'f32
	b = 3.14'f64
	`

	expected := `
@a = global float 0x40091EB860000000
@b = global double 3.140000e+00
`

	hasTestTemplate(t, src, expected)

}

func TestEq(t *testing.T) {
	src := `
	a = 1
	b = 1

	fn main() {
	  c = a == b
	}
	`

	expected := `
  %.eq_tmp = icmp eq i32 %0, %1
`

	hasTestTemplate(t, src, expected)
}
