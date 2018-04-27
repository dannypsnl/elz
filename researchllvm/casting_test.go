package researchllvm

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestCasting(t *testing.T) {
	c := NewTestContext()

	argsRef := []llvm.Type{}
	arrT := llvm.PointerType(llvm.Int8Type(), 0)
	argsRef = append(argsRef, arrT)
	pft := llvm.FunctionType(llvm.Int32Type(), argsRef, false)
	printf := llvm.AddFunction(c.module, "printf", pft)

	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	main := llvm.AddFunction(c.module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")

	c.builder.SetInsertPointAtEnd(entry)

	bar := c.builder.CreateAlloca(llvm.Int32Type(), "bar")
	c.builder.CreateStore(llvm.ConstInt(llvm.Int32Type(), 10, true), bar)
	barLoad := c.builder.CreateLoad(bar, "bar.load")

	cast(llvm.ZExt, llvm.Int64Type(), c.builder, barLoad, printf)
	cast(llvm.SExt, llvm.Int64Type(), c.builder, barLoad, printf)
	cast(llvm.Trunc, llvm.Int16Type(), c.builder, barLoad, printf)

	c.builder.CreateRet(barLoad)

	println(c.module.String())
}

func cast(castBy llvm.Opcode, castTo llvm.Type, builder llvm.Builder, v llvm.Value, printf llvm.Value) {
	s := builder.CreateGlobalStringPtr("hello world\n", "format")
	builder.CreateCast(v, castBy, castTo, "tmp")
	builder.CreateCall(printf, []llvm.Value{s}, "p")
}
