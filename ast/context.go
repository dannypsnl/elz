package ast

import (
	"bytes"

	"github.com/elz-lang/elz/errors"

	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		Parent:    nil,
		Reporter:  errors.NewReporter(),
		Module:    llvm.NewModule("main"),
		Context:   llvm.NewContext(),
		Vars:      make(map[string]llvm.Value),
		VarsType:  make(map[string]string),
		Types:     make(map[string]llvm.Type),
		functions: make(map[string]*Function),
		Builder:   llvm.NewBuilder(),
	}
}

type Context struct {
	Parent    *Context
	Reporter  *errors.Reporter
	Module    llvm.Module
	Context   llvm.Context
	Vars      map[string]llvm.Value
	VarsType  map[string]string
	Types     map[string]llvm.Type
	functions map[string]*Function
	Builder   llvm.Builder
}

type Function struct {
	value   llvm.Value
	retType string
}

func (c *Context) NewType(name string, t llvm.Type) {
	c.Types[name] = t
}

func (c *Context) Type(name string) llvm.Type {
	typ := LLVMType(name)
	if typ.String() != "VoidType" {
		return typ
	}
	if name == "()" {
		return llvm.VoidType()
	}
	return llvm.PointerType(c.Module.GetTypeByName(name), 0)
}

func (c *Context) Call(funcName string, exprs ...Expr) llvm.Value {
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(funcName)
	buf.WriteRune('(')
	for i, e := range exprs {
		e.Type(c)
		buf.WriteString(e.Type(c))
		if len(exprs)-1 > i {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(')')
	signature := buf.String()

	args := []llvm.Value{}
	for _, e := range exprs {
		args = append(args, e.Codegen(c))
	}

	if c.funcRetTyp(signature) != nil {
		fn := c.funcRetTyp(signature).value
		return c.Builder.CreateCall(fn, args, "")
	}
	return llvm.Value{}
}

func (c *Context) funcRetTyp(signature string) *Function {
	if f, ok := c.functions[signature]; ok {
		return f
	}
	if c.Parent != nil {
		return c.Parent.funcRetTyp(signature)
	}
	return nil
}

func (c *Context) VarValue(name string, value llvm.Value) {
	c.Vars[name] = value
}

func (c *Context) NewVar(name string, typ string) {
	// FIXME: let vars contains Var Node only, then Var Node contains more info is better.
	c.VarsType[name] = typ
	// FIXME: Missing export & mutable or not info
}

func (c *Context) Var(name string) (llvm.Value, bool) {
	v, ok := c.Vars[name]
	if ok {
		return v, true
	}
	if c.Parent != nil {
		return c.Parent.Var(name)
	}
	return llvm.Value{}, false
	// It will cause easy panic in llvm system
	// To match the type have to write down this line
	// p.s. Because var not found is common, we can't panic this
}

func (c *Context) VarType(name string) (string, bool) {
	v, ok := c.VarsType[name]
	if ok {
		return v, true
	}
	if c.Parent != nil {
		return c.Parent.VarType(name)
	}
	return "no this var", false
}
