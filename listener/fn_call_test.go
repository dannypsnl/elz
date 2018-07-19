package listener

import (
	"testing"
	"github.com/dannypsnl/assert"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestFnCall(t *testing.T) {
	assert := assert.NewTester(t)

	src := `
	fn add(l, r: i32) -> i32 { return l + r }
	`

	l := listener(src)
	ee, err := llvm.NewExecutionEngine(l.context.Module)
	if err != nil {
		panic(err)
	}
	gv := ee.RunFunction(ee.FindFunction("add"), []llvm.GenericValue{
		llvm.NewGenericValueFromInt(llvm.Int32Type(), 20, true),
		llvm.NewGenericValueFromInt(llvm.Int32Type(), 10, true),
	})

	assert.Eq(gv.IntWidth(), 32)
	assert.Eq(gv.Int(true), uint64(30))
}

func TestFnCallAsStatement(t *testing.T) {
	src := `
	fn foo() {
		// do nothing
	}
	fn main() {
		foo()
	}
	`

	expected := `call`

	hasTestTemplate(t, src, expected)
}

func TestCallTheFuncNoParam(t *testing.T) {
	src := `
	fn foo() -> i32 { return 10 }
	fn main() {
		let a = foo()
	}
	`

	expected := `call i32 @foo()`

	hasTestTemplate(t, src, expected)
}
