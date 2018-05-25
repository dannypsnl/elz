package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type ArrayElement struct {
	E     Expr
	Index int
}

func (ae *ArrayElement) Check(c *Context) {
	ae.E.Check(c)
}

func (ae *ArrayElement) Codegen(c *Context) llvm.Value {
	expr := ae.E.Codegen(c)
	return c.Builder.CreateExtractValue(expr, ae.Index, "")
}

func (ae *ArrayElement) Type(c *Context) string {
	return elemType(ae.E.Type(c))
}
