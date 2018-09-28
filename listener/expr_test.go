package listener

import (
	"testing"

	"github.com/elz-lang/elz/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/dannypsnl/assert"

	"llvm.org/llvm/bindings/go/llvm"
)

func (s *ElzListener) runElzFunction(name string) llvm.GenericValue {
	ee, err := llvm.NewExecutionEngine(s.context.Module)
	if err != nil {
		panic(err)
	}
	return ee.RunFunction(ee.FindFunction(name), []llvm.GenericValue{})
}

func TestAccessChain(t *testing.T) {
	src := `
type Bar (
	i: i32
)

fn main() {
	let b = Bar(10)
	let i = b.i
}
`

	expected := `getelementptr inbounds %Bar, %Bar* %1, i32 0, i32 0`

	hasTestTemplate(t, src, expected)
}

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

func TestBinaryOperator(t *testing.T) {
	assert := assert.NewTester(t)

	assertIsTrue := func(t *testing.T, value llvm.GenericValue) {
		t.Helper()
		assert.Eq(value.IntWidth(), 1)
		assert.Eq(value.Int(false), uint64(1))
	}

	t.Run("Eq, ==", func(t *testing.T) {
		l := listener(`
		fn equal() -> bool {
			return 1 == 1
		}
		`)

		assertIsTrue(t, l.runElzFunction("equal"))
	})
	t.Run("Greater, >", func(t *testing.T) {

		l := listener(`
		fn greater() -> bool {
			return 1 > 0
		}

		fn greater_64() -> bool {
    		return 30'i64 > 0'i64
		}
		`)

		assertIsTrue(t, l.runElzFunction("greater"))
		assertIsTrue(t, l.runElzFunction("greater_64"))
	})
}
