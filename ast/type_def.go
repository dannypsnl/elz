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

// NOTE: TypeDef is a statement, so should not get value from this AST's Codegen
func (typ *TypeDef) Codegen(c *Context) llvm.Value {
	types := make([]llvm.Type, 0)
	for _, attr := range typ.Attrs {
		types = append(types, LLVMType(attr.Type))
	}
	t := c.Context.StructCreateNamed(typ.Name)
	t.StructSetBody(types, true)

	c.NewType(typ.Name, t)

	return llvm.Value{}
}
