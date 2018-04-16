package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Param struct {
	Name string
	Type string
}

type FnDef struct {
	Export  bool
	Name    string
	Params  []*Param
	Body    []Stat
	RetType string
}

func (f *FnDef) Check(ctx *Context) {
	for _, stat := range f.Body {
		stat.Check(ctx)
	}
}

func (f *FnDef) Codegen(ctx *Context) llvm.Value {
	f.setupMissParamType()

	fn := llvm.AddFunction(ctx.Module, f.Name,
		llvm.FunctionType(f.returnType(), f.paramsType(), false),
	)

	for i, param := range fn.Params() {
		param.SetName(f.Params[i].Name)
	}

	entryPoint := llvm.AddBasicBlock(fn, "entry")

	ctx.Builder.SetInsertPointAtEnd(entryPoint)
	for _, stat := range f.Body {
		stat.Codegen(ctx)
	}
	ctx.Builder.ClearInsertionPoint()
	if f.Name == "main" {
		generateMainFn(ctx.Builder, entryPoint)
	}
	return fn
}

func generateMainFn(builder llvm.Builder, entryPoint llvm.BasicBlock) {
	builder.SetInsertPointAtEnd(entryPoint)
	builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), 0, false))
	builder.ClearInsertionPoint()
}

func (f *FnDef) setupMissParamType() {
	var cache string
	for i := len(f.Params); i > 0; i-- {
		if f.Params[i-1].Type != "" {
			cache = f.Params[i-1].Type
		} else {
			f.Params[i-1].Type = cache
		}
	}
}

func (f *FnDef) returnType() llvm.Type {
	rt := f.RetType
	if rt == "" {
		rt = "()"
	}
	// FIXME: if main function define it's return type, it's an error
	if f.Name == "main" && rt == "()" {
		rt = "i32"
	}
	retT := convertToLLVMType(rt)
	return retT
}

func (f *FnDef) paramsType() []llvm.Type {
	paramsT := []llvm.Type{}
	for _, v := range f.Params {
		paramsT = append(paramsT, convertToLLVMType(v.Type))
	}
	return paramsT
}
