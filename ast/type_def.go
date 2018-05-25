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
	Name      string
	Attrs     []TypeAttr
	signature string
}

func (typ *TypeDef) Check(c *Context) {
	signature := bytes.NewBuffer([]byte{})
	signature.WriteString(typ.Name)
	signature.WriteString("(")
	for i, attr := range typ.Attrs {
		signature.WriteString(attr.Type)
		if i < len(typ.Attrs)-1 {
			signature.WriteRune(',')
		}
	}
	signature.WriteString(")")

	typ.signature = signature.String()
	println(typ.signature)
	c.functions[typ.signature] = &Function{
		value:   llvm.Value{},
		retType: typ.Name,
	}
}

// NOTE: TypeDef is a statement, so should not get value from this AST's Codegen
func (typ *TypeDef) Codegen(c *Context) llvm.Value {
	types := make([]llvm.Type, 0)
	for _, attr := range typ.Attrs {
		types = append(types, c.Type(attr.Type))
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

	c.functions[typ.signature].value = f
	return llvm.Value{}
}
