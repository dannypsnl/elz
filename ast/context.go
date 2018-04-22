package ast

import (
	"github.com/elz-lang/elz/errors"

	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		Parent:   nil,
		Reporter: errors.NewReporter(),
		Module:   llvm.NewModule("main"),
		Context:  llvm.NewContext(),
		Vars:     make(map[string]llvm.Value),
		VarsType: make(map[string]string),
		Builder:  llvm.NewBuilder(),
	}
}

type Context struct {
	Parent   *Context
	Reporter *errors.Reporter
	Module   llvm.Module
	Context  llvm.Context
	Vars     map[string]llvm.Value
	VarsType map[string]string
	Builder  llvm.Builder
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
