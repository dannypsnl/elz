package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Ref struct {
	E          Expr
	compilable bool
}

func (r *Ref) Check(c *Context) {
	r.E.Check(c)

	if _, ok := r.E.(*Id); ok {
		r.compilable = true
	}
}
func (r *Ref) Codegen(c *Context) llvm.Value {
	if r.compilable {
		e := r.E.(*Id)
		v, ok := c.Var(e.Val)
		if ok {
			return c.Builder.CreateGEP(v, []llvm.Value{
				llvm.ConstInt(llvm.Int32Type(), 0, true),
			}, "")
		}
	}
	return llvm.Value{}
}
func (r *Ref) Type(c *Context) string {
	return "ref<" + r.E.Type(c) + ">"
}
