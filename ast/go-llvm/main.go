package main

import (
	"fmt"
	"llvm.org/llvm/bindings/go/llvm"
)

type Context struct {
	m       llvm.Module
	ctx     llvm.Context
	builder llvm.Builder
}

func main() {
	ctx := &Context{
		m:       llvm.NewModule("main"),
		ctx:     llvm.NewContext(),
		builder: llvm.NewBuilder(),
	}

	// How to create string
	aStr := llvm.ConstString(`\\a你好, llvm, $@#%^!&!)~!#*(@#+_)(*&GBJNLSfdlbc)`, false)
	gVal := llvm.AddGlobal(ctx.m, aStr.Type(), "main::string")
	gVal.SetInitializer(aStr)

	// How to create a function
	ft := llvm.FunctionType(aStr.Type(), []llvm.Type{aStr.Type()}, false)
	llvm.AddFunction(ctx.m, "main::foo_string_string", ft)

	ft = llvm.FunctionType(ctx.ctx.FloatType(), []llvm.Type{}, false)
	llvm.AddFunction(ctx.m, "main::foo_float", ft)

	// How to create a user define type
	structTp := llvm.StructType([]llvm.Type{
		llvm.FloatType(),
	}, false)
	insStruct := llvm.ConstNamedStruct(structTp, []llvm.Value{llvm.ConstFloat(llvm.FloatType(), 3.14)})
	gStruct := llvm.AddGlobal(ctx.m, structTp, "main::A struct")
	gStruct.SetInitializer(insStruct)

	fmt.Println(ctx.m)
}
