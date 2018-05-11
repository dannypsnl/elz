package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type UnaryExpr struct {
	E  Expr
	Op string
}

func (u *UnaryExpr) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstFNeg(u.E.Codegen(ctx))
}
func (u *UnaryExpr) Type(ctx *Context) string {
	return u.E.Type(ctx)
}
