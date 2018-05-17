package ast

import (
	"fmt"
	"strings"

	"llvm.org/llvm/bindings/go/llvm"
)

type DeRef struct {
	E  Expr
	ok bool
}

func (dr *DeRef) Check(c *Context) {
	dr.E.Check(c)

	exprTyp := dr.E.Type(c)
	if strings.HasPrefix(exprTyp, "ref<") && strings.HasSuffix(exprTyp, ">") {
		dr.ok = true
	} else {
		c.Reporter.Emit(fmt.Sprintf("only can dereference a reference, but type is %s", exprTyp))
	}
}

func (dr *DeRef) Codegen(c *Context) llvm.Value {
	if dr.ok {
		p := dr.E.Codegen(c)
		return c.Builder.CreateLoad(p, "")
	}
	return llvm.Value{}
}

func (dr *DeRef) Type(c *Context) string {
	exprT := dr.E.Type(c)
	if dr.ok {
		return exprT[4 : len(exprT)-1]
	}
	return "can't dereference a not pointer value"
}
