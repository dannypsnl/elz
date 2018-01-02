package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Id struct {
	Val string
}

func (i *Id) Codegen(ctx *Context) llvm.Value {
	// FIXME: load value
	return ctx.Vars[i.Val].v
}

// At here we can see, ident's type need to logging in Context
// So Context should send into Type method and Context::Vars
// need a new structure for usage
func (i *Id) Type(ctx *Context) string {
	return ctx.Vars[i.Val].Type
}

type Str struct {
	Val string
}

func (s *Str) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstString(s.Val, false)
}

func (s *Str) Type(*Context) string {
	return "str"
}

type Number struct {
	Val string
}

func (n *Number) Codegen(*Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.FloatType(), n.Val)
}

func (n *Number) Type(*Context) string {
	return "num"
}
