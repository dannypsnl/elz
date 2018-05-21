package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Assign struct {
	VarName string
	E       Expr
}

func (a *Assign) Check(c *Context) {
	a.E.Check(c)

}

func (a *Assign) Codegen(c *Context) llvm.Value {
	expr := a.E.Codegen(c)
	val, ok := c.Var(a.VarName)
	if ok {
		return c.Builder.CreateStore(expr, val)
	}
	return llvm.Value{}
}
