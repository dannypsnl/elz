package researchllvm

import (
	"testing"

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

func TestArrayAlloc(t *testing.T) {
	ctx := NewTestContext()

	argsRef := []llvm.Type{
		llvm.PointerType(llvm.Int8Type(), 0),
	}
	pft := llvm.FunctionType(llvm.Int32Type(), argsRef, true)
	printf := llvm.AddFunction(ctx.module, "printf", pft)

	arrayType := llvm.ArrayType(llvm.Int32Type(), 5)
	arrayV := llvm.ConstArray(llvm.Int32Type(), []llvm.Value{
		llvm.ConstInt(llvm.Int32Type(), 1, true),
		llvm.ConstInt(llvm.Int32Type(), 2, true),
		llvm.ConstInt(llvm.Int32Type(), 3, true),
		llvm.ConstInt(llvm.Int32Type(), 4, true),
		llvm.ConstInt(llvm.Int32Type(), 5, true),
	})
	array := llvm.AddGlobal(ctx.module, arrayType, "arr")
	array.SetInitializer(arrayV)

	ft := llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	main := llvm.AddFunction(ctx.module, "main", ft)

	entry := llvm.AddBasicBlock(main, "entry")

	ctx.builder.SetInsertPointAtEnd(entry)

	s := ctx.builder.CreateGlobalStringPtr("value: %d\n", "format")

	arr0Addr := ctx.builder.CreateGEP(array, []llvm.Value{
		llvm.ConstInt(llvm.Int64Type(), 0, false),
		llvm.ConstInt(llvm.Int32Type(), 0, false),
	}, "&arr[0]")

	arr0Value := ctx.builder.CreateLoad(arr0Addr, "arr[0]")
	ctx.builder.CreateCall(printf, []llvm.Value{s, arr0Value}, "")

	arrayAlloc := ctx.builder.CreateAlloca(arrayType, "array.alloc")
	//arrayLoad := ctx.builder.CreateLoad(array, "array.load")
	ctx.builder.CreateStore(arrayV, arrayAlloc)
	arr1Addr := ctx.builder.CreateGEP(arrayAlloc, []llvm.Value{
		llvm.ConstInt(llvm.Int64Type(), 0, false),
		llvm.ConstInt(llvm.Int32Type(), 1, false),
	}, "&arr[1]")
	arr1Value := ctx.builder.CreateLoad(arr1Addr, "arr[1]")
	ctx.builder.CreateCall(printf, []llvm.Value{s, arr1Value}, "")

	ctx.builder.CreateRetVoid()

	println(ctx.module.String())
}
