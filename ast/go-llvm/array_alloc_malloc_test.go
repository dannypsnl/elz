package main

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

	ar := ctx.builder.CreateArrayAlloca(llvm.Int32Type(), llvm.ConstInt(llvm.Int32Type(), 15, false), "ar")
	ctx.builder.CreateGEP(ar, []llvm.Value{llvm.ConstInt(llvm.Int32Type(), 0, false)}, "sp")
	arrLoad := ctx.builder.CreateLoad(array, "array.load")
	ctx.builder.CreateExtractValue(arrLoad, 0, "arr0")

	println(ctx.module.String())
}
