package ast

import (
	"bytes"

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
	t := c.Module.Context().StructCreateNamed(typ.Name)
	t.StructSetBody(types, true)

	c.NewType(typ.Name, t)

	ft := llvm.FunctionType(llvm.PointerType(t, 0), types, false)
	f := llvm.AddFunction(c.Module, typ.Name, ft)

	entry := llvm.AddBasicBlock(f, "entry")

	c.Builder.SetInsertPointAtEnd(entry)

	object := c.Builder.CreateMalloc(t, "object")

	for i, attr := range typ.Attrs {
		valueI := c.Builder.CreateGEP(object, []llvm.Value{
			llvm.ConstInt(llvm.Int32Type(), 0, false),
			llvm.ConstInt(llvm.Int32Type(), uint64(i), false),
		}, attr.Name)
		c.Builder.CreateStore(f.Param(i), valueI)
	}

	c.Builder.CreateRet(object)

	c.Builder.ClearInsertionPoint()

	signature := bytes.NewBuffer([]byte{})
	signature.WriteString("(")
	for i, attr := range typ.Attrs {
		if i != 0 {
			signature.WriteString("," + attr.Type)
		} else {
			signature.WriteString(attr.Type)
		}
	}
	signature.WriteString(")")

	c.functions[signature.String()] = &Function{
		value:   f,
		retType: typ.Name,
	}
	return llvm.Value{}
}
