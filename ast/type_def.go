package ast

import (
	"bytes"

	"llvm.org/llvm/bindings/go/llvm"
)

// TypeAttr record attribute's name & type
type TypeAttr struct {
	Export bool
	Name   string
	Typ    string
}

// TypeDef is AST for code like:
//
// ```
// type Foo (
//   attr: i32
// )
// ```
type TypeDef struct {
	Name      string
	Attrs     []TypeAttr
	signature string
	types     []llvm.Type
	t         llvm.Type
}

// Check for TypeDef
func (typ *TypeDef) Check(c *Context) {
	typ.types = make([]llvm.Type, 0)
	for _, attr := range typ.Attrs {
		typ.types = append(typ.types, c.Type(attr.Typ))
	}
	typ.t = c.Module.Context().StructCreateNamed(typ.Name)
	typ.t.StructSetBody(typ.types, true)

	c.NewType(typ.Name, typ.t)

	signature := bytes.NewBuffer([]byte{})
	signature.WriteString(typ.Name)
	signature.WriteString("(")
	for i, attr := range typ.Attrs {
		signature.WriteString(attr.Typ)
		if i < len(typ.Attrs)-1 {
			signature.WriteRune(',')
		}
	}
	signature.WriteString(")")

	typ.signature = signature.String()
	returnType := typ.Name
	c.NewFunc(typ.signature, returnType)
}

// Codegen of TypeDef
//
// NOTE: TypeDef is a statement, so should not get value from this AST's Codegen
func (typ *TypeDef) Codegen(c *Context) llvm.Value {
	ft := llvm.FunctionType(llvm.PointerType(typ.t, 0), typ.types, false)
	typeConstructor := llvm.AddFunction(c.Module, typ.Name, ft)

	entry := llvm.AddBasicBlock(typeConstructor, "entry")

	c.Builder.SetInsertPointAtEnd(entry)

	object := c.Builder.CreateMalloc(typ.t, "object")

	for i, attr := range typ.Attrs {
		valueI := c.Builder.CreateGEP(object, []llvm.Value{
			llvm.ConstInt(llvm.Int32Type(), 0, false),
			llvm.ConstInt(llvm.Int32Type(), uint64(i), false),
		}, attr.Name)
		c.Builder.CreateStore(typeConstructor.Param(i), valueI)
	}

	c.Builder.CreateRet(object)

	c.Builder.ClearInsertionPoint()

	c.FuncValue(typ.signature, typeConstructor)
	return llvm.Value{}
}
