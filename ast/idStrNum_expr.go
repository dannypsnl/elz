package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Id struct {
	Val string
}

func (i *Id) Codegen(ctx *Context) llvm.Value {
	println(i.Val)
	// FIXME: load value
	return ctx.Builder.CreateLoad(
		ctx.Vars[i.Val].v,
		i.Val,
	)
}

// At here we can see, ident's type need to logging in Context
// So Context should send into Type method and Context::Vars
// need a new structure for usage
func (i *Id) Type(ctx *Context) string {
	println(i.Val)
	return ctx.VarsType[i.Val]
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
