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

	// How to create global const float
	aNum := llvm.ConstFloatFromString(llvm.FloatType(), `3.1415926`)
	bNum := llvm.ConstFloatFromString(llvm.FloatType(), `2.5`)
	//resNum := llvm.ConstFAdd(aNum, bNum)
	//resNum = llvm.ConstFMul(aNum, resNum)
	resNum := ctx.builder.CreateFAdd(aNum, bNum, "tmp")
	resNum = ctx.builder.CreateFMul(aNum, resNum, "tmp")
	gNum := llvm.AddGlobal(ctx.m, llvm.FloatType(), "main::float")
	gNum.SetInitializer(resNum)

	// How to create string
	aStr := llvm.ConstString(`\\a你好, llvm, $@#%^!&!)~!#*(@#+_)(*&GBJNLSfdlbc)`, false)
	gVal := llvm.AddGlobal(ctx.m, aStr.Type(), "main::string")
	gVal.SetInitializer(aStr)

	// How to create a function
	// llvm.FunctionType( retrun type, params types, is var arg: bool )
	ft := llvm.FunctionType(aStr.Type(), []llvm.Type{aStr.Type()}, false)
	llvm.AddFunction(ctx.m, "main::foo_string_string", ft)

	ft = llvm.FunctionType(llvm.Int32Type(), []llvm.Type{
		llvm.Int32Type(),
		llvm.Int32Type(),
	}, false)
	addFn := llvm.AddFunction(ctx.m, "main::add", ft)
	args := []string{"lv", "rv"}
	i := 0
	for _, param := range addFn.Params() {
		param.SetName(args[i])
		i++
	}
	addBlock := llvm.AddBasicBlock(addFn, "entry")
	ctx.builder.SetInsertPointAtEnd(addBlock)
	returnVal := ctx.builder.CreateAdd(addFn.Param(0), addFn.Param(1), "result")
	ctx.builder.CreateRet(
		returnVal,
	)
	ctx.builder.ClearInsertionPoint()

	// main function
	ft = llvm.FunctionType(ctx.ctx.Int32Type(), []llvm.Type{}, false)
	mainFn := llvm.AddFunction(ctx.m, "main", ft)
	mainBlock := llvm.AddBasicBlock(mainFn, "entry")
	ctx.builder.SetInsertPointAtEnd(mainBlock)
	ctx.builder.CreateCall(addFn, []llvm.Value{
		llvm.ConstInt(llvm.Int32Type(), 2, false),
		llvm.ConstInt(llvm.Int32Type(), 3, false),
	}, "tmp")
	ctx.builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), 0, false))
	ctx.builder.ClearInsertionPoint()

	// How to create a user define type
	structTp := llvm.StructType([]llvm.Type{
		llvm.FloatType(),
	}, false)
	insStruct := llvm.ConstNamedStruct(structTp, []llvm.Value{llvm.ConstFloat(llvm.FloatType(), 3.14)})
	gStruct := llvm.AddGlobal(ctx.m, structTp, "main::A struct")
	gStruct.SetInitializer(insStruct)

	engine, err := llvm.NewExecutionEngine(ctx.m)
	if err != nil {
		panic(err)
	}
	result := engine.RunFunction(addFn, []llvm.GenericValue{
		llvm.NewGenericValueFromInt(llvm.Int32Type(), 2, false),
		llvm.NewGenericValueFromInt(llvm.Int32Type(), 3, false),
	})
	fmt.Println(result.Int(false))

	fmt.Println("===================================")
	fmt.Println(ctx.m)
}
