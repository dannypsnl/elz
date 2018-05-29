package ast

import (
	"bytes"

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
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(fc.Name)
	buf.WriteRune('(')
	for i, at := range fc.Args {
		buf.WriteString(at.Type(c))
		if i < len(fc.Args)-1 {
			buf.WriteRune(',')
		}
	}
	buf.WriteRune(')')
	return c.funcRetTyp(buf.String()).retType
}
