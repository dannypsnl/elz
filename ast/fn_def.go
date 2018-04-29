package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Param struct {
	Name string
	Type string
}

type FnDef struct {
	Export      bool
	Name        string
	Params      []*Param
	Body        []Stat
	RetType     string
	Ctx         *Context
	IsExternDef bool
}

func (f *FnDef) Check(ctx *Context) {
	f.setupMissParamType()
	f.Ctx = &Context{
		Parent:   ctx,
		Reporter: ctx.Reporter,
		Module:   ctx.Module,
		Context:  ctx.Context,
		Vars:     make(map[string]llvm.Value),
		VarsType: make(map[string]string),
		Builder:  llvm.NewBuilder(),
	}

	for _, p := range f.Params {
		f.Ctx.VarsType[p.Name] = p.Type
	}
	for _, stat := range f.Body {
		stat.Check(f.Ctx)
	}
}

func (f *FnDef) Codegen(ctx *Context) llvm.Value {
	fn := llvm.AddFunction(ctx.Module, f.Name,
		llvm.FunctionType(f.returnType(), f.paramsType(), false),
	)

	// is a declaration in extern block for ffi we don't generate the statement for it
	if !f.IsExternDef {
		entryPoint := llvm.AddBasicBlock(fn, "entry")
		f.Ctx.Builder.SetInsertPointAtEnd(entryPoint)

		for i, param := range fn.Params() {
			param.SetName(f.Params[i].Name)
			f.Ctx.Vars[f.Params[i].Name] = param
		}

		for _, stat := range f.Body {
			stat.Codegen(f.Ctx)
		}
		f.Ctx.Builder.ClearInsertionPoint()
		if f.Name == "main" {
			generateMainFn(f.Ctx.Builder, entryPoint)
		}
	}
	return fn
}

func generateMainFn(builder llvm.Builder, entryPoint llvm.BasicBlock) {
	builder.SetInsertPointAtEnd(entryPoint)
	builder.CreateRet(llvm.ConstInt(LLVMType("i32"), 0, false))
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
	retT := LLVMType(rt)
	return retT
}

func (f *FnDef) paramsType() []llvm.Type {
	paramsT := []llvm.Type{}
	for _, v := range f.Params {
		paramsT = append(paramsT, LLVMType(v.Type))
	}
	return paramsT
}
