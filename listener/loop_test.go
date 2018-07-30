package listener

import (
	"testing"
	"github.com/dannypsnl/assert"

	"llvm.org/llvm/bindings/go/llvm"
)

func Test_Loop_statement(t *testing.T) {
	assert := assert.NewTester(t)

	l := listener(`
fn loop_it() -> i32 {
	let mut i = 1
	loop {
		match i {
			10 => return i,
			_ => i = i + 1,
		}
	}
}
`)

	ee, err := llvm.NewExecutionEngine(l.context.Module)
	if err != nil {
		panic(err)
	}
	gv := ee.RunFunction(ee.FindFunction("loop_it"), []llvm.GenericValue{})

	assert.Eq(gv.IntWidth(), 32)
	assert.Eq(gv.Int(true), uint64(10))
}
