package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		llvm.NewModule("main"),
		llvm.NewContext(),
		make(map[string]llvm.Value),
		llvm.NewBuilder(),
	}
}

type Context struct {
	Module  llvm.Module
	Context llvm.Context
	Vars    map[string]llvm.Value
	Builder llvm.Builder
}
