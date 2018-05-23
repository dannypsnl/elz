package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type TypeAttr struct {
	Export bool
	Name   string
	Type   string
}

type TypeDef struct {
	Name  string
	Attrs []TypeAttr
}

func (typ *TypeDef) Check(c *Context) {}

func (typ *TypeDef) Codegen(c *Context) llvm.Value {
	return llvm.Value{}
}
