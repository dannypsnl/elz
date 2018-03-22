package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		Module:     llvm.NewModule("main"),
		Context:    llvm.NewContext(),
		GlobalVars: make(map[string]*VarNode),
		VarsType:   make(map[string]string),
		Builder:    llvm.NewBuilder(),
	}
}

type VarNode struct {
	v    llvm.Value
	Type string
}

type Context struct {
	Module     llvm.Module
	Context    llvm.Context
	GlobalVars map[string]*VarNode
	VarsType   map[string]string
	Builder    llvm.Builder
}
