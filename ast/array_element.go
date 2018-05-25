package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type ArrayElement struct {
	E     Expr
	Index int

	ok bool
}

func (ae *ArrayElement) Check(c *Context) {
	ae.E.Check(c)

	ae.ok = isArrayType(ae.E.Type(c))
}

func (ae *ArrayElement) Codegen(c *Context) llvm.Value {
	if ae.ok {
		expr := ae.E.Codegen(c)
		return c.Builder.CreateExtractValue(expr, ae.Index, "")
	}
	return llvm.Value{}
}

func (ae *ArrayElement) Type(c *Context) string {
	if ae.ok {
		return elemType(ae.E.Type(c))
	}
	return ""
}
