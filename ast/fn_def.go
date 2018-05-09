package ast

import (
	"bytes"

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
	fcache      string
}

func (f *FnDef) Check(c *Context) {
	f.completeParamType()
	f.Ctx = &Context{
		Parent:   c,
		Reporter: c.Reporter,
		Module:   c.Module,
		Context:  c.Context,
		Vars:     make(map[string]llvm.Value),
		VarsType: make(map[string]string),
		Builder:  llvm.NewBuilder(),
	}

	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(f.Name)
	buf.WriteRune('(')
	for i, p := range f.Params {
		f.Ctx.VarsType[p.Name] = p.Type
		buf.WriteString(p.Type)
		if i != len(f.Params)-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(')')
	f.fcache = buf.String()
	c.functions[f.fcache] = &Function{
		value:   llvm.Value{},
		retType: f.RetType,
	}

	for _, stat := range f.Body {
		stat.Check(f.Ctx)
	}
}

func (f *FnDef) Codegen(c *Context) llvm.Value {
	fn := llvm.AddFunction(c.Module, f.Name,
		llvm.FunctionType(f.returnType(c), f.paramsType(), false),
	)

	fc := c.functions[f.fcache]
	fc.value = fn

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

func (f *FnDef) completeParamType() {
	var cache string
	for i := len(f.Params); i > 0; i-- {
		if f.Params[i-1].Type != "" {
			cache = f.Params[i-1].Type
		} else {
			f.Params[i-1].Type = cache
		}
	}
}

func (f *FnDef) returnType(c *Context) llvm.Type {
	rt := f.RetType
	if rt == "" {
		rt = "()"
	}
	if f.Name == "main" {
		if rt == "()" {
			rt = "i32"
		} else {
			c.Reporter.Emit("you can't have return type for main function!")
		}
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
