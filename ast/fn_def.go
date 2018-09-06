package ast

import (
	"bytes"
	"fmt"

	"github.com/elz-lang/elz/util"

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
	Notations   []util.Notation
	isExternDef bool
	fnSignature string
}

func (f *FnDef) Check(c *Context) {
	f.completeParamType()
	f.Ctx = NewContext()
	f.Ctx.Parent = c
	f.Ctx.Reporter = c.Reporter
	f.Ctx.Module = c.Module
	f.Ctx.Context = c.Context
	f.Ctx.Builder = llvm.NewBuilder()

	for _, nota := range f.Notations {
		if nota.Leading == "extern" {
			if nota.Content[0] == "c" {
				f.isExternDef = true
			} else {
				c.Reporter.Emit(fmt.Sprintf("Not support #[extern(%s)] yet", nota.Content[0]))
			}
		} else {
			c.Reporter.Emit(fmt.Sprintf("Not support notation #[%s] yet", nota.Leading))
		}
	}

	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(f.Name)
	buf.WriteRune('(')
	for i, p := range f.Params {
		f.Ctx.NewVar(p.Name, p.Type)
		buf.WriteString(p.Type)
		if i != len(f.Params)-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(')')
	f.fnSignature = buf.String()
	c.NewFunc(f.fnSignature, f.RetType)

	for _, stat := range f.Body {
		stat.Check(f.Ctx)
	}
}

func (f *FnDef) Codegen(c *Context) llvm.Value {
	fn := llvm.AddFunction(c.Module, f.Name,
		llvm.FunctionType(f.returnType(c), f.paramsType(c), false),
	)

	c.FuncValue(f.fnSignature, fn)

	// is a declaration in extern block for ffi we don't generate the statement for it
	if !f.isExternDef {
		entryPoint := llvm.AddBasicBlock(fn, "entry")
		f.Ctx.Builder.SetInsertPointAtEnd(entryPoint)

		for i, param := range fn.Params() {
			param.SetName(f.Params[i].Name)
			f.Ctx.VarValue(f.Params[i].Name+" param", param)
		}

		for _, stat := range f.Body {
			stat.Codegen(f.Ctx)
		}
		if f.Name == "main" {
			f.Ctx.Builder.CreateRet(llvm.ConstInt(llvm.Int32Type(), 0, false))
		} else if f.RetType == "" {
			f.Ctx.Builder.CreateRetVoid()
		}
	}
	return fn
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
		if rt != "()" {
			c.Reporter.Emit("you can't have return type for main function!")
		}
		rt = "i32"
	}
	return c.Type(rt)
}

func (f *FnDef) paramsType(c *Context) []llvm.Type {
	paramsT := make([]llvm.Type, 0)
	for _, v := range f.Params {
		paramsT = append(paramsT, c.Type(v.Type))
	}
	return paramsT
}
