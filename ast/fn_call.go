package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type FnCall struct {
	Name string
	Args []Expr
}

func (fc *FnCall) Check(c *Context) {
	for _, at := range fc.Args {
		at.Check(c)
	}

}

func (fc *FnCall) Codegen(c *Context) llvm.Value {
	return c.Call(fc.Name, fc.Args...)
}

func (fc *FnCall) Type(c *Context) string {
	return c.funcRetTyp(c.signature(fc.Name, fc.Args...)).retType
}
