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
	retT := convertToLLVMType(f.RetType)
	ft := llvm.FunctionType(retT, paramsT, false)
	fn := llvm.AddFunction(ctx.Module, f.Name, ft)
	entryPoint := llvm.AddBasicBlock(fn, "entry")

	ctx.Builder.SetInsertPointAtEnd(entryPoint)
	// TODO: for _, stat := range f.Body {
	//     stat.Codegen(ctx)
	// }
	ctx.Builder.CreateRet(llvm.ConstFloat(llvm.FloatType(), 3.14))
	ctx.Builder.ClearInsertionPoint()
	return fn
}
