package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Return struct {
	Expr Expr
}

func (r *Return) Check(c *Context) {
	r.Expr.Check(c)
}

func (r *Return) Codegen(c *Context) llvm.Value {
	return c.Builder.CreateRet(r.Expr.Codegen(c))
}

func (r *Return) Type(c *Context) string {
	return r.Expr.Type(c)
}
