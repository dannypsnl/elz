package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Ref struct {
	E *Id
}

func (r *Ref) Check(c *Context) {
	r.E.Check(c)
}
func (r *Ref) Codegen(c *Context) llvm.Value {
	v, ok := c.Var(r.E.Val)
	if ok {
		return c.Builder.CreateGEP(v, []llvm.Value{
			llvm.ConstInt(llvm.Int32Type(), 0, true),
		}, "")
	}
	return llvm.Value{}
}
func (r *Ref) Type(c *Context) string {
	return "ref<" + r.E.Type(c) + ">"
}
