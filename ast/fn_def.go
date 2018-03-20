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

	builder := llvm.NewBuilder()
	builder.SetInsertPointAtEnd(entryPoint)
	// TODO: for _, stat := range f.Body {
	//builder.Insert(stat.Codegen(ctx))
	//}
	builder.CreateRet(llvm.ConstFloat(llvm.FloatType(), 3.14))
	builder.ClearInsertionPoint()
	return fn
}
