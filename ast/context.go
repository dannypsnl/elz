package ast

import (
	"github.com/elz-lang/elz/errors"

	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		Reporter:   errors.NewReporter(),
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
	Reporter   *errors.Reporter
	Module     llvm.Module
	Context    llvm.Context
	GlobalVars map[string]*VarNode
	VarsType   map[string]string
	Builder    llvm.Builder
}
