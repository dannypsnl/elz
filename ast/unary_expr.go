package ast

import (
	"fmt"

	"llvm.org/llvm/bindings/go/llvm"
)

type UnaryExpr struct {
	E         Expr
	Op        string
	checkedOk bool
}

func (u *UnaryExpr) Check(c *Context) {
	u.E.Check(c)

	// initial it first
	u.checkedOk = true
	exprTyp := u.E.Type(c)

	switch u.Op {
	case "*":
		u.checkedOk = isRefType(exprTyp)
		if !u.checkedOk {
			c.Reporter.Emit(fmt.Sprintf("only can dereference a reference, but type is %s", exprTyp))
		}
	default:
		u.checkedOk = false
	}
}

func (u *UnaryExpr) Codegen(c *Context) llvm.Value {
	if u.checkedOk {
		switch u.Op {
		case "*":
			p := u.E.Codegen(c)
			return c.Builder.CreateLoad(p, "")
		}
	}
	return llvm.Value{}
}

func (u *UnaryExpr) Type(c *Context) string {
	switch u.Op {
	case "*":
		return elemType(u.E.Type(c))
	default:
		return u.E.Type(c)
	}
}
