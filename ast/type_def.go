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
	types     []llvm.Type
	t         llvm.Type
}

func (typ *TypeDef) Check(c *Context) {
	typ.types = make([]llvm.Type, 0)
	for _, attr := range typ.Attrs {
		typ.types = append(typ.types, c.Type(attr.Type))
	}
	typ.t = c.Module.Context().StructCreateNamed(typ.Name)
	typ.t.StructSetBody(typ.types, true)

	c.NewType(typ.Name, typ.t)

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
	c.functions[typ.signature] = &Function{
		value:   llvm.Value{},
		retType: typ.Name,
	}
}

// NOTE: TypeDef is a statement, so should not get value from this AST's Codegen
func (typ *TypeDef) Codegen(c *Context) llvm.Value {
	ft := llvm.FunctionType(llvm.PointerType(typ.t, 0), typ.types, false)
	f := llvm.AddFunction(c.Module, typ.Name, ft)

	entry := llvm.AddBasicBlock(f, "entry")

	c.Builder.SetInsertPointAtEnd(entry)

	object := c.Builder.CreateMalloc(typ.t, "object")

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
