package researchllvm

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type TestContext struct {
	module  llvm.Module
	context llvm.Context
	builder llvm.Builder
}

func NewTestContext() *TestContext {
	return &TestContext{
		module:  llvm.NewModule("test"),
		context: llvm.NewContext(),
		builder: llvm.NewBuilder(),
	}
}

func (c *TestContext) call(fn llvm.Value, args ...llvm.Value) llvm.Value {
	return c.builder.CreateCall(fn, args, ".call_tmp")
}
