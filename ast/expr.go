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
	From Expr
	Get  Expr
}

func (a *Access) Check(c *Context) {}
func (a *Access) Codegen(c *Context) llvm.Value {
	return llvm.Value{}
}
func (a *Access) Type(*Context) string {
	return "fail"
}
