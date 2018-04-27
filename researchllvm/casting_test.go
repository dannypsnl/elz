package researchllvm

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestCasting(t *testing.T) {
	c := NewTestContext()

	argsRef := []llvm.Type{
		llvm.PointerType(llvm.Int8Type(), 0),
	}
	pft := llvm.FunctionType(llvm.Int32Type(), argsRef, true)
	printf := llvm.AddFunction(c.module, "printf", pft)

	ft := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	main := llvm.AddFunction(c.module, "main", ft)
	entry := llvm.AddBasicBlock(main, "entry")

	c.builder.SetInsertPointAtEnd(entry)

	bar := c.builder.CreateAlloca(llvm.Int32Type(), "bar")
	c.builder.CreateStore(llvm.ConstInt(llvm.Int32Type(), 10, true), bar)
	barLoad := c.builder.CreateLoad(bar, "bar.load")

	v := cast(llvm.ZExt, llvm.Int64Type(), c.builder, barLoad)
	s := c.builder.CreateGlobalStringPtr("hello world, number: %d\n", "format")
	c.builder.CreateCall(printf, []llvm.Value{s, v}, "p")
	v = cast(llvm.SExt, llvm.Int64Type(), c.builder, barLoad)
	c.builder.CreateCall(printf, []llvm.Value{s, v}, "p")
	v = cast(llvm.Trunc, llvm.Int16Type(), c.builder, barLoad)
	c.builder.CreateCall(printf, []llvm.Value{s, v}, "p")

	c.builder.CreateRet(barLoad)

	println(c.module.String())
}

func cast(castBy llvm.Opcode, castTo llvm.Type, builder llvm.Builder, v llvm.Value) llvm.Value {
	return builder.CreateCast(v, castBy, castTo, "tmp")
}
