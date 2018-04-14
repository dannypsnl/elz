package main

import (
	"testing"

	"llvm.org/llvm/bindings/go/llvm"
)

func TestGlobalVarLoad(t *testing.T) {
	mod := llvm.NewModule("test")
	//	context := llvm.NewContext()
	builder := llvm.NewBuilder()

	y := llvm.AddGlobal(mod, llvm.Int32Type(), "y")
	x := llvm.AddGlobal(mod, llvm.Int32Type(), "x")
	x.SetInitializer(llvm.ConstInt(llvm.Int32Type(), 10, false))
	println(x.IsConstant())
	println(x.IsGlobalConstant())
	println(x.IsDeclaration())
	println(llvm.ConstInt(llvm.Int32Type(), 10, true).IsConstant())

	fnType := llvm.FunctionType(llvm.Int32Type(), []llvm.Type{}, false)
	foo := llvm.AddFunction(mod, "foo", fnType)

	block := llvm.AddBasicBlock(foo, "entry")

	builder.SetInsertPointAtEnd(block)

	xLoad := builder.CreateLoad(x, "x.load")
	println(xLoad.IsConstant())
	builder.CreateRet(xLoad)

	builder.ClearInsertionPoint()

	// ======================================================

	fnType = llvm.FunctionType(llvm.VoidType(), []llvm.Type{}, false)
	init := llvm.AddFunction(mod, "init", fnType)

	block = llvm.AddBasicBlock(init, "entry")

	builder.SetInsertPointAtEnd(block)

	fooRes := builder.CreateCall(foo, []llvm.Value{}, "foo_result")
	builder.CreateStore(fooRes, y)

	println(mod.String())
}
