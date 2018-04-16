package main

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestLocalLoad(t *testing.T) {
	c := NewTestContext()

	ft := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	main := llvm.AddFunction(c.module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")

	c.builder.SetInsertPointAtEnd(entry)

	bar := c.builder.CreateAlloca(llvm.Int32Type(), "bar")
	c.builder.CreateStore(llvm.ConstInt(llvm.Int32Type(), 10, true), bar)
	barLoad := c.builder.CreateLoad(bar, "bar.load")
	c.builder.CreateICmp(llvm.IntEQ, barLoad, llvm.ConstInt(llvm.Int32Type(), 10, true), "cmp")

	c.builder.ClearInsertionPoint()

	println(c.module.String())
}
