package ast

import (
	"bytes"

	"llvm.org/llvm/bindings/go/llvm"
)

type FnCall struct {
	Name    string
	Args    []Expr
	fcache  string
	retType string // Setting by parser
}

func (fc *FnCall) Check(c *Context) {
	for _, at := range fc.Args {
		at.Check(c)
	}

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
	fc.fcache = buf.String()
	fc.retType = c.funcRetTyp(fc.fcache).retType
}

func (fc *FnCall) Codegen(c *Context) llvm.Value {
	fn := c.funcRetTyp(fc.fcache).value

	args := []llvm.Value{}
	for _, a := range fc.Args {
		args = append(args, a.Codegen(c))
	}

	return c.Builder.CreateCall(fn, args, "")
}

func (fc *FnCall) Type(*Context) string {
	return fc.retType
}
