package ast

import (
	"github.com/elz-lang/elz/errors"

	"llvm.org/llvm/bindings/go/llvm"
)

func NewContext() *Context {
	return &Context{
		Reporter: errors.NewReporter(),
		Module:   llvm.NewModule("main"),
		Context:  llvm.NewContext(),
		Vars:     make(map[string]llvm.Value),
		VarsType: make(map[string]string),
		Builder:  llvm.NewBuilder(),
	}
}

type Context struct {
	Reporter *errors.Reporter
	Module   llvm.Module
	Context  llvm.Context
	Vars     map[string]llvm.Value
	VarsType map[string]string
	Builder  llvm.Builder
}
