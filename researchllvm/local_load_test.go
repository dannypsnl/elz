package researchllvm

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestLocalLoad(t *testing.T) {
	c := NewTestContext()

	pft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{llvm.PointerType(llvm.Int8Type(), 0)}, true)
	printf := llvm.AddFunction(c.module, "printf", pft)

	ft := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	main := llvm.AddFunction(c.module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")

	c.builder.SetInsertPointAtEnd(entry)

	tstr := c.builder.CreateGlobalStringPtr("%d\n", "")

	bar := c.builder.CreateAlloca(llvm.Int32Type(), "bar")
	c.builder.CreateStore(llvm.ConstInt(llvm.Int32Type(), 10, true), bar)
	barLoad := c.builder.CreateLoad(bar, "bar.load")
	c.call(printf, tstr, barLoad)
	c.builder.CreateICmp(llvm.IntEQ, barLoad, llvm.ConstInt(llvm.Int32Type(), 10, true), "cmp")

	c.builder.ClearInsertionPoint()

	println(c.module.String())
}
