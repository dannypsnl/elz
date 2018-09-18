package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

// Expr required method that an expression node have to implement
type Expr interface {
	// Codegen return a llvm.Value
	Codegen(*Context) llvm.Value
	// Type return a type info by string format.
	// It help elz's type system working with AST.
	Type(*Context) string
	Check(*Context)
}

type Access struct {
	From, Get Expr
}

func (a *Access) Check(c *Context) {
	switch a.From.(type) {
	case *Id:
	default:
		panic("Not support anything is not an Identifier in access chain yet!")
	}
}
func (a *Access) Codegen(c *Context) llvm.Value {
	valName := a.From.(*Id).Val
	fieldName := a.Get.(*Id).Val
	typeName, ok := c.VarType(valName)
	if !ok {
		return llvm.Value{}
	}
	idx := c.FindFieldIndexOfType(typeName, fieldName)
	llvmV, ok := c.LLVMValueOfVar(valName)
	if !ok {
		return llvm.Value{}
	}
	ptrToField := c.Builder.CreateInBoundsGEP(
		// struct type is `ref<ref<Type>>` this format, but we can't do GEP on it. So we have to extract first pointer
		c.Builder.CreateLoad(llvmV, ""),
		[]llvm.Value{
			// extract second pointer
			llvm.ConstInt(llvm.Int32Type(), 0, false),
			// get pointer of field
			llvm.ConstInt(llvm.Int32Type(), uint64(idx), false),
		}, "")
	// load value of field
	return c.Builder.CreateLoad(ptrToField, "")
}
func (a *Access) Type(c *Context) string {
	valName := a.From.(*Id).Val
	fieldName := a.Get.(*Id).Val
	typeName, ok := c.VarType(valName)
	if !ok {
		return "not found"
	}
	return c.FieldTypeOf(typeName, fieldName)
}
