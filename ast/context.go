package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		llvm.NewModule("main"),
		llvm.NewContext(),
		make(map[string]*VarNode),
		llvm.NewBuilder(),
	}
}

type VarNode struct {
	v    llvm.Value
	Type string
}

type Context struct {
	Module  llvm.Module
	Context llvm.Context
	Vars    map[string]*VarNode
	Builder llvm.Builder
}
