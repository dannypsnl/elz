package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Ref struct {
	E Expr
}

func (r *Ref) Check(c *Context) {
	r.E.Check(c)
}
func (r *Ref) Codegen(c *Context) llvm.Value {
	return c.Builder.CreatePointerCast(
		r.E.Codegen(c),
		LLVMType(r.E.Type(c)),
		".ref_tmp",
	)
}
func (r *Ref) Type(c *Context) string {
	return "ref<" + r.E.Type(c) + ">"
}
