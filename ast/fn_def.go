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

func (f *FnDef) Codegen(ctx *Context) llvm.Value {
	paramsT := []llvm.Type{}
	for _, v := range f.Params {
		paramsT = append(paramsT, convertToLLVMType(v.Type))
	}

	rt := f.RetType
	if rt == "" {
		rt = "()"
	}
	if f.Name == "main" && rt == "()" {
		rt = "i32"
	}
	retT := convertToLLVMType(rt)

	ft := llvm.FunctionType(retT, paramsT, false)
	fn := llvm.AddFunction(ctx.Module, f.Name, ft)

	for i, param := range fn.Params() {
		param.SetName(f.Params[i].Name)
	}

	entryPoint := llvm.AddBasicBlock(fn, "entry")

	ctx.Builder.SetInsertPointAtEnd(entryPoint)
	for _, stat := range f.Body {
		stat.Codegen(ctx)
	}
	ctx.Builder.ClearInsertionPoint()
	return fn
}
