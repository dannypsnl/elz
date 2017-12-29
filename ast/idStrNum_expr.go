package ast

import (
	"llvm.org/llvm/bindings/go/llvm"
)

type Id struct {
	Val string
}

func (i *Id) Codegen(ctx *Context) llvm.Value {
	return ctx.Vars[i.Val]
}

// At here we can see, ident's type need to logging in Context
// So Context should send into Type method and Context::Vars
// need a new structure for usage
func (i *Id) Type() string {
	return ""
}

type Str struct {
	Val string
}

func (s *Str) Codegen(ctx *Context) llvm.Value {
	return llvm.ConstString(s.Val, false)
}

func (s *Str) Type() string {
	return "str"
}

type Number struct {
	Val string
}

func (n *Number) Codegen(*Context) llvm.Value {
	return llvm.ConstFloatFromString(llvm.FloatType(), n.Val)
}

func (n *Number) Type() string {
	return "num"
}
