package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

// Assign is AST for code like:
//
//  fn main() {
//    let mut a = 1
//    a = 2
//  }
type Assign struct {
	VarName string
	E       Expr
}

func (a *Assign) Check(c *Context) {
	a.E.Check(c)
}

func (a *Assign) Codegen(c *Context) llvm.Value {
	expr := a.E.Codegen(c)
	val, ok := c.LLVMValueOfVar(a.VarName)
	if ok {
		return c.Builder.CreateStore(expr, val)
	}
	return llvm.Value{}
}
